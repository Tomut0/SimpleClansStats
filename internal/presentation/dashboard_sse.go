package presentation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/application"
	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/google/uuid"
)

func (s *ServerHandler) GetDashboardLayout(w http.ResponseWriter, r *http.Request) {
	layout, err := s.dashboard.Layout(r.Context())
	if err != nil {
		http.Error(w, "could not load layout", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(layout)
}

func (s *ServerHandler) SaveDashboardLayout(w http.ResponseWriter, r *http.Request) {
	if status := s.adminStatus(r.Context()); status != http.StatusOK {
		http.Error(w, http.StatusText(status), status)
		return
	}
	userID, ok := s.sessions.Get(r.Context(), "user_id").(int64)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, 16<<10))
	if err != nil {
		http.Error(w, "invalid layout", http.StatusBadRequest)
		return
	}
	var layout domain.DashboardLayout
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&layout) != nil || decoder.Decode(new(any)) != io.EOF {
		http.Error(w, "invalid layout", http.StatusBadRequest)
		return
	}
	if err := s.dashboard.SaveLayout(r.Context(), userID, layout); err != nil {
		if err == application.ErrInvalidDashboardLayout {
			http.Error(w, "invalid layout", http.StatusBadRequest)
		} else {
			http.Error(w, "could not save layout", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DashboardEvents streams the latest collector snapshot whenever it changes.
func (s *ServerHandler) DashboardEvents(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.sessions.Get(r.Context(), "user_id").(int64); !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	serverID := r.URL.Query().Get("serverId")
	if _, err := uuid.Parse(serverID); err != nil {
		http.Error(w, "serverId is required", http.StatusBadRequest)
		return
	}
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("X-Accel-Buffering", "no")
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	heartbeat := time.NewTicker(15 * time.Second)
	defer heartbeat.Stop()
	last := ""
	for {
		if s.dashboard == nil {
			http.Error(w, "dashboard unavailable", http.StatusServiceUnavailable)
			return
		}
		payload, err := s.dashboard.Data(r.Context(), serverID)
		encoded, _ := json.Marshal(payload)
		if err == nil && len(payload.State) > 0 && string(encoded) != last {
			data := encoded
			if _, err = fmt.Fprintf(w, "event: snapshot\ndata: %s\n\n", data); err != nil {
				return
			}
			flusher.Flush()
			last = string(encoded)
		} else if err != nil {
			return
		}
		select {
		case <-r.Context().Done():
			return
		case <-ticker.C:
		case <-heartbeat.C:
			if _, err := fmt.Fprint(w, ": keepalive\n\n"); err != nil {
				return
			}
			flusher.Flush()
		}
	}
}

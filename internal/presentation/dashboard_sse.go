package presentation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// DashboardEvents streams the latest collector snapshot whenever it changes.
func (s *ServerHandler) DashboardEvents(w http.ResponseWriter, r *http.Request) {
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

	if s.dashboard == nil {
		http.Error(w, "dashboard unavailable", http.StatusServiceUnavailable)
		return
	}

	ctx := r.Context()
	poll := time.NewTicker(2 * time.Second)
	defer poll.Stop()
	keepalive := time.NewTicker(15 * time.Second)
	defer keepalive.Stop()
	lastSnapshot := ""
	for {
		payload, err := s.dashboard.Data(ctx, serverID)
		if err != nil {
			return
		}
		encoded, _ := json.Marshal(payload)
		snapshot := string(encoded)
		if len(payload.State) > 0 && snapshot != lastSnapshot {
			if _, err := fmt.Fprintf(w, "event: snapshot\ndata: %s\n\n", snapshot); err != nil {
				return
			}
			flusher.Flush()
			lastSnapshot = snapshot
		}
		select {
		case <-ctx.Done():
			return
		case <-poll.C:
		case <-keepalive.C:
			if _, err := fmt.Fprint(w, ": keepalive\n\n"); err != nil {
				return
			}
			flusher.Flush()
		}
	}
}

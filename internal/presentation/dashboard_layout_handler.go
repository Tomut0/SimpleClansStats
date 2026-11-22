package presentation

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Tomut0/SimpleClansStats/internal/application"
	"github.com/Tomut0/SimpleClansStats/internal/domain"
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
	ctx := r.Context()
	if status := s.adminStatus(ctx); status != http.StatusOK {
		http.Error(w, http.StatusText(status), status)
		return
	}
	userID, ok := s.sessions.Get(ctx, "user_id").(int64)
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

	if err := s.dashboard.SaveLayout(ctx, userID, layout); err != nil {
		if err == application.ErrInvalidDashboardLayout {
			http.Error(w, "invalid layout", http.StatusBadRequest)
		} else {
			http.Error(w, "could not save layout", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

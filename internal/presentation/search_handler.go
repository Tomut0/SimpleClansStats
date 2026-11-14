package presentation

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/google/uuid"
)

func (s *ServerHandler) Search(w http.ResponseWriter, r *http.Request) {
	serverID := r.URL.Query().Get("serverId")
	if _, err := uuid.Parse(serverID); err != nil {
		http.Error(w, "serverId is required", http.StatusBadRequest)
		return
	}
	if s.search == nil {
		http.Error(w, "search unavailable", http.StatusServiceUnavailable)
		return
	}
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query == "" {
		writeIngestJSON(w, http.StatusOK, []domain.SearchResult{})
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	results, err := s.search.Search(r.Context(), serverID, query, limit)
	if err != nil {
		http.Error(w, "could not search", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Cache-Control", "no-store")
	writeIngestJSON(w, http.StatusOK, results)
}

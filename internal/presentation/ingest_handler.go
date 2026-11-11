package presentation

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime"
	"net/http"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/application"
)

const maxIngestBytes = 8 << 20

func (s *ServerHandler) Ingest(w http.ResponseWriter, r *http.Request) {
	keys := r.Header.Values("X-API-Key")
	if len(keys) != 1 || keys[0] == "" {
		ingestError(w, http.StatusUnauthorized, "unauthorized", "Invalid collector key")
		return
	}
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/json" {
		ingestError(w, http.StatusUnsupportedMediaType, "unsupported_media_type", "Expected application/json")
		return
	}
	body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, maxIngestBytes))
	var limit *http.MaxBytesError
	if errors.As(err, &limit) {
		ingestError(w, http.StatusRequestEntityTooLarge, "snapshot_too_large", "Snapshot exceeds 8 MiB")
		return
	}
	if err != nil {
		ingestError(w, http.StatusBadRequest, "invalid_snapshot", "Invalid snapshot")
		return
	}
	if s.ingest == nil {
		ingestError(w, http.StatusInternalServerError, "internal_error", "Ingest unavailable")
		return
	}
	now := time.Now().UTC()
	err = s.ingest.Ingest(r.Context(), keys[0], body, now)
	if errors.Is(err, application.ErrInvalidIngest) {
		ingestError(w, http.StatusBadRequest, "invalid_snapshot", "Invalid snapshot state or events")
		return
	}
	if errors.Is(err, application.ErrUnauthorizedIngest) {
		ingestError(w, http.StatusUnauthorized, "unauthorized", "Invalid collector key")
		return
	}
	if err != nil {
		log.Printf("save ingest snapshot: %v", err)
		ingestError(w, http.StatusInternalServerError, "internal_error", "Could not save snapshot")
		return
	}
	writeIngestJSON(w, http.StatusAccepted, struct {
		Accepted    bool      `json:"accepted"`
		ProcessedAt time.Time `json:"processedAt"`
	}{true, now})
}

func ingestError(w http.ResponseWriter, status int, code, message string) {
	writeIngestJSON(w, status, Error{Code: code, Message: message})
}

func writeIngestJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

package presentation

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime"
	"net/http"
	"strings"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite/generated"
	"github.com/google/uuid"
)

const maxIngestBytes = 8 << 20

type ingestSnapshot struct {
	TakenAt  time.Time `json:"takenAt"`
	ServerID string    `json:"serverId"`
	Clans    []struct {
		ExternalID string `json:"externalId"`
		Tag        string `json:"tag"`
		Name       string `json:"name"`
	} `json:"clans"`
	Players []struct {
		UUID uuid.UUID `json:"uuid"`
		Name string    `json:"name"`
	} `json:"players"`
}

func (s *ServerHandler) storeIngest(ctx context.Context, serverID, takenAt, receivedAt, payload string, state []byte, events []ingestEvent) error {
	hash := sha256.Sum256(state)
	stateHash := hex.EncodeToString(hash[:])
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	queries := s.settings.WithTx(tx)
	updated, err := queries.RecordIngestReceipt(ctx, generated.RecordIngestReceiptParams{
		ServerID: serverID, ReceivedAt: receivedAt,
	})
	if err != nil {
		return err
	}
	if updated == 0 {
		return sql.ErrNoRows
	}

	for _, event := range events {
		if err := queries.SaveIngestEvent(ctx, generated.SaveIngestEventParams{
			ServerID: serverID, EventID: event.id, ReceivedAt: receivedAt, Payload: event.payload,
		}); err != nil {
			return err
		}
	}
	if err := queries.SaveIngestSnapshot(ctx, generated.SaveIngestSnapshotParams{
		ServerID: serverID, TakenAt: takenAt, ReceivedAt: receivedAt, Payload: payload, StateHash: stateHash,
	}); err != nil {
		return err
	}
	if err := queries.UpdateIngestState(ctx, generated.UpdateIngestStateParams{
		ServerID: serverID, TakenAt: takenAt, StateHash: stateHash, Payload: string(state),
	}); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *ServerHandler) Ingest(w http.ResponseWriter, r *http.Request) {
	keys := r.Header.Values("X-API-Key")
	if len(keys) != 1 || keys[0] == "" {
		ingestError(w, http.StatusUnauthorized, "unauthorized", "Invalid collector key")
		return
	}
	hash := sha256.Sum256([]byte(keys[0]))
	serverID, err := s.settings.GetActiveServerIDByKeyHash(r.Context(), hash[:])
	if errors.Is(err, sql.ErrNoRows) {
		ingestError(w, http.StatusUnauthorized, "unauthorized", "Invalid collector key")
		return
	}
	if err != nil {
		log.Printf("authorize ingest: %v", err)
		ingestError(w, http.StatusInternalServerError, "internal_error", "Could not authorize ingest")
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
	var snapshot ingestSnapshot
	if err != nil || json.Unmarshal(body, &snapshot) != nil || !snapshot.valid(serverID) {
		ingestError(w, http.StatusBadRequest, "invalid_snapshot", "Invalid snapshot or server ID")
		return
	}

	state, events, err := ingestState(body)
	if err != nil {
		ingestError(w, http.StatusBadRequest, "invalid_snapshot", "Invalid snapshot state or events")
		return
	}
	now := time.Now().UTC()
	err = s.storeIngest(r.Context(), serverID, ingestTime(snapshot.TakenAt), ingestTime(now), string(body), state, events)
	if errors.Is(err, sql.ErrNoRows) {
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

func (s ingestSnapshot) valid(serverID string) bool {
	if s.TakenAt.IsZero() || s.Clans == nil || s.Players == nil || (s.ServerID != "" && !strings.EqualFold(s.ServerID, serverID)) {
		return false
	}
	for _, clan := range s.Clans {
		if strings.TrimSpace(clan.ExternalID) == "" || strings.TrimSpace(clan.Tag) == "" || strings.TrimSpace(clan.Name) == "" {
			return false
		}
	}
	for _, player := range s.Players {
		if player.UUID == uuid.Nil || strings.TrimSpace(player.Name) == "" {
			return false
		}
	}
	return true
}

func ingestTime(t time.Time) string {
	// Fixed-width UTC timestamps sort chronologically as SQLite TEXT.
	return t.UTC().Format("2006-01-02T15:04:05.000000000Z")
}

func ingestError(w http.ResponseWriter, status int, code, message string) {
	writeIngestJSON(w, status, Error{Code: code, Message: message})
}

func writeIngestJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

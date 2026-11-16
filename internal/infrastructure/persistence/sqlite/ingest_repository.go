package sqlite

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite/generated"
)

type ingestRepository struct {
	db      *sql.DB
	queries *generated.Queries
}

func NewSQLiteIngestRepository(db *sql.DB) ports.IngestRepository {
	return &ingestRepository{db: db, queries: generated.New(db)}
}

func (r *ingestRepository) GetServerIDByKeyHash(ctx context.Context, hash []byte) (string, error) {
	return r.queries.GetActiveServerIDByKeyHash(ctx, hash)
}

func (r *ingestRepository) SaveIngest(ctx context.Context, serverID string, takenAt, receivedAt time.Time, payload string, state []byte, events []domain.IngestEvent) error {
	hash := sha256.Sum256(state)
	taken := FormatIngestTime(takenAt)
	stateHash := hex.EncodeToString(hash[:])
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	queries := r.queries.WithTx(tx)
	received := FormatIngestTime(receivedAt)
	updated, err := queries.RecordIngestReceipt(ctx, generated.RecordIngestReceiptParams{ServerID: serverID, ReceivedAt: received})
	if err != nil {
		return err
	}
	if updated == 0 {
		return sql.ErrNoRows
	}
	var currentHash, currentTaken string
	if err := tx.QueryRowContext(ctx, `SELECT state_hash, last_taken_at FROM servers WHERE id = ?`, serverID).Scan(&currentHash, &currentTaken); err != nil {
		return err
	}
	stateChanged := taken > currentTaken && stateHash != currentHash
	for _, event := range events {
		if err := queries.SaveIngestEvent(ctx, generated.SaveIngestEventParams{ServerID: serverID, EventID: event.ID, ReceivedAt: received, Payload: event.Payload}); err != nil {
			return err
		}
	}
	if err := queries.SaveIngestSnapshot(ctx, generated.SaveIngestSnapshotParams{ServerID: serverID, TakenAt: taken, ReceivedAt: received, Payload: payload, StateHash: stateHash}); err != nil {
		return err
	}
	if err := queries.UpdateIngestState(ctx, generated.UpdateIngestStateParams{ServerID: serverID, TakenAt: taken, StateHash: stateHash, Payload: string(state)}); err != nil {
		return err
	}
	if stateChanged {
		if err := indexSearchDocuments(ctx, tx, serverID, payload); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func FormatIngestTime(t time.Time) string { return t.UTC().Format("2006-01-02T15:04:05.000000000Z") }

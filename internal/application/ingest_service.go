package application

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
)

var ErrInvalidIngest = errors.New("invalid ingest")
var ErrUnauthorizedIngest = errors.New("invalid collector key")

type IngestService struct{ repository ports.IngestRepository }

func NewIngestService(repository ports.IngestRepository) *IngestService {
	return &IngestService{repository: repository}
}

func (s *IngestService) Ingest(ctx context.Context, key string, payload []byte, receivedAt time.Time) error {
	hash := sha256.Sum256([]byte(key))
	serverID, err := s.repository.GetServerIDByKeyHash(ctx, hash[:])
	if errors.Is(err, sql.ErrNoRows) {
		return ErrUnauthorizedIngest
	}
	if err != nil {
		return err
	}
	var snapshot domain.IngestSnapshot
	if json.Unmarshal(payload, &snapshot) != nil || !snapshot.Valid(serverID) {
		return ErrInvalidIngest
	}
	state, events, err := NormalizeIngestState(payload)
	if err != nil {
		return ErrInvalidIngest
	}
	err = s.repository.SaveIngest(ctx, serverID, snapshot.TakenAt, receivedAt, string(payload), state, events)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrUnauthorizedIngest
	}
	return err
}

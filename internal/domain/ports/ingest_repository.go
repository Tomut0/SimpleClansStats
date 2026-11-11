package ports

import (
	"context"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
)

type IngestRepository interface {
	GetServerIDByKeyHash(ctx context.Context, hash []byte) (string, error)
	SaveIngest(ctx context.Context, serverID string, takenAt, receivedAt time.Time, payload string, state []byte, events []domain.IngestEvent) error
}

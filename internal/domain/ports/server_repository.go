package ports

import (
	"context"
	"github.com/Tomut0/SimpleClansStats/internal/domain"
)

type ServerRepository interface {
	ListServers(ctx context.Context) ([]domain.ServerRecord, error)
	CreateServer(ctx context.Context, id, name string, keyHash []byte) error
	DeleteServer(ctx context.Context, id string) error
}

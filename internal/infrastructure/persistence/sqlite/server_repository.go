package sqlite

import (
	"context"
	"database/sql"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite/generated"
)

type serverRepository struct{ queries *generated.Queries }

func NewSQLiteServerRepository(db *sql.DB) ports.ServerRepository {
	return &serverRepository{queries: generated.New(db)}
}

func (r *serverRepository) ListServers(ctx context.Context) ([]domain.ServerRecord, error) {
	rows, err := r.queries.ListServers(ctx)
	if err != nil {
		return nil, err
	}
	servers := make([]domain.ServerRecord, 0, len(rows))
	for _, row := range rows {
		servers = append(servers, domain.ServerRecord{ID: row.ID, DisplayName: row.DisplayName, Disabled: row.Disabled, LastReceivedAt: row.LastReceivedAt, OnlineTimeoutMinutes: row.ServerOnlineTimeoutMinutes})
	}
	return servers, nil
}

func (r *serverRepository) CreateServer(ctx context.Context, id, name string, keyHash []byte) error {
	return r.queries.CreateServer(ctx, generated.CreateServerParams{ID: id, DisplayName: name, ApiKeyHash: keyHash})
}

func (r *serverRepository) DeleteServer(ctx context.Context, id string) error {
	return r.queries.DeleteServer(ctx, id)
}

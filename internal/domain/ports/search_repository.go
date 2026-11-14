package ports

import (
	"context"
	"github.com/Tomut0/SimpleClansStats/internal/domain"
)

type SearchRepository interface {
	Search(ctx context.Context, serverID, query string, limit int) (domain.SearchData, error)
}

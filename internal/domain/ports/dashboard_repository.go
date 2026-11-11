package ports

import "context"

import "github.com/Tomut0/SimpleClansStats/internal/domain"

type DashboardRepository interface {
	GetDashboardData(ctx context.Context, serverID string) (domain.DashboardData, error)
	GetLayout(ctx context.Context, userID int64) (domain.DashboardLayout, error)
	SaveLayout(ctx context.Context, userID int64, layout domain.DashboardLayout) error
}

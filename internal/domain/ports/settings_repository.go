package ports

import (
	"context"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
)

type SettingsRepository interface {
	GetAppSettings(ctx context.Context) (domain.AppSettings, error)
	GetAppLogo(ctx context.Context) ([]byte, error)
	SaveAppSettings(ctx context.Context, update domain.AppSettingsUpdate) (domain.AppSettings, error)
	SaveAppTheme(ctx context.Context, preset string) error
}

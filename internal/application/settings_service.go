package application

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"
	_ "time/tzdata"
	"unicode/utf8"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
)

type SettingsService struct{ repository ports.SettingsRepository }

var ErrInvalidSettings = errors.New("invalid application settings")

func NewSettingsService(repository ports.SettingsRepository) *SettingsService {
	return &SettingsService{repository: repository}
}

func (s *SettingsService) Get(ctx context.Context) (domain.AppSettings, error) {
	return s.repository.GetAppSettings(ctx)
}

func (s *SettingsService) Logo(ctx context.Context) ([]byte, error) {
	return s.repository.GetAppLogo(ctx)
}

func (s *SettingsService) Update(ctx context.Context, update domain.AppSettingsUpdate) (domain.AppSettings, error) {
	update.AppName = strings.TrimSpace(update.AppName)
	update.PlayerAvatarURL = strings.TrimSpace(update.PlayerAvatarURL)
	if update.AppName == "" || !utf8.ValidString(update.AppName) || utf8.RuneCountInString(update.AppName) > 100 ||
		(update.DateFormat != "DD.MM.YYYY" && update.DateFormat != "YYYY-MM-DD" && update.DateFormat != "MM/DD/YYYY") ||
		(update.TimeFormat != "24h" && update.TimeFormat != "12h") || update.TimeZone == "" || update.TimeZone == "Local" ||
		!validPlayerAvatarURL(update.PlayerAvatarURL) {
		return domain.AppSettings{}, ErrInvalidSettings
	}
	if _, err := time.LoadLocation(update.TimeZone); err != nil {
		return domain.AppSettings{}, ErrInvalidSettings
	}
	if update.ServerOnlineTimeoutMinutes != nil && (*update.ServerOnlineTimeoutMinutes < 1 || *update.ServerOnlineTimeoutMinutes > 1440) {
		return domain.AppSettings{}, ErrInvalidSettings
	}
	if update.IngestRetentionDays != nil && (*update.IngestRetentionDays < 0 || *update.IngestRetentionDays > 36500) {
		return domain.AppSettings{}, ErrInvalidSettings
	}
	return s.repository.SaveAppSettings(ctx, update)
}

func validPlayerAvatarURL(value string) bool {
	if len(value) > 500 || strings.Count(value, "{username}") != 1 {
		return false
	}
	parsed, err := url.Parse(strings.ReplaceAll(value, "{username}", "player"))
	return err == nil && parsed.Scheme == "https" && parsed.Host != "" && parsed.User == nil
}

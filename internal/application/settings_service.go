package application

import (
	"context"
	"errors"
	"net/url"
	"regexp"
	"slices"
	"strings"
	"sync"
	"time"
	_ "time/tzdata"
	"unicode/utf8"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
)

type SettingsService struct {
	repository ports.SettingsRepository
	themeMu    sync.Mutex
}

var localePattern = regexp.MustCompile(`^[A-Za-z]{2,3}(-[A-Za-z0-9]{2,8})*$`)
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
	if !validSettingsUpdate(update) {
		return domain.AppSettings{}, ErrInvalidSettings
	}
	return s.repository.SaveAppSettings(ctx, update)
}

func validSettingsUpdate(update domain.AppSettingsUpdate) bool {
	validAppName := update.AppName != "" && utf8.ValidString(update.AppName) && utf8.RuneCountInString(update.AppName) <= 100
	validLanguage := len(update.PrimaryLanguage) <= 35 && localePattern.MatchString(update.PrimaryLanguage)
	if !validAppName || !validLanguage {
		return false
	}

	validDateFormat := slices.Contains([]string{"DD.MM.YYYY", "YYYY-MM-DD", "MM/DD/YYYY"}, update.DateFormat)
	validTimeFormat := update.TimeFormat == "24h" || update.TimeFormat == "12h"
	if !validDateFormat || !validTimeFormat || update.TimeZone == "" || update.TimeZone == "Local" {
		return false
	}
	if !validPlayerAvatarURL(update.PlayerAvatarURL) {
		return false
	}
	if _, err := time.LoadLocation(update.TimeZone); err != nil {
		return false
	}

	return optionalSettingInRange(update.ServerOnlineTimeoutMinutes, 1, 1440) &&
		optionalSettingInRange(update.IngestRetentionDays, 0, 36500)
}

func optionalSettingInRange(value *int64, min, max int64) bool {
	return value == nil || (*value >= min && *value <= max)
}

func validPlayerAvatarURL(value string) bool {
	if len(value) > 500 || strings.Count(value, "{username}") != 1 {
		return false
	}
	parsed, err := url.Parse(strings.ReplaceAll(value, "{username}", "player"))
	return err == nil && parsed.Scheme == "https" && parsed.Host != "" && parsed.User == nil
}

package application

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
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

const DefaultThemePreset = "b2pklPa1A"

var themePresetPattern = regexp.MustCompile(`^[A-Za-z0-9]{6,32}$`)
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

func (s *SettingsService) ThemeCSS(ctx context.Context) (string, error) {
	s.themeMu.Lock()
	defer s.themeMu.Unlock()
	cssPath := themeCSSPath()
	css, err := os.ReadFile(cssPath)
	if err == nil {
		return string(css), nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	settings, err := s.repository.GetAppSettings(ctx)
	if err != nil {
		return "", err
	}
	if settings.Theme == DefaultThemePreset {
		return "", nil
	}
	generatedCSS, _, err := runThemeCLI(ctx, settings.Theme)
	if err != nil {
		return "", err
	}
	if err := writeThemeCSS(generatedCSS); err != nil {
		return "", err
	}
	return generatedCSS, nil
}

func (s *SettingsService) ApplyTheme(ctx context.Context, preset string) error {
	if !themePresetPattern.MatchString(preset) {
		return ErrInvalidSettings
	}
	s.themeMu.Lock()
	defer s.themeMu.Unlock()
	css, restore, err := runThemeCLI(ctx, preset)
	if err != nil {
		return err
	}
	if err := s.repository.SaveAppTheme(ctx, preset); err != nil {
		restore()
		return fmt.Errorf("save applied theme: %w", err)
	}
	if err := writeThemeCSS(css); err != nil {
		_ = os.Remove(themeCSSPath())
		return fmt.Errorf("save applied theme stylesheet: %w", err)
	}
	return nil
}

func themeCSSPath() string {
	if path := os.Getenv("APP_THEME_CSS_PATH"); path != "" {
		return path
	}
	return filepath.Join("data", "theme.css")
}

func writeThemeCSS(css string) error {
	path := themeCSSPath()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	temp, err := os.CreateTemp(filepath.Dir(path), ".theme-*.css")
	if err != nil {
		return err
	}
	tempPath := temp.Name()
	defer os.Remove(tempPath)
	if _, err = temp.WriteString(css); err != nil {
		_ = temp.Close()
		return err
	}
	if err = temp.Close(); err != nil {
		return err
	}
	return os.Rename(tempPath, path)
}

func runThemeCLI(ctx context.Context, preset string) (string, func(), error) {
	if !themePresetPattern.MatchString(preset) {
		return "", func() {}, ErrInvalidSettings
	}
	projectDir := os.Getenv("SHADCN_PROJECT_DIR")
	if projectDir == "" {
		projectDir = "web"
	}
	cssPath := filepath.Join(projectDir, "src", "routes", "layout.css")
	originalCSS, err := os.ReadFile(cssPath)
	if err != nil {
		return "", func() {}, fmt.Errorf("read theme source: %w", err)
	}
	restore := func() { _ = os.WriteFile(cssPath, originalCSS, 0o644) }
	cmd := exec.CommandContext(ctx, "npx", "--yes", "shadcn-svelte@latest", "apply", preset, "--only", "theme,font", "--yes")
	cmd.Dir = projectDir
	if output, err := cmd.CombinedOutput(); err != nil {
		restore()
		return "", func() {}, fmt.Errorf("apply shadcn preset: %w: %s", err, strings.TrimSpace(string(output)))
	}
	css, err := os.ReadFile(cssPath)
	if err != nil {
		restore()
		return "", func() {}, fmt.Errorf("read applied theme: %w", err)
	}
	// Browsers do not process Tailwind's build-time @theme directive.
	return strings.ReplaceAll(string(css), "@theme inline {", ":root {"), restore, nil
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

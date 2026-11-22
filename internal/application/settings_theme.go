package application

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const DefaultThemePreset = "b2pklPa1A"

var themePresetPattern = regexp.MustCompile(`^[A-Za-z0-9]{6,32}$`)
var themeImportPattern = regexp.MustCompile(`(?m)^@import[^;]+;\r?\n?`)

func (s *SettingsService) ThemeCSS(ctx context.Context) (string, error) {
	s.themeMu.Lock()
	defer s.themeMu.Unlock()

	css, err := os.ReadFile(themeCSSPath())
	if err == nil {
		return cleanThemeCSS(string(css)), nil
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

	generatedCSS, _, err := buildThemeCSS(ctx, settings.Theme)
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

	css, restoreSource, err := buildThemeCSS(ctx, preset)
	if err != nil {
		return err
	}
	if err := s.repository.SaveAppTheme(ctx, preset); err != nil {
		restoreSource()
		return fmt.Errorf("save applied theme: %w", err)
	}
	if err := writeThemeCSS(css); err != nil {
		// A missing stylesheet lets ThemeCSS rebuild it from the saved preset.
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
	directory := filepath.Dir(path)
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return err
	}
	temp, err := os.CreateTemp(directory, ".theme-*.css")
	if err != nil {
		return err
	}
	tempPath := temp.Name()
	defer os.Remove(tempPath)

	if _, err := temp.WriteString(css); err != nil {
		_ = temp.Close()
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	return os.Rename(tempPath, path)
}

func buildThemeCSS(ctx context.Context, preset string) (string, func(), error) {
	if !themePresetPattern.MatchString(preset) {
		return "", nil, ErrInvalidSettings
	}
	projectDir := os.Getenv("SHADCN_PROJECT_DIR")
	if projectDir == "" {
		projectDir = "web"
	}
	cssPath := filepath.Join(projectDir, "src", "routes", "layout.css")
	originalCSS, err := os.ReadFile(cssPath)
	if err != nil {
		return "", nil, fmt.Errorf("read theme source: %w", err)
	}
	restoreSource := func() { _ = os.WriteFile(cssPath, originalCSS, 0o644) }

	apply := exec.CommandContext(ctx, "npx", "--yes", "shadcn-svelte@latest", "apply", preset, "--yes")
	apply.Dir = projectDir
	if output, err := apply.CombinedOutput(); err != nil {
		restoreSource()
		return "", nil, fmt.Errorf("apply shadcn preset: %w: %s", err, strings.TrimSpace(string(output)))
	}

	build := exec.CommandContext(ctx, "npm", "run", "build")
	build.Dir = projectDir
	if output, err := build.CombinedOutput(); err != nil {
		restoreSource()
		return "", nil, fmt.Errorf("build frontend after applying theme: %w: %s", err, strings.TrimSpace(string(output)))
	}

	css, err := os.ReadFile(cssPath)
	if err != nil {
		restoreSource()
		return "", nil, fmt.Errorf("read applied theme: %w", err)
	}
	return cleanThemeCSS(string(css)), restoreSource, nil
}

func cleanThemeCSS(css string) string {
	// The saved stylesheet is served directly to browsers.
	css = strings.ReplaceAll(css, "@theme inline {", ":root {")
	return themeImportPattern.ReplaceAllString(css, "")
}

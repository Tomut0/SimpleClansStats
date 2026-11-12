package sqlite

import (
	"context"
	"database/sql"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite/generated"
)

type settingsRepository struct {
	db      *sql.DB
	queries *generated.Queries
}

func NewSQLiteSettingsRepository(db *sql.DB) ports.SettingsRepository {
	return &settingsRepository{db: db, queries: generated.New(db)}
}

func (r *settingsRepository) GetAppSettings(ctx context.Context) (domain.AppSettings, error) {
	row, err := r.queries.GetAppSettings(ctx)
	return settingsDomain(row), err
}

func (r *settingsRepository) GetAppLogo(ctx context.Context) ([]byte, error) {
	return r.queries.GetAppLogo(ctx)
}

func (r *settingsRepository) SaveAppTheme(ctx context.Context, preset string) error {
	_, err := r.db.ExecContext(ctx, `UPDATE app_settings SET theme = ? WHERE id = 1`, preset)
	return err
}

func (r *settingsRepository) SaveAppSettings(ctx context.Context, update domain.AppSettingsUpdate) (domain.AppSettings, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.AppSettings{}, err
	}
	defer tx.Rollback()
	queries := r.queries.WithTx(tx)
	if err = queries.UpdateAppSettings(ctx, generated.UpdateAppSettingsParams{
		AppName: update.AppName, RemoveAuthor: update.RemoveAuthor, RemoveTelemetry: update.RemoveTelemetry, PrimaryLanguage: update.PrimaryLanguage, Multilingual: update.Multilingual, TimeZone: update.TimeZone,
		DateFormat: update.DateFormat, TimeFormat: update.TimeFormat, PlayerAvatarUrl: update.PlayerAvatarURL, Logo: update.Logo,
	}); err != nil {
		return domain.AppSettings{}, err
	}
	if err = queries.UpdateIngestSettings(ctx, generated.UpdateIngestSettingsParams{
		ServerOnlineTimeoutMinutes: nullInt(update.ServerOnlineTimeoutMinutes), IngestRetentionDays: nullInt(update.IngestRetentionDays),
	}); err != nil {
		return domain.AppSettings{}, err
	}
	row, err := queries.GetAppSettings(ctx)
	if err != nil {
		return domain.AppSettings{}, err
	}
	if err = tx.Commit(); err != nil {
		return domain.AppSettings{}, err
	}
	return settingsDomain(row), nil
}

func nullInt(value *int64) sql.NullInt64 {
	if value == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: *value, Valid: true}
}

func settingsDomain(row generated.GetAppSettingsRow) domain.AppSettings {
	return domain.AppSettings{AppName: row.AppName, RemoveAuthor: row.RemoveAuthor, RemoveTelemetry: row.RemoveTelemetry, TimeZone: row.TimeZone,
		DateFormat: row.DateFormat, TimeFormat: row.TimeFormat, PlayerAvatarURL: row.PlayerAvatarUrl, PrimaryLanguage: row.PrimaryLanguage, Multilingual: row.Multilingual, LogoVersion: row.LogoVersion,
		ServerOnlineTimeoutMinutes: row.ServerOnlineTimeoutMinutes, IngestRetentionDays: row.IngestRetentionDays, Theme: row.Theme}
}

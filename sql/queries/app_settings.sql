-- name: GetAppSettings :one
SELECT app_name, remove_author, remove_telemetry, time_zone, date_format, time_format, player_avatar_url, logo_version,
       server_online_timeout_minutes, ingest_retention_days, theme
FROM app_settings CROSS JOIN ingest_settings WHERE app_settings.id = 1 AND ingest_settings.id = 1;

-- name: UpdateAppSettings :exec
UPDATE app_settings
SET app_name = sqlc.arg(app_name),
    remove_author = sqlc.arg(remove_author),
    remove_telemetry = sqlc.arg(remove_telemetry),
    time_zone = sqlc.arg(time_zone),
    date_format = sqlc.arg(date_format),
    time_format = sqlc.arg(time_format),
    player_avatar_url = sqlc.arg(player_avatar_url),
    logo = COALESCE(sqlc.narg(logo), logo),
    logo_version = logo_version + (sqlc.narg(logo) IS NOT NULL)
WHERE id = 1;

-- name: UpdateIngestSettings :exec
UPDATE ingest_settings
SET server_online_timeout_minutes = COALESCE(sqlc.narg(server_online_timeout_minutes), server_online_timeout_minutes),
    ingest_retention_days = COALESCE(sqlc.narg(ingest_retention_days), ingest_retention_days)
WHERE id = 1;

-- name: GetIngestRetentionDays :one
SELECT ingest_retention_days FROM ingest_settings WHERE id = 1;

-- name: GetAppLogo :one
SELECT logo FROM app_settings WHERE id = 1;

-- name: GetAppSettings :one
SELECT app_name, remove_author, time_zone, date_format, time_format, logo_version
FROM app_settings WHERE id = 1;

-- name: UpdateAppSettings :one
UPDATE app_settings
SET app_name = sqlc.arg(app_name),
    remove_author = sqlc.arg(remove_author),
    time_zone = sqlc.arg(time_zone),
    date_format = sqlc.arg(date_format),
    time_format = sqlc.arg(time_format),
    logo = COALESCE(sqlc.narg(logo), logo),
    logo_version = logo_version + (sqlc.narg(logo) IS NOT NULL)
WHERE id = 1
RETURNING app_name, remove_author, time_zone, date_format, time_format, logo_version;

-- name: GetAppLogo :one
SELECT logo FROM app_settings WHERE id = 1;

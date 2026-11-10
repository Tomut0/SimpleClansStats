-- name: CreateServer :exec
INSERT INTO servers (id, display_name, api_key_hash) VALUES (?, ?, ?);

-- name: DeleteServer :exec
DELETE FROM servers WHERE id = ?;

-- name: GetActiveServerIDByKeyHash :one
SELECT id FROM servers WHERE api_key_hash = ? AND NOT disabled;

-- name: ListServers :many
SELECT servers.id, display_name, disabled, last_received_at, server_online_timeout_minutes
FROM servers CROSS JOIN ingest_settings
WHERE ingest_settings.id = 1
ORDER BY display_name COLLATE NOCASE, servers.id;

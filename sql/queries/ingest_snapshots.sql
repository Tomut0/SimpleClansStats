-- name: SaveIngestSnapshot :exec
INSERT INTO ingest_snapshots (server_id, taken_at, received_at, payload)
SELECT id, sqlc.arg(taken_at), sqlc.arg(received_at), sqlc.arg(payload)
FROM servers
WHERE id = sqlc.arg(server_id)
  AND last_taken_at < sqlc.arg(taken_at) AND state_hash <> sqlc.arg(state_hash);

-- name: RecordIngestReceipt :execrows
UPDATE servers SET last_received_at = MAX(last_received_at, CAST(sqlc.arg(received_at) AS TEXT))
WHERE id = sqlc.arg(server_id) AND NOT disabled;

-- name: UpdateIngestState :exec
UPDATE servers
SET last_taken_at = sqlc.arg(taken_at), state_hash = sqlc.arg(state_hash),
    state_payload = CASE WHEN state_hash <> sqlc.arg(state_hash) THEN sqlc.arg(payload) ELSE state_payload END
WHERE id = sqlc.arg(server_id) AND last_taken_at < sqlc.arg(taken_at);

-- name: SaveIngestEvent :exec
INSERT INTO ingest_events (server_id, event_id, received_at, payload)
VALUES (?, ?, ?, ?)
ON CONFLICT (server_id, event_id) DO NOTHING;

-- name: PurgeIngestSnapshots :exec
DELETE FROM ingest_snapshots WHERE received_at < ?;

-- name: PurgeIngestEvents :exec
DELETE FROM ingest_events WHERE received_at < ?;

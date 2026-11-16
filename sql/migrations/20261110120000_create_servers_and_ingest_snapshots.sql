-- +goose Up
CREATE TABLE ingest_settings (
  id INTEGER PRIMARY KEY CHECK (id = 1),
  server_online_timeout_minutes INTEGER NOT NULL DEFAULT 10 CHECK (server_online_timeout_minutes BETWEEN 1 AND 1440),
  ingest_retention_days INTEGER NOT NULL DEFAULT 30 CHECK (ingest_retention_days BETWEEN 0 AND 36500)
);
INSERT INTO ingest_settings (id) VALUES (1);

CREATE TABLE servers (
  id TEXT PRIMARY KEY NOT NULL,
  display_name TEXT NOT NULL,
  api_key_hash BLOB NOT NULL UNIQUE CHECK (length(api_key_hash) = 32),
  disabled BOOLEAN NOT NULL DEFAULT FALSE,
  last_received_at TEXT NOT NULL DEFAULT '',
  last_taken_at TEXT NOT NULL DEFAULT '',
  state_hash TEXT NOT NULL DEFAULT '',
  state_payload TEXT NOT NULL DEFAULT ''
);

CREATE TABLE ingest_snapshots (
  server_id TEXT NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
  taken_at TEXT NOT NULL,
  received_at TEXT NOT NULL,
  payload TEXT NOT NULL CHECK (json_valid(payload)),
  PRIMARY KEY (server_id, taken_at)
);

CREATE TABLE ingest_events (
  server_id TEXT NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
  event_id TEXT NOT NULL,
  received_at TEXT NOT NULL,
  payload TEXT NOT NULL CHECK (json_valid(payload)),
  PRIMARY KEY (server_id, event_id)
);

CREATE INDEX ingest_snapshots_received_at ON ingest_snapshots (received_at);
CREATE INDEX ingest_events_received_at ON ingest_events (received_at);

-- +goose Down
DROP TABLE IF EXISTS ingest_events;
DROP TABLE ingest_snapshots;
DROP TABLE servers;
DROP TABLE IF EXISTS ingest_settings;

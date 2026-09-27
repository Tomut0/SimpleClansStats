-- +goose Up
CREATE TABLE IF NOT EXISTS sessions (
  id               TEXT PRIMARY KEY,
  user_id          INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  token_hash       BLOB NOT NULL UNIQUE,
  csrf_token       TEXT NOT NULL,
  user_agent       TEXT,
  ip               TEXT,
  created_at       TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now')),
  last_activity_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now')),
  expires_at       TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS sessions;

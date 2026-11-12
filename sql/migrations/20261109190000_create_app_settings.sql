-- +goose Up
CREATE TABLE app_settings (
  id INTEGER PRIMARY KEY CHECK (id = 1),
  app_name TEXT NOT NULL DEFAULT 'SCStats',
  remove_author BOOLEAN NOT NULL DEFAULT FALSE,
  remove_telemetry BOOLEAN NOT NULL DEFAULT FALSE,
  installation_id TEXT NOT NULL DEFAULT '',
  time_zone TEXT NOT NULL DEFAULT 'Europe/Moscow',
  date_format TEXT NOT NULL DEFAULT 'DD.MM.YYYY',
  time_format TEXT NOT NULL DEFAULT '24h',
  player_avatar_url TEXT NOT NULL DEFAULT 'https://mc-heads.net/avatar/{username}/32',
  logo BLOB,
  logo_version INTEGER NOT NULL DEFAULT 0,
  theme TEXT NOT NULL DEFAULT 'b2pklPa1A',
  primary_language TEXT NOT NULL DEFAULT 'en',
  multilingual BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO app_settings (id) VALUES (1);

-- +goose Down
DROP TABLE app_settings;

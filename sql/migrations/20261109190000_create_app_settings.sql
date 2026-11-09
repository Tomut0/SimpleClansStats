-- +goose Up
CREATE TABLE app_settings (
  id INTEGER PRIMARY KEY CHECK (id = 1),
  app_name TEXT NOT NULL DEFAULT 'SCStats',
  remove_author BOOLEAN NOT NULL DEFAULT FALSE,
  time_zone TEXT NOT NULL DEFAULT 'Europe/Moscow',
  date_format TEXT NOT NULL DEFAULT 'DD.MM.YYYY',
  time_format TEXT NOT NULL DEFAULT '24h',
  logo BLOB,
  logo_version INTEGER NOT NULL DEFAULT 0
);

INSERT INTO app_settings (id) VALUES (1);

-- +goose Down
DROP TABLE app_settings;

-- +goose Up
CREATE TABLE dashboard_layouts (
  user_id INTEGER PRIMARY KEY NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  layout TEXT NOT NULL CHECK (json_valid(layout))
);

-- +goose Down
DROP TABLE IF EXISTS dashboard_layouts;

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite"
)

func TestCleanupUsesStoredRetention(t *testing.T) {
	db := sqlite.New(filepath.Join(t.TempDir(), "cleanup.sqlite"))
	defer db.Close()
	migration, err := os.ReadFile("../sql/migrations/20261110120000_create_servers_and_ingest_snapshots.sql")
	if err != nil {
		t.Fatal(err)
	}
	up, _, _ := strings.Cut(string(migration), "-- +goose Down")
	if _, err := db.Db.Exec(up); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Db.Exec(`
INSERT INTO servers (id, display_name, api_key_hash) VALUES ('s', 'test', zeroblob(32));
INSERT INTO ingest_events (server_id, event_id, received_at, payload)
VALUES ('s', 'old', '2026-11-07T12:00:00.000000000Z', '{}'),
       ('s', 'cutoff', '2026-11-08T12:00:00.000000000Z', '{}'),
       ('s', 'recent', '2026-11-09T12:00:00.000000000Z', '{}');
UPDATE ingest_settings SET ingest_retention_days = 0;`); err != nil {
		t.Fatal(err)
	}
	now := time.Date(2026, 11, 10, 12, 0, 0, 0, time.UTC)
	check := func(want int) {
		t.Helper()
		if err := cleanupIngest(t.Context(), db.Queries, now); err != nil {
			t.Fatal(err)
		}
		var count int
		if err := db.Db.QueryRow("SELECT count(*) FROM ingest_events").Scan(&count); err != nil || count != want {
			t.Fatalf("after cleanup got %d events, want %d (%v)", count, want, err)
		}
	}
	check(3)
	if _, err := db.Db.Exec("UPDATE ingest_settings SET ingest_retention_days = 2"); err != nil {
		t.Fatal(err)
	}
	check(2)
}

package sqlite

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

func EnsureInstallationID(db *sql.DB) (string, error) {
	candidate, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	if _, err := db.ExecContext(context.Background(), `UPDATE app_settings SET installation_id = ? WHERE id = 1 AND installation_id = ''`, candidate.String()); err != nil {
		return "", err
	}
	var id string
	err = db.QueryRowContext(context.Background(), `SELECT installation_id FROM app_settings WHERE id = 1`).Scan(&id)
	return id, err
}

package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
)

type dashboardRepository struct {
	db *sql.DB
}

func NewSQLiteDashboardRepository(db *sql.DB) ports.DashboardRepository {
	return &dashboardRepository{db: db}
}

func (r *dashboardRepository) GetDashboardData(ctx context.Context, serverID string) (domain.DashboardData, error) {
	var payload string
	err := r.db.QueryRowContext(ctx, `SELECT state_payload FROM servers WHERE id = ? AND disabled = FALSE`, serverID).Scan(&payload)
	if err != nil {
		return domain.DashboardData{}, err
	}
	data := domain.DashboardData{State: json.RawMessage(payload), Events: []json.RawMessage{}, History: []domain.DashboardHistoryPoint{}}
	rows, err := r.db.QueryContext(ctx, `SELECT payload FROM ingest_events WHERE server_id = ? ORDER BY received_at DESC LIMIT 5`, serverID)
	if err != nil {
		return data, err
	}
	defer rows.Close()
	for rows.Next() {
		var event string
		if err := rows.Scan(&event); err != nil {
			return data, err
		}
		data.Events = append(data.Events, json.RawMessage(event))
	}
	if err := rows.Err(); err != nil {
		return data, err
	}
	rows.Close()
	snapshots, err := r.db.QueryContext(ctx, `SELECT payload FROM ingest_snapshots WHERE server_id = ? ORDER BY taken_at DESC LIMIT 48`, serverID)
	if err != nil {
		return data, err
	}
	defer snapshots.Close()
	for snapshots.Next() {
		var snapshot string
		if err := snapshots.Scan(&snapshot); err != nil {
			return data, err
		}
		data.Snapshots = append(data.Snapshots, json.RawMessage(snapshot))
	}
	return data, snapshots.Err()
}

func (r *dashboardRepository) GetLayout(ctx context.Context) (domain.DashboardLayout, error) {
	var value string
	err := r.db.QueryRowContext(ctx, `SELECT dl.layout FROM dashboard_layouts dl JOIN users u ON u.id = dl.user_id WHERE u.role = 'admin' ORDER BY dl.rowid DESC LIMIT 1`).Scan(&value)
	if err == sql.ErrNoRows {
		return domain.DashboardLayout{Version: 1, Widgets: []domain.DashboardWidget{}}, nil
	}
	if err != nil {
		return domain.DashboardLayout{}, err
	}
	var layout domain.DashboardLayout
	err = json.Unmarshal([]byte(value), &layout)
	return layout, err
}

func (r *dashboardRepository) SaveLayout(ctx context.Context, userID int64, layout domain.DashboardLayout) error {
	value, err := json.Marshal(layout)
	if err != nil {
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err = tx.ExecContext(ctx, `DELETE FROM dashboard_layouts WHERE user_id IN (SELECT id FROM users WHERE role = 'admin' AND id != ?)`, userID); err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx, `INSERT INTO dashboard_layouts (user_id, layout) VALUES (?, ?) ON CONFLICT(user_id) DO UPDATE SET layout = excluded.layout`, userID, string(value)); err != nil {
		return err
	}
	err = tx.Commit()
	return err
}

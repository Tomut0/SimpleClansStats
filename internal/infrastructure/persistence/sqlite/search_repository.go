package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
)

type searchRepository struct{ db *sql.DB }

func NewSQLiteSearchRepository(db *sql.DB) ports.SearchRepository {
	return &searchRepository{db: db}
}

func (r *searchRepository) Search(ctx context.Context, serverID, query string, limit int) (domain.SearchData, error) {
	data := domain.SearchData{Matches: []domain.SearchMatch{}}
	terms := strings.Fields(query)
	if len(terms) == 0 {
		return data, nil
	}
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return data, err
	}
	defer tx.Rollback()
	var statePayload string
	if err := tx.QueryRowContext(ctx, `SELECT state_payload FROM servers WHERE id = ?`, serverID).Scan(&statePayload); err != nil {
		if err == sql.ErrNoRows {
			return data, nil
		}
		return data, err
	}
	if statePayload != "" {
		var snapshot struct {
			Clans   []json.RawMessage `json:"clans"`
			Players []json.RawMessage `json:"players"`
		}
		if err := json.Unmarshal([]byte(statePayload), &snapshot); err != nil {
			return data, err
		}
		data.Clans, data.Players = snapshot.Clans, snapshot.Players
	}
	for i, term := range terms {
		terms[i] = `"` + strings.ReplaceAll(term, `"`, `""`) + `"*`
	}
	rows, err := tx.QueryContext(ctx, `SELECT entity_type, entity_id, name FROM search_documents WHERE server_id = ? AND search_text MATCH ? ORDER BY rank LIMIT ?`, serverID, strings.Join(terms, " AND "), limit)
	if err != nil {
		return data, err
	}
	defer rows.Close()
	for rows.Next() {
		var match domain.SearchMatch
		if err := rows.Scan(&match.Type, &match.ID, &match.Name); err != nil {
			return data, err
		}
		data.Matches = append(data.Matches, match)
	}
	if err := rows.Err(); err != nil {
		return data, err
	}
	return data, tx.Commit()
}

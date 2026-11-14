package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"
)

func indexSearchDocuments(ctx context.Context, tx *sql.Tx, serverID, payload string) error {
	var snapshot struct {
		Clans   []json.RawMessage `json:"clans"`
		Players []json.RawMessage `json:"players"`
	}
	if err := json.Unmarshal([]byte(payload), &snapshot); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM search_documents WHERE server_id = ?`, serverID); err != nil {
		return err
	}
	insert := `INSERT INTO search_documents (server_id, entity_type, entity_id, name, payload, search_text) VALUES (?, ?, ?, ?, ?, ?)`
	for _, group := range []struct {
		kind string
		rows []json.RawMessage
	}{{"clan", snapshot.Clans}, {"player", snapshot.Players}} {
		for _, raw := range group.rows {
			var fields map[string]any
			if err := json.Unmarshal(raw, &fields); err != nil {
				return err
			}
			idKey, nameKey := "externalId", "tag"
			if group.kind == "player" {
				idKey, nameKey = "uuid", "name"
			}
			id, _ := fields[idKey].(string)
			name, _ := fields[nameKey].(string)
			searchText := []string{name}
			if group.kind == "clan" {
				searchText = append(searchText, stringField(fields, "name"))
			} else {
				searchText = append(searchText, stringField(fields, "clanTag"), stringField(fields, "clanName"))
			}
			if _, err := tx.ExecContext(ctx, insert, serverID, group.kind, id, name, string(raw), strings.Join(searchText, " ")); err != nil {
				return err
			}
		}
	}
	return nil
}

func stringField(fields map[string]any, key string) string {
	value, _ := fields[key].(string)
	return value
}

package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"
)

type searchSnapshot struct {
	Clans   []json.RawMessage `json:"clans"`
	Players []json.RawMessage `json:"players"`
}

func indexSearchDocuments(ctx context.Context, tx *sql.Tx, serverID, payload string) error {
	var snapshot searchSnapshot
	if err := json.Unmarshal([]byte(payload), &snapshot); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM search_documents WHERE server_id = ?`, serverID); err != nil {
		return err
	}
	const insert = `
		INSERT INTO search_documents (server_id, entity_type, entity_id, name, payload, search_text)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	for _, group := range []struct {
		kind        string
		rows        []json.RawMessage
		idField     string
		nameField   string
		extraFields []string
	}{
		{
			kind: "clan", rows: snapshot.Clans, idField: "externalId", nameField: "tag",
			extraFields: []string{"name"},
		},
		{
			kind: "player", rows: snapshot.Players, idField: "uuid", nameField: "name",
			extraFields: []string{"clanTag", "clanName"},
		},
	} {
		for _, raw := range group.rows {
			var fields map[string]any
			if err := json.Unmarshal(raw, &fields); err != nil {
				return err
			}
			id := stringField(fields, group.idField)
			name := stringField(fields, group.nameField)
			searchText := []string{name}
			for _, field := range group.extraFields {
				searchText = append(searchText, stringField(fields, field))
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

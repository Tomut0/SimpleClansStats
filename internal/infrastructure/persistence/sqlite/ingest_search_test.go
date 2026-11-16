package sqlite

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/Tomut0/SimpleClansStats/internal/application"
)

func TestSearchMatchesClanTagAndNamePrefixes(t *testing.T) {
	db := New(filepath.Join(t.TempDir(), "search.sqlite"))
	defer db.Close()
	if _, err := db.Db.Exec(`CREATE TABLE servers (id TEXT PRIMARY KEY, state_payload TEXT); CREATE VIRTUAL TABLE search_documents USING fts5(server_id UNINDEXED, entity_type UNINDEXED, entity_id UNINDEXED, name UNINDEXED, payload UNINDEXED, search_text, tokenize='unicode61 remove_diacritics 2'); INSERT INTO servers VALUES ('server', '{"clans":[{"externalId":"id","tag":"ABC","name":"Blue Wolves"}],"players":[]}'); INSERT INTO search_documents VALUES ('server', 'clan', 'id', 'ABC', '{"tag":"ABC","name":"Blue Wolves"}', 'ABC Blue Wolves')`); err != nil {
		t.Fatal(err)
	}
	search := application.NewSearchService(NewSQLiteSearchRepository(db.Db))
	for _, query := range []string{"ABC", "Blue", "Wolves"} {
		results, err := search.Search(t.Context(), "server", query, 10)
		if err != nil {
			t.Fatalf("search %q: %v", query, err)
		}
		if len(results) != 1 {
			t.Errorf("search %q returned %d results, want 1", query, len(results))
		}
	}
}

func TestSearchClanDetailsResolveMembersAndRelations(t *testing.T) {
	db := New(filepath.Join(t.TempDir(), "search-details.sqlite"))
	defer db.Close()
	if _, err := db.Db.Exec(`CREATE TABLE servers (id TEXT PRIMARY KEY, state_payload TEXT); CREATE VIRTUAL TABLE search_documents USING fts5(server_id UNINDEXED, entity_type UNINDEXED, entity_id UNINDEXED, name UNINDEXED, payload UNINDEXED, search_text)`); err != nil {
		t.Fatal(err)
	}
	payload := `{"clans":[{"externalId":"1","tag":"ONE","name":"One","leaderExternalId":"uuid-1","memberExternalIds":["uuid-1","uuid-2"],"allyExternalIds":["2"],"rivalExternalIds":["3"]},{"externalId":"2","tag":"ALLY","name":"Ally"},{"externalId":"3","tag":"RIVAL","name":"Rival"}],"players":[{"uuid":"uuid-1","name":"Alice","clanExternalId":"1"},{"uuid":"uuid-2","name":"Bob","role":"member","clanExternalId":"1"}]}`
	if _, err := db.Db.Exec(`INSERT INTO servers VALUES ('server', ?)`, payload); err != nil {
		t.Fatal(err)
	}
	tx, err := db.Db.BeginTx(t.Context(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := indexSearchDocuments(t.Context(), tx, "server", payload); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
	search := application.NewSearchService(NewSQLiteSearchRepository(db.Db))
	results, err := search.Search(t.Context(), "server", "One", 10)
	if err != nil || len(results) != 1 {
		t.Fatalf("search returned %d results (%v)", len(results), err)
	}
	var details struct {
		Members []struct{ Name, Role string } `json:"members"`
		Allies  []string                      `json:"allies"`
		Rivals  []string                      `json:"rivals"`
	}
	if err := json.Unmarshal(results[0].Data, &details); err != nil {
		t.Fatal(err)
	}
	if len(details.Members) != 2 || details.Members[0].Name != "Alice" || details.Members[0].Role != "leader" || details.Members[1].Name != "Bob" {
		t.Fatalf("unexpected members: %+v", details.Members)
	}
	if len(details.Allies) != 1 || details.Allies[0] != "ALLY" || len(details.Rivals) != 1 || details.Rivals[0] != "RIVAL" {
		t.Fatalf("unexpected relations: allies=%v rivals=%v", details.Allies, details.Rivals)
	}
	players, err := search.Search(t.Context(), "server", "Alice", 10)
	if err != nil || len(players) != 1 {
		t.Fatalf("player search returned %d results (%v)", len(players), err)
	}
	var playerData map[string]any
	if err := json.Unmarshal(players[0].Data, &playerData); err != nil {
		t.Fatal(err)
	}
	if _, includesUUID := playerData["uuid"]; includesUUID {
		t.Fatal("player UUID must not be returned by search")
	}
}

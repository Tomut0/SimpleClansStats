package server

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/application"
	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite"
	"github.com/Tomut0/SimpleClansStats/internal/presentation"
	"github.com/alexedwards/scs/v2"
)

type changingSnapshotWriter struct {
	*httptest.ResponseRecorder
	cancel    context.CancelFunc
	flushes   int
	update    func() error
	updateErr error
}

func (w *changingSnapshotWriter) Flush() {
	w.ResponseRecorder.Flush()
	w.flushes++
	if w.flushes == 1 {
		w.updateErr = w.update()
	}
	if w.flushes == 2 || w.updateErr != nil {
		w.cancel()
	}
}

type ingestTestUsers struct {
	ports.UserRepository
	lookupErr bool
}

func (u ingestTestUsers) GetUserById(id int64) (domain.User, error) {
	if u.lookupErr {
		return domain.User{}, errors.New("user lookup failed")
	}
	role := domain.RoleMember
	if id == 1 {
		role = domain.RoleAdmin
	}
	return *domain.NewUser("", "", role, nil, time.Time{}, time.Time{}), nil
}

func TestServerRegistrationAndIngest(t *testing.T) {
	db := sqlite.New(filepath.Join(t.TempDir(), "ingest.sqlite"))
	defer db.Close()
	for _, name := range []string{
		"20260925204348_create_users_table.sql",
		"20261109190000_create_app_settings.sql",
		"20261110120000_create_servers_and_ingest_snapshots.sql",
		"20261111100000_create_dashboard_layouts.sql",
		"20261112100000_create_search_index.sql",
	} {
		migration, err := os.ReadFile("../../sql/migrations/" + name)
		if err != nil {
			t.Fatal(err)
		}
		up, _, _ := strings.Cut(string(migration), "-- +goose Down")
		if _, err := db.Db.Exec(up); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := db.Db.Exec(`INSERT INTO users (id, email, username, password_hash, role) VALUES (1, 'admin@example.com', 'admin', 'hash', 'admin'), (2, 'member@example.com', 'member', 'hash', 'member')`); err != nil {
		t.Fatal(err)
	}

	sessions := scs.New()
	dashboard := application.NewDashboardService(sqlite.NewSQLiteDashboardRepository(db.Db))
	ingest := application.NewIngestService(sqlite.NewSQLiteIngestRepository(db.Db))
	search := application.NewSearchService(sqlite.NewSQLiteSearchRepository(db.Db))
	settings := application.NewSettingsService(sqlite.NewSQLiteSettingsRepository(db.Db))
	servers := application.NewServerService(sqlite.NewSQLiteServerRepository(db.Db))
	router := newRouter(presentation.NewServerHandler(application.NewAuthService(ingestTestUsers{}, nil), sessions, dashboard, ingest, settings, servers, search), sessions, "", false)
	request := func(method, path, body, key string, userID int64, wantStatus int) *httptest.ResponseRecorder {
		t.Helper()
		req := httptest.NewRequest(method, "/api/v1"+path, strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		if key != "" {
			req.Header.Set("X-API-Key", key)
		}
		if userID != 0 {
			ctx, err := sessions.Load(req.Context(), "")
			if err != nil {
				t.Fatal(err)
			}
			sessions.Put(ctx, "user_id", userID)
			token, _, err := sessions.Commit(ctx)
			if err != nil {
				t.Fatal(err)
			}
			req.AddCookie(&http.Cookie{Name: sessions.Cookie.Name, Value: token})
		}
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		if res.Code != wantStatus {
			t.Fatalf("%s %s: got %d, want %d (%s)", method, path, res.Code, wantStatus, res.Body)
		}
		return res
	}

	name := `{"displayName":" Выживание "}`
	request("POST", "/servers", name, "", 0, 401)
	request("POST", "/servers", name, "", 2, 403)
	res := request("POST", "/servers", name, "", 1, 201)
	var registered presentation.ServerRegistration
	if err := json.Unmarshal(res.Body.Bytes(), &registered); err != nil {
		t.Fatal(err)
	}
	if res.Header().Get("Cache-Control") != "no-store" || registered.Server.DisplayName != "Выживание" ||
		!strings.HasPrefix(registered.ApiKey, "sc-") {
		t.Fatal("server registration did not return a name and a private key")
	}
	id := registered.Server.Id.String()
	var storedHash []byte
	if err := db.Db.QueryRow("SELECT api_key_hash FROM servers WHERE id = ?", id).Scan(&storedHash); err != nil {
		t.Fatal(err)
	}
	hash := sha256.Sum256([]byte(registered.ApiKey))
	if string(storedHash) != string(hash[:]) {
		t.Fatal("collector key was not stored as a hash")
	}
	request("GET", "/dashboard/layout", "", "", 0, 200)
	request("PUT", "/dashboard/layout", `{"version":1,"widgets":[{"id":"clans","x":0,"y":0,"w":4,"h":2}]}`, "", 2, 403)
	request("PUT", "/dashboard/layout", `{"version":1,"widgets":[{"id":"clans","x":0,"y":0,"w":4,"h":2}]}`, "", 1, 204)
	request("PUT", "/dashboard/layout", `{"version":1,"widgets":[{"id":"custom-html","x":0,"y":0,"w":4,"h":2}]}`, "", 1, 400)
	layout := request("GET", "/dashboard/layout", "", "", 2, 200)
	if !strings.Contains(layout.Body.String(), `"id":"clans"`) {
		t.Fatalf("layout did not persist: %s", layout.Body.String())
	}
	list := func(online, hasSnapshot bool) {
		t.Helper()
		res := request("GET", "/servers", "", "", 0, 200)
		var servers []presentation.Server
		if err := json.Unmarshal(res.Body.Bytes(), &servers); err != nil {
			t.Fatal(err)
		}
		if len(servers) != 1 || servers[0].Online != online || (servers[0].LastIngestAt != nil) != hasSnapshot || servers[0].Id.String() != id ||
			strings.Contains(res.Body.String(), registered.ApiKey) {
			t.Fatal("unexpected public server list")
		}
	}
	list(false, false)

	payload := `{"takenAt":"2026-11-10T15:00:00.123+03:00","clans":[],"players":[],"extensions":{"scCollectorV1":{"fixture":"` +
		strings.Repeat("x", 5000) + `"}}}`
	request("POST", "/ingest", payload, "", 0, 401)
	request("POST", "/ingest", payload, "invalid", 0, 401)
	request("POST", "/ingest", payload+` {}`, registered.ApiKey, 0, 400)
	request("POST", "/ingest", strings.Replace(payload, `"clans":[]`, `"clans":null`, 1), registered.ApiKey, 0, 400)
	request("POST", "/ingest", `{"serverId":"another-server",`+payload[1:], registered.ApiKey, 0, 400)
	request("POST", "/ingest", strings.Repeat(" ", (8<<20)+1), registered.ApiKey, 0, 413)
	ack := request("POST", "/ingest", payload, registered.ApiKey, 0, 202)
	if !strings.Contains(ack.Body.String(), `"accepted":true`) {
		t.Fatal("collector did not receive an acknowledgment")
	}
	request("POST", "/ingest", strings.Replace(payload, "15:00:00.123+03:00", "12:00:00.123Z", 1), registered.ApiKey, 0, 202)
	list(true, true)
	streamContext, cancelStream := context.WithCancel(context.Background())
	streamRequest := httptest.NewRequest("GET", "/api/v1/dashboard/events?serverId="+id, nil)
	streamContext, err := sessions.Load(streamContext, "")
	if err != nil {
		t.Fatal(err)
	}
	sessions.Put(streamContext, "user_id", int64(1))
	token, _, err := sessions.Commit(streamContext)
	if err != nil {
		t.Fatal(err)
	}
	streamRequest.AddCookie(&http.Cookie{Name: sessions.Cookie.Name, Value: token})
	streamRequest = streamRequest.WithContext(streamContext)
	stream := httptest.NewRecorder()
	streamWriter := &changingSnapshotWriter{
		ResponseRecorder: stream,
		cancel:           cancelStream,
		update: func() error {
			_, err := db.Db.Exec(`UPDATE servers SET state_payload = ? WHERE id = ?`, `{"clans":[{"tag":"updated"}]}`, id)
			return err
		},
	}
	router.ServeHTTP(streamWriter, streamRequest)
	if streamWriter.updateErr != nil {
		t.Fatal(streamWriter.updateErr)
	}
	if streamWriter.flushes != 2 || strings.Count(stream.Body.String(), "event: snapshot") != 2 ||
		!strings.Contains(stream.Body.String(), `"tag":"updated"`) || !strings.Contains(stream.Body.String(), `"events":[]`) {
		t.Fatalf("SSE snapshot was not streamed: %s", stream.Body.String())
	}

	var count int
	if err := db.Db.QueryRow("SELECT count(*) FROM ingest_snapshots").Scan(&count); err != nil || count != 1 {
		t.Fatalf("expected one stored snapshot after retry, got %d (%v)", count, err)
	}
	var takenAt, storedPayload string
	if err := db.Db.QueryRow("SELECT taken_at, payload FROM ingest_snapshots WHERE server_id = ?", id).Scan(&takenAt, &storedPayload); err != nil {
		t.Fatal(err)
	}
	if takenAt != "2026-11-10T12:00:00.123000000Z" || storedPayload != payload {
		t.Fatal("snapshot must preserve its body and use a canonical UTC timestamp")
	}
	const oldReceipt = "2000-01-01T00:00:00.000000000Z"
	if _, err := db.Db.Exec("UPDATE servers SET last_received_at = ? WHERE id = ?", oldReceipt, id); err != nil {
		t.Fatal(err)
	}
	list(false, true)

	heartbeat := strings.Replace(payload, "15:00:00.123+03:00", "15:05:00+03:00", 1)
	heartbeat = strings.Replace(heartbeat, `"fixture":`, `"events":[{"id":"event-1","type":"player.joinedClan"}],"fixture":`, 1)
	request("POST", "/ingest", heartbeat, registered.ApiKey, 0, 202)
	request("POST", "/ingest", heartbeat, registered.ApiKey, 0, 202)
	list(true, true)
	fiveMinutesAgo := time.Now().UTC().Add(-5 * time.Minute).Format(time.RFC3339Nano)
	if _, err := db.Db.Exec("UPDATE servers SET last_received_at = ?", fiveMinutesAgo); err != nil {
		t.Fatal(err)
	}
	list(true, true)
	if _, err := db.Db.Exec("UPDATE ingest_settings SET server_online_timeout_minutes = 3"); err != nil {
		t.Fatal(err)
	}
	list(false, true)
	if err := db.Db.QueryRow("SELECT count(*) FROM ingest_snapshots").Scan(&count); err != nil || count != 1 {
		t.Fatalf("heartbeat or an event created a snapshot: %d (%v)", count, err)
	}
	if err := db.Db.QueryRow("SELECT count(*) FROM ingest_events").Scan(&count); err != nil || count != 1 {
		t.Fatalf("expected one event after retry, got %d (%v)", count, err)
	}

	changed := strings.Replace(heartbeat, "15:05:00", "15:10:00", 1)
	changed = strings.Replace(changed, `"clans":[]`, `"clans":[{"externalId":"x","tag":"X","name":"Clan X"}]`, 1)
	request("POST", "/ingest", changed, registered.ApiKey, 0, 202)
	searchResponse := request("GET", "/search?serverId="+id+"&q=Clan", "", "", 0, 200)
	if !strings.Contains(searchResponse.Body.String(), `"name":"Clan X"`) || strings.Contains(searchResponse.Body.String(), `"uuid"`) {
		t.Fatalf("public search did not return sanitized clan details: %s", searchResponse.Body.String())
	}
	request("POST", "/ingest", heartbeat, registered.ApiKey, 0, 202)
	if err := db.Db.QueryRow("SELECT count(*) FROM ingest_snapshots").Scan(&count); err != nil || count != 2 {
		t.Fatalf("expected initial and changed snapshots only, got %d (%v)", count, err)
	}
	var current string
	if err := db.Db.QueryRow("SELECT state_payload FROM servers WHERE id = ?", id).Scan(&current); err != nil ||
		!strings.Contains(current, "Clan X") || strings.Contains(current, "event-1") {
		t.Fatalf("current state was lost or contains events (%v)", err)
	}
	if _, err := db.Db.Exec("UPDATE ingest_snapshots SET received_at = ? WHERE taken_at = ?", oldReceipt, takenAt); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Db.Exec("UPDATE ingest_events SET received_at = ?", oldReceipt); err != nil {
		t.Fatal(err)
	}
	cutoff := sqlite.FormatIngestTime(time.Now().UTC().AddDate(0, 0, -30))
	if err := db.Queries.PurgeIngestSnapshots(t.Context(), cutoff); err != nil {
		t.Fatal(err)
	}
	if err := db.Queries.PurgeIngestEvents(t.Context(), cutoff); err != nil {
		t.Fatal(err)
	}
	if err := db.Db.QueryRow("SELECT count(*) FROM ingest_snapshots").Scan(&count); err != nil || count != 1 {
		t.Fatalf("cleanup must remove only expired snapshots, got %d (%v)", count, err)
	}
	if err := db.Db.QueryRow("SELECT count(*) FROM ingest_events").Scan(&count); err != nil || count != 0 {
		t.Fatalf("cleanup did not remove expired events: %d (%v)", count, err)
	}
	if _, err := db.Db.Exec("UPDATE ingest_snapshots SET received_at = ?", oldReceipt); err != nil {
		t.Fatal(err)
	}
	if err := db.Queries.PurgeIngestSnapshots(t.Context(), cutoff); err != nil {
		t.Fatal(err)
	}
	unchanged := strings.Replace(changed, "15:10:00", "15:15:00", 1)
	unchanged = strings.Replace(unchanged, `"events":[{"id":"event-1","type":"player.joinedClan"}],`, "", 1)
	request("POST", "/ingest", unchanged, registered.ApiKey, 0, 202)
	if err := db.Db.QueryRow("SELECT count(*) FROM ingest_snapshots").Scan(&count); err != nil || count != 0 {
		t.Fatalf("expired history must not reset the state hash, got %d (%v)", count, err)
	}
	list(true, true)
	if _, err := db.Db.Exec("UPDATE servers SET disabled = TRUE WHERE id = ?", id); err != nil {
		t.Fatal(err)
	}
	request("POST", "/ingest", payload, registered.ApiKey, 0, 401)
	list(false, true)

	request("DELETE", "/servers/"+id, "", "", 0, 401)
	request("DELETE", "/servers/"+id, "", "", 2, 403)
	request("DELETE", "/servers/"+id, "", "", 1, 204)
	request("DELETE", "/servers/"+id, "", "", 1, 204)
	if res := request("GET", "/servers", "", "", 0, 200); strings.TrimSpace(res.Body.String()) != "[]" {
		t.Fatal("deleted server is still listed")
	}
	if err := db.Db.QueryRow("SELECT count(*) FROM ingest_snapshots").Scan(&count); err != nil || count != 0 {
		t.Fatalf("deleted server left snapshots behind: %d (%v)", count, err)
	}
	request("POST", "/ingest", payload, registered.ApiKey, 0, 401)
	router = newRouter(presentation.NewServerHandler(application.NewAuthService(ingestTestUsers{lookupErr: true}, nil), sessions, dashboard, ingest, settings, servers, search), sessions, "", false)
	request("POST", "/servers", name, "", 1, 401)
	request("DELETE", "/servers/"+id, "", "", 1, 401)
}

package presentation

import (
	"bytes"
	"encoding/json"
	"errors"
	"image"
	"image/png"
	"mime/multipart"
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
	"github.com/alexedwards/scs/v2"
)

type settingsTestUsers struct {
	ports.UserRepository
	lookupErr bool
}

func (u settingsTestUsers) GetUserById(id int64) (domain.User, error) {
	if u.lookupErr {
		return domain.User{}, errors.New("user lookup failed")
	}
	role := domain.RoleMember
	if id == 1 {
		role = domain.RoleAdmin
	}
	return *domain.NewUser("", "", role, nil, time.Time{}, time.Time{}), nil
}

func TestAppSettings(t *testing.T) {
	db := sqlite.New(filepath.Join(t.TempDir(), "settings.sqlite"))
	defer db.Close()
	for _, name := range []string{
		"20261109190000_create_app_settings.sql",
		"20261110120000_create_servers_and_ingest_snapshots.sql",
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
	sessions := scs.New()
	dashboard := application.NewDashboardService(sqlite.NewSQLiteDashboardRepository(db.Db))
	ingest := application.NewIngestService(sqlite.NewSQLiteIngestRepository(db.Db))
	search := application.NewSearchService(sqlite.NewSQLiteSearchRepository(db.Db))
	settings := application.NewSettingsService(sqlite.NewSQLiteSettingsRepository(db.Db))
	servers := application.NewServerService(sqlite.NewSQLiteServerRepository(db.Db))
	handler := Handler(NewStrictHandler(NewServerHandler(application.NewAuthService(settingsTestUsers{}, nil), sessions, dashboard, ingest, settings, servers, search), nil))
	request := func(method, path, name string, logo []byte, userID int64, wantStatus int, extra ...map[string]string) *httptest.ResponseRecorder {
		t.Helper()
		var body bytes.Buffer
		var contentType string
		if method == http.MethodPost {
			form := multipart.NewWriter(&body)
			for key, value := range map[string]string{
				"appName": name, "removeAuthor": "true", "removeTelemetry": "true", "timeZone": "America/New_York",
				"dateFormat": "YYYY-MM-DD", "timeFormat": "12h",
				"playerAvatarUrl": "https://mc-heads.net/avatar/{username}/32",
			} {
				if err := form.WriteField(key, value); err != nil {
					t.Fatal(err)
				}
			}
			for _, fields := range extra {
				for key, value := range fields {
					if err := form.WriteField(key, value); err != nil {
						t.Fatal(err)
					}
				}
			}
			if logo != nil {
				part, err := form.CreateFormFile("logo", "logo.png")
				if err != nil {
					t.Fatal(err)
				}
				if _, err := part.Write(logo); err != nil {
					t.Fatal(err)
				}
			}
			if err := form.Close(); err != nil {
				t.Fatal(err)
			}
			contentType = form.FormDataContentType()
		}
		req := httptest.NewRequest(method, path, &body)
		req.Header.Set("Content-Type", contentType)
		res := httptest.NewRecorder()
		sessions.LoadAndSave(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if userID != 0 {
				sessions.Put(r.Context(), "user_id", userID)
			}
			handler.ServeHTTP(w, r)
		})).ServeHTTP(res, req)
		if res.Code != wantStatus {
			t.Fatalf("%s %s: got %d, want %d (%s)", method, path, res.Code, wantStatus, res.Body)
		}
		return res
	}

	initial := request("GET", "/settings", "", nil, 0, 200)
	var defaults AppSettings
	if err := json.Unmarshal(initial.Body.Bytes(), &defaults); err != nil {
		t.Fatal(err)
	}
	if defaults.AppName != "SCStats" {
		t.Fatalf("initial app name: got %q, want SCStats", defaults.AppName)
	}
	if defaults.ServerOnlineTimeoutMinutes != 10 || defaults.IngestRetentionDays != 30 {
		t.Fatalf("unexpected default ingest settings: %+v", defaults)
	}

	var logo bytes.Buffer
	if err := png.Encode(&logo, image.NewRGBA(image.Rect(0, 0, 2, 2))); err != nil {
		t.Fatal(err)
	}
	request("GET", "/settings/logo", "", nil, 0, 404)
	request("POST", "/settings", "Clan Stats", logo.Bytes(), 0, 401)
	request("POST", "/settings", "Clan Stats", logo.Bytes(), 2, 403)
	request("POST", "/settings", "Clan Stats", logo.Bytes(), 1, 200,
		map[string]string{"serverOnlineTimeoutMinutes": "15", "ingestRetentionDays": "0"})
	request("POST", "/settings", "Rejected", nil, 1, 400, map[string]string{"ingestRetentionDays": "-1"})
	request("POST", "/settings", "Rejected", nil, 1, 400, map[string]string{"playerAvatarUrl": "javascript:alert(1)"})
	request("POST", "/settings", "Rejected", []byte("not an image"), 1, 400)

	readSettings := func(wantName string) {
		t.Helper()
		res := request("GET", "/settings", "", nil, 0, 200)
		var got AppSettings
		if err := json.Unmarshal(res.Body.Bytes(), &got); err != nil {
			t.Fatal(err)
		}
		want := AppSettings{AppName: wantName, RemoveAuthor: true, RemoveTelemetry: true, TimeZone: "America/New_York", DateFormat: "YYYY-MM-DD",
			TimeFormat: "12h", PlayerAvatarUrl: "https://mc-heads.net/avatar/{username}/32", LogoVersion: 1, ServerOnlineTimeoutMinutes: 15, IngestRetentionDays: 0}
		if got != want {
			t.Fatalf("stored settings: got %+v, want %+v", got, want)
		}
	}
	readSettings("Clan Stats")
	request("POST", "/settings", "Updated", nil, 1, 200)
	readSettings("Updated")
	res := request("GET", "/settings/logo", "", nil, 0, 200)
	if res.Header().Get("Content-Type") != "image/png" || !bytes.Equal(res.Body.Bytes(), logo.Bytes()) {
		t.Fatal("stored logo differs from the uploaded PNG")
	}
	handler = Handler(NewStrictHandler(NewServerHandler(application.NewAuthService(settingsTestUsers{lookupErr: true}, nil), sessions, dashboard, ingest, settings, servers, search), nil))
	request("POST", "/settings", "Rejected", nil, 1, 401)
}

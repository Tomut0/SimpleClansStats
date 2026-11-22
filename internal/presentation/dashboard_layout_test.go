package presentation

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Tomut0/SimpleClansStats/internal/application"
	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
	"github.com/alexedwards/scs/v2"
)

type dashboardLayoutRepositoryTest struct{ layout domain.DashboardLayout }

func (r *dashboardLayoutRepositoryTest) GetDashboardData(context.Context, string) (domain.DashboardData, error) {
	return domain.DashboardData{}, nil
}
func (r *dashboardLayoutRepositoryTest) GetLayout(context.Context) (domain.DashboardLayout, error) {
	return r.layout, nil
}
func (r *dashboardLayoutRepositoryTest) SaveLayout(_ context.Context, _ int64, layout domain.DashboardLayout) error {
	r.layout = layout
	return nil
}

var _ ports.DashboardRepository = (*dashboardLayoutRepositoryTest)(nil)

func TestDashboardLayoutAccess(t *testing.T) {
	sessions := scs.New()
	clans := domain.DashboardWidget{
		ID: "widget-123e4567-e89b-12d3-a456-426614174000", Type: "counter", Field: "clans", Format: "number",
		W: 12, H: 2,
	}
	repository := &dashboardLayoutRepositoryTest{layout: domain.DashboardLayout{
		Version: 1,
		Widgets: []domain.DashboardWidget{clans},
	}}
	service := application.NewDashboardService(repository)
	handler := NewServerHandler(application.NewAuthService(settingsTestUsers{}, nil), sessions, service, nil, nil, nil, nil)

	get := httptest.NewRecorder()
	sessions.LoadAndSave(http.HandlerFunc(handler.GetDashboardLayout)).ServeHTTP(get, httptest.NewRequest(http.MethodGet, "/dashboard/layout", nil))
	if get.Code != http.StatusOK || !strings.Contains(get.Body.String(), `"id":"`+clans.ID+`"`) {
		t.Fatalf("public layout read: got %d %s", get.Code, get.Body.String())
	}

	put := func(userID int64) *httptest.ResponseRecorder {
		t.Helper()
		response := httptest.NewRecorder()
		body := `{"version":1,"widgets":[{"id":"widget-123e4567-e89b-12d3-a456-426614174001","type":"counter","field":"players","format":"number","x":0,"y":0,"w":12,"h":2}]}`
		request := httptest.NewRequest(http.MethodPut, "/dashboard/layout", strings.NewReader(body))
		sessions.LoadAndSave(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessions.Put(r.Context(), "user_id", userID)
			handler.SaveDashboardLayout(w, r)
		})).ServeHTTP(response, request)
		return response
	}

	if response := put(2); response.Code != http.StatusForbidden {
		t.Fatalf("member layout save: got %d, want %d", response.Code, http.StatusForbidden)
	}
	if response := put(1); response.Code != http.StatusNoContent {
		t.Fatalf("admin layout save: got %d, want %d", response.Code, http.StatusNoContent)
	}
	if repository.layout.Widgets[0].Field != "players" {
		t.Fatalf("shared layout not saved: %+v", repository.layout)
	}
}

package application

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
)

type dashboardRepositoryTest struct {
	saved domain.DashboardLayout
	data  domain.DashboardData
}

func (r *dashboardRepositoryTest) GetDashboardData(context.Context, string) (domain.DashboardData, error) {
	return r.data, nil
}

func TestDashboardDataKeepsHistoryAndOnlyNewestSnapshots(t *testing.T) {
	repository := &dashboardRepositoryTest{data: domain.DashboardData{Snapshots: []json.RawMessage{
		[]byte(`{"takenAt":"2026-11-12T00:00:00Z","clans":[{"tag":"A","balance":35}]}`),
		[]byte(`{"takenAt":"2026-11-11T00:00:00Z","clans":[{"tag":"A","balance":20}]}`),
		[]byte(`{"takenAt":"2026-11-10T00:00:00Z","clans":[{"tag":"A","balance":10}]}`),
	}}}
	data, err := NewDashboardService(repository).Data(context.Background(), "server")
	if err != nil {
		t.Fatal(err)
	}
	if len(data.History) != 3 || !data.History[0].At.Equal(mustParseTime(t, "2026-11-10T00:00:00Z")) {
		t.Fatalf("unexpected history ordering: %+v", data.History)
	}
	if len(data.ClanBalanceChanges) != 1 || string(data.ClanBalanceChanges[0]) != `{"balanceChange":15,"tag":"A"}` {
		t.Fatalf("expected newest balance delta, got %s", data.ClanBalanceChanges)
	}
	if data.Snapshots != nil {
		t.Fatal("snapshots should not be returned")
	}
}

func TestDashboardDataReturnsPartialHistoryOnDecodeError(t *testing.T) {
	repository := &dashboardRepositoryTest{data: domain.DashboardData{Snapshots: []json.RawMessage{
		[]byte(`{"takenAt":"2026-11-10T00:00:00Z"}`),
		[]byte(`not json`),
		[]byte(`{"takenAt":"2026-11-12T00:00:00Z"}`),
	}}}
	data, err := NewDashboardService(repository).Data(context.Background(), "server")
	if err == nil || len(data.History) != 1 || !data.History[0].At.Equal(mustParseTime(t, "2026-11-12T00:00:00Z")) {
		t.Fatalf("expected error and partial newest history, got data=%+v err=%v", data, err)
	}
}

func TestDashboardDataWithOneNullSnapshotHasEmptyBalanceChanges(t *testing.T) {
	repository := &dashboardRepositoryTest{data: domain.DashboardData{Snapshots: []json.RawMessage{[]byte(`null`)}}}
	data, err := NewDashboardService(repository).Data(context.Background(), "server")
	if err != nil || len(data.History) != 1 || data.ClanBalanceChanges == nil || len(data.ClanBalanceChanges) != 0 {
		t.Fatalf("expected one history point and empty changes, got data=%+v err=%v", data, err)
	}
}

func mustParseTime(t *testing.T, value string) time.Time {
	t.Helper()
	parsed, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		t.Fatal(err)
	}
	return parsed
}

func TestDashboardHistoryPointReadsSnapshotMetrics(t *testing.T) {
	snapshot := map[string]any{"takenAt": "2026-11-11T10:00:00Z", "clans": []any{map[string]any{"balance": float64(120)}, map[string]any{"balance": float64(30)}}, "players": []any{map[string]any{}}, "extensions": map[string]any{"scCollectorV1": map[string]any{"wars": []any{map[string]any{}}, "onlinePlayers": []any{map[string]any{}, map[string]any{}}, "killsByType": map[string]any{"allies": float64(4), "rivals": float64(3)}}}}
	encoded, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	var decoded dashboardSnapshot
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	point := dashboardHistoryPoint(decoded)
	if point.Clans != 2 || point.Players != 1 || point.Wars != 1 || point.OnlinePlayers != 2 || point.Balance != 150 || point.Kills != 7 || point.At.IsZero() {
		t.Fatalf("unexpected history point: %+v", point)
	}
}

func TestDashboardClanBalanceChangesUsesLatestMatchingClans(t *testing.T) {
	previous := dashboardSnapshot{Clans: []domain.Clan{{"externalId": "a", "tag": "A", "balance": float64(50)}, {"externalId": "b", "tag": "B", "balance": float64(100)}}}
	current := dashboardSnapshot{Clans: []domain.Clan{{"externalId": "a", "tag": "A", "balance": float64(80)}}}
	changes := dashboardClanBalanceChanges(current, previous)
	if len(changes) != 1 || string(changes[0]) != `{"balanceChange":30,"tag":"A"}` {
		t.Fatalf("unexpected balance changes: %s", changes)
	}
}
func (r *dashboardRepositoryTest) GetLayout(context.Context) (domain.DashboardLayout, error) {
	return r.saved, nil
}
func (r *dashboardRepositoryTest) SaveLayout(_ context.Context, _ int64, layout domain.DashboardLayout) error {
	r.saved = layout
	return nil
}

func TestSaveDashboardLayoutValidatesWidgetBounds(t *testing.T) {
	repository := &dashboardRepositoryTest{}
	service := NewDashboardService(repository)
	valid := domain.DashboardLayout{Version: 1, Widgets: []domain.DashboardWidget{{ID: "clans", X: 0, Y: 0, W: 12, H: 2}}}
	if err := service.SaveLayout(context.Background(), 7, valid); err != nil {
		t.Fatal(err)
	}
	if repository.saved.Widgets[0].ID != "clans" {
		t.Fatal("valid layout was not saved")
	}
	invalid := domain.DashboardLayout{Version: 1, Widgets: []domain.DashboardWidget{{ID: "clans", X: 1, Y: 0, W: 12, H: 2}}}
	if err := service.SaveLayout(context.Background(), 7, invalid); err != ErrInvalidDashboardLayout {
		t.Fatalf("expected invalid layout error, got %v", err)
	}
}

func TestSaveDashboardLayoutAcceptsConfiguredWidgetTypes(t *testing.T) {
	tests := []struct {
		typeName string
		field    string
		format   string
	}{
		{"counter", "clans", "number"},
		{"leaderboard", "clanBalance", ""},
		{"activityFeed", "recentActions", ""},
		{"onlineSnapshot", "onlinePlayers", ""},
		{"distribution", "killTypes", ""},
		{"timeSeries", "balance", "compact"},
		{"clanCard", "clan", ""},
		{"playerCard", "player", ""},
		{"warCenter", "wars", ""},
		{"clanRoster", "clanPlayers", ""},
		{"network", "clanRelations", ""},
		{"compare", "clanMetrics", ""},
		{"heatmap", "playerActivity", ""},
		{"gainersLosers", "clanBalanceChange", ""},
		{"economy", "clanBalance", ""},
		{"killFeed", "recentKills", ""},
		{"peakOnline", "onlinePlayers", ""},
		{"pendingInvites", "pendingInvites", ""},
		{"claims", "claims", ""},
	}
	service := NewDashboardService(&dashboardRepositoryTest{})
	for _, tc := range tests {
		t.Run(tc.typeName, func(t *testing.T) {
			id := "widget-123e4567-e89b-12d3-a456-426614174000"
			if tc.typeName == "counter" {
				id = "counter-123e4567-e89b-12d3-a456-426614174000"
			}
			layout := domain.DashboardLayout{Version: 1, Widgets: []domain.DashboardWidget{{
				ID: id, Type: tc.typeName, Field: tc.field, Format: tc.format, W: 2, H: 2,
			}}}
			if err := service.SaveLayout(context.Background(), 1, layout); err != nil {
				t.Fatalf("expected %s widget to be accepted, got %v", tc.typeName, err)
			}
		})
	}
}

func TestSaveDashboardLayoutRejectsInvalidCounterSettings(t *testing.T) {
	service := NewDashboardService(&dashboardRepositoryTest{})
	invalid := domain.DashboardLayout{Version: 1, Widgets: []domain.DashboardWidget{{
		ID: "counter-123e4567-e89b-12d3-a456-426614174000", Type: "counter", Field: "unknown", Format: "money", W: 2, H: 2,
	}}}
	if err := service.SaveLayout(context.Background(), 1, invalid); err != ErrInvalidDashboardLayout {
		t.Fatalf("expected invalid counter settings error, got %v", err)
	}
}

func TestSaveDashboardLayoutPreservesLegacyAndTableLimits(t *testing.T) {
	service := NewDashboardService(&dashboardRepositoryTest{})
	for _, widget := range []domain.DashboardWidget{
		{ID: "clans", Limit: 99, W: 2, H: 2},
		{ID: "widget-123e4567-e89b-12d3-a456-426614174000", Type: "table", Field: "wars", Limit: 0, W: 2, H: 2},
		{ID: "widget-123e4567-e89b-12d3-a456-426614174000", Type: "table", Field: "wars", Limit: 20, W: 2, H: 2},
	} {
		layout := domain.DashboardLayout{Version: 1, Widgets: []domain.DashboardWidget{widget}}
		if err := service.SaveLayout(context.Background(), 1, layout); err != nil {
			t.Fatalf("expected widget limit %d to be accepted: %v", widget.Limit, err)
		}
	}
	invalid := domain.DashboardLayout{Version: 1, Widgets: []domain.DashboardWidget{{
		ID: "widget-123e4567-e89b-12d3-a456-426614174000", Type: "table", Field: "wars", Limit: 7, W: 2, H: 2,
	}}}
	if err := service.SaveLayout(context.Background(), 1, invalid); err != ErrInvalidDashboardLayout {
		t.Fatalf("expected unsupported table limit to fail, got %v", err)
	}
}

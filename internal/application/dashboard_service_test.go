package application

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
)

type dashboardRepositoryTest struct{ saved domain.DashboardLayout }

func (r *dashboardRepositoryTest) GetDashboardData(context.Context, string) (domain.DashboardData, error) {
	return domain.DashboardData{}, nil
}

func TestDashboardHistoryPointReadsSnapshotMetrics(t *testing.T) {
	snapshot := map[string]any{"takenAt": "2026-11-11T10:00:00Z", "clans": []any{map[string]any{"balance": float64(120)}, map[string]any{"balance": float64(30)}}, "players": []any{map[string]any{}}, "extensions": map[string]any{"scCollectorV1": map[string]any{"wars": []any{map[string]any{}}, "onlinePlayers": []any{map[string]any{}, map[string]any{}}, "killsByType": map[string]any{"allies": float64(4), "rivals": float64(3)}}}}
	encoded, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	point := dashboardHistoryPoint(decoded)
	if point.Clans != 2 || point.Players != 1 || point.Wars != 1 || point.OnlinePlayers != 2 || point.Balance != 150 || point.Kills != 7 || point.At.IsZero() {
		t.Fatalf("unexpected history point: %+v", point)
	}
}

func TestDashboardClanBalanceChangesUsesLatestMatchingClans(t *testing.T) {
	previous := map[string]any{"clans": []any{map[string]any{"externalId": "a", "tag": "A", "balance": float64(50)}, map[string]any{"externalId": "b", "tag": "B", "balance": float64(100)}}}
	current := map[string]any{"clans": []any{map[string]any{"externalId": "a", "tag": "A", "balance": float64(80)}}}
	changes := dashboardClanBalanceChanges(current, previous)
	if len(changes) != 1 || string(changes[0]) != `{"balanceChange":30,"tag":"A"}` {
		t.Fatalf("unexpected balance changes: %s", changes)
	}
}
func (r *dashboardRepositoryTest) GetLayout(context.Context, int64) (domain.DashboardLayout, error) {
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

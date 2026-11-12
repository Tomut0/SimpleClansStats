package application

import (
	"context"
	"encoding/json"
	"errors"
	"regexp"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
)

type DashboardService struct {
	repository ports.DashboardRepository
}

type dashboardSnapshot struct {
	TakenAt    string          `json:"takenAt"`
	Clans      []domain.Clan   `json:"clans"`
	Players    []domain.Player `json:"players"`
	Extensions map[string]any  `json:"extensions"`
}

var ErrInvalidDashboardLayout = errors.New("invalid dashboard layout")

var dashboardWidgetIDs = map[string]bool{"clans": true, "players": true, "wars": true, "balance": true, "kills": true, "status": true, "actions": true, "top": true, "online": true}
var dashboardWidgetFields = map[string]map[string]bool{
	"counter":        {"clans": true, "players": true, "wars": true, "balance": true, "onlinePlayers": true, "kills": true},
	"table":          {"wars": true, "clanPlayers": true, "pendingInvites": true, "clanBalanceChange": true},
	"leaderboard":    {"clanBalance": true},
	"activityFeed":   {"recentActions": true},
	"onlineSnapshot": {"onlinePlayers": true},
	"distribution":   {"killTypes": true, "clanStatus": true},
	"timeSeries":     {"clans": true, "players": true, "wars": true, "balance": true, "onlinePlayers": true, "kills": true},
	"clanCard":       {"clan": true},
	"playerCard":     {"player": true},
	"warCenter":      {"wars": true},
	"clanRoster":     {"clanPlayers": true},
	"network":        {"clanRelations": true},
	"compare":        {"clanMetrics": true},
	"heatmap":        {"playerActivity": true},
	"gainersLosers":  {"clanBalanceChange": true},
	"economy":        {"clanBalance": true},
	"killFeed":       {"recentKills": true},
	"peakOnline":     {"onlinePlayers": true},
	"pendingInvites": {"pendingInvites": true},
	"claims":         {"claims": true},
}
var dashboardCounterFormats = map[string]bool{"number": true, "compact": true, "plain": true}
var dashboardCounterIDs = regexp.MustCompile(`^counter-[0-9a-fA-F-]{36}$`)
var dashboardWidgetInstanceIDs = regexp.MustCompile(`^widget-[0-9a-fA-F-]{36}$`)

func NewDashboardService(repository ports.DashboardRepository) *DashboardService {
	return &DashboardService{repository: repository}
}

func (s *DashboardService) Data(ctx context.Context, serverID string) (domain.DashboardData, error) {
	data, err := s.repository.GetDashboardData(ctx, serverID)
	if err != nil {
		return data, err
	}
	var current, previous dashboardSnapshot
	for i := len(data.Snapshots) - 1; i >= 0; i-- {
		var snapshot dashboardSnapshot
		if err := json.Unmarshal(data.Snapshots[i], &snapshot); err != nil {
			return data, err
		}
		if i == 0 {
			current = snapshot
		} else if i == 1 {
			previous = snapshot
		}
		data.History = append(data.History, dashboardHistoryPoint(snapshot))
	}
	if len(data.Snapshots) > 1 {
		data.ClanBalanceChanges = dashboardClanBalanceChanges(current, previous)
	}
	if data.ClanBalanceChanges == nil {
		data.ClanBalanceChanges = []json.RawMessage{}
	}
	data.Snapshots = nil
	return data, nil
}

func dashboardClanBalanceChanges(current, previous dashboardSnapshot) []json.RawMessage {
	oldBalances := map[string]float64{}
	for _, clan := range previous.Clans {
		oldBalances[clan.IdentityKey()] = clan.Balance()
	}
	var changes []json.RawMessage
	for _, clan := range current.Clans {
		previousBalance, exists := oldBalances[clan.IdentityKey()]
		if !exists {
			continue
		}
		change, _ := json.Marshal(map[string]any{"tag": clan["tag"], "balanceChange": clan.Balance() - previousBalance})
		changes = append(changes, change)
	}
	return changes
}

func dashboardHistoryPoint(snapshot dashboardSnapshot) domain.DashboardHistoryPoint {
	point := domain.DashboardHistoryPoint{At: time.Time{}}
	point.At, _ = time.Parse(time.RFC3339Nano, snapshot.TakenAt)
	point.Clans, point.Players = len(snapshot.Clans), len(snapshot.Players)
	for _, clan := range snapshot.Clans {
		point.Balance += clan.Balance()
	}
	if extensions := snapshot.Extensions; extensions != nil {
		if collector, ok := extensions["scCollectorV1"].(map[string]any); ok {
			wars, _ := collector["wars"].([]any)
			online, _ := collector["onlinePlayers"].([]any)
			point.Wars, point.OnlinePlayers = len(wars), len(online)
			kills := collector["killsByType"]
			if killRows, ok := kills.([]any); ok {
				for _, row := range killRows {
					if item, ok := row.(map[string]any); ok {
						point.Kills += numberValue(item["count"])
					}
				}
			} else if killMap, ok := kills.(map[string]any); ok {
				for _, count := range killMap {
					point.Kills += numberValue(count)
				}
			}
		}
	}
	return point
}

func numberValue(value any) float64 {
	switch number := value.(type) {
	case float64:
		return number
	case json.Number:
		value, _ := number.Float64()
		return value
	default:
		return 0
	}
}

func (s *DashboardService) Layout(ctx context.Context) (domain.DashboardLayout, error) {
	return s.repository.GetLayout(ctx)
}

func (s *DashboardService) SaveLayout(ctx context.Context, userID int64, layout domain.DashboardLayout) error {
	if layout.Version != 1 || len(layout.Widgets) == 0 || len(layout.Widgets) > 24 {
		return ErrInvalidDashboardLayout
	}
	seen := make(map[string]bool, len(layout.Widgets))
	for _, widget := range layout.Widgets {
		if !validDashboardWidget(widget) || seen[widget.ID] || widget.X < 0 || widget.X > 11 || widget.Y < 0 || widget.Y > 10000 || widget.W < 1 || widget.W > 12 || widget.H < 1 || widget.H > 12 || widget.X+widget.W > 12 {
			return ErrInvalidDashboardLayout
		}
		seen[widget.ID] = true
	}
	return s.repository.SaveLayout(ctx, userID, layout)
}

func validDashboardWidget(widget domain.DashboardWidget) bool {
	if widget.Type == "" {
		return dashboardWidgetIDs[widget.ID] && widget.Field == "" && widget.Format == ""
	}
	fields, supportedType := dashboardWidgetFields[widget.Type]
	if !supportedType || !fields[widget.Field] {
		return false
	}
	if widget.Type == "counter" {
		legacyField, legacy := map[string]string{"clans": "clans", "players": "players", "wars": "wars", "balance": "balance"}[widget.ID]
		if !((legacy && legacyField == widget.Field) || dashboardCounterIDs.MatchString(widget.ID)) || !dashboardCounterFormats[widget.Format] {
			return false
		}
	} else if !dashboardWidgetInstanceIDs.MatchString(widget.ID) || (widget.Format != "" && !dashboardCounterFormats[widget.Format]) {
		return false
	}
	if widget.Type == "table" {
		return widget.Limit == 0 || widget.Limit == 5 || widget.Limit == 10 || widget.Limit == 20
	}
	return widget.Limit == 0
}

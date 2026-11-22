package application

import (
	"context"
	"errors"
	"regexp"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
)

const (
	dashboardColumns    = 12
	dashboardMaxRow     = 10000
	dashboardMaxHeight  = 12
	dashboardMaxWidgets = 24
)

var ErrInvalidDashboardLayout = errors.New("invalid dashboard layout")

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

var dashboardWidgetFormats = map[string]bool{"number": true, "compact": true, "plain": true}
var dashboardWidgetIDPattern = regexp.MustCompile(`^widget-[0-9a-fA-F-]{36}$`)

func (s *DashboardService) Layout(ctx context.Context) (domain.DashboardLayout, error) {
	layout, err := s.repository.GetLayout(ctx)
	if err != nil || validDashboardLayout(layout) {
		return layout, err
	}
	// An empty list tells the web dashboard to render its built-in layout.
	return domain.DashboardLayout{Version: domain.DashboardLayoutVersion, Widgets: []domain.DashboardWidget{}}, nil
}

func (s *DashboardService) SaveLayout(ctx context.Context, userID int64, layout domain.DashboardLayout) error {
	if !validDashboardLayout(layout) {
		return ErrInvalidDashboardLayout
	}
	return s.repository.SaveLayout(ctx, userID, layout)
}

func validDashboardLayout(layout domain.DashboardLayout) bool {
	if layout.Version != domain.DashboardLayoutVersion || len(layout.Widgets) == 0 || len(layout.Widgets) > dashboardMaxWidgets {
		return false
	}

	seenIDs := make(map[string]bool, len(layout.Widgets))
	for _, widget := range layout.Widgets {
		if seenIDs[widget.ID] || !validDashboardWidget(widget) || !widgetFitsDashboard(widget) {
			return false
		}
		seenIDs[widget.ID] = true
	}
	return true
}

func widgetFitsDashboard(widget domain.DashboardWidget) bool {
	validPosition := widget.X >= 0 && widget.X < dashboardColumns && widget.Y >= 0 && widget.Y <= dashboardMaxRow
	validSize := widget.W >= 1 && widget.W <= dashboardColumns && widget.H >= 1 && widget.H <= dashboardMaxHeight
	return validPosition && validSize && widget.X+widget.W <= dashboardColumns
}

func validDashboardWidget(widget domain.DashboardWidget) bool {
	if !dashboardWidgetIDPattern.MatchString(widget.ID) || !dashboardWidgetFields[widget.Type][widget.Field] {
		return false
	}

	if widget.Type == "counter" {
		return dashboardWidgetFormats[widget.Format] && widget.Limit == 0
	}

	if widget.Format != "" && !dashboardWidgetFormats[widget.Format] {
		return false
	}
	if widget.Type == "table" {
		return widget.Limit == 0 || widget.Limit == 5 || widget.Limit == 10 || widget.Limit == 20
	}
	return widget.Limit == 0
}

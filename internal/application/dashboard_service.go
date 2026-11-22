package application

import (
	"context"
	"encoding/json"
	"slices"
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

func NewDashboardService(repository ports.DashboardRepository) *DashboardService {
	return &DashboardService{repository: repository}
}

func (s *DashboardService) Data(ctx context.Context, serverID string) (domain.DashboardData, error) {
	data, err := s.repository.GetDashboardData(ctx, serverID)
	if err != nil {
		return data, err
	}
	var latest, previous dashboardSnapshot
	// The repository returns newest first. Charts read the history oldest first.
	for index, raw := range slices.Backward(data.Snapshots) {
		var snapshot dashboardSnapshot
		if err := json.Unmarshal(raw, &snapshot); err != nil {
			return data, err
		}
		switch index {
		case 0:
			latest = snapshot
		case 1:
			previous = snapshot
		}
		data.History = append(data.History, dashboardHistoryPoint(snapshot))
	}
	if len(data.Snapshots) > 1 {
		data.ClanBalanceChanges = dashboardClanBalanceChanges(latest, previous)
	}
	if data.ClanBalanceChanges == nil {
		data.ClanBalanceChanges = []json.RawMessage{}
	}
	data.Snapshots = nil
	return data, nil
}

func dashboardClanBalanceChanges(latest, previous dashboardSnapshot) []json.RawMessage {
	previousBalances := make(map[string]float64, len(previous.Clans))
	for _, clan := range previous.Clans {
		previousBalances[clan.IdentityKey()] = clan.Balance()
	}
	var changes []json.RawMessage
	for _, clan := range latest.Clans {
		previousBalance, exists := previousBalances[clan.IdentityKey()]
		if !exists {
			continue
		}
		change, _ := json.Marshal(map[string]any{
			"tag":           clan["tag"],
			"balanceChange": clan.Balance() - previousBalance,
		})
		changes = append(changes, change)
	}
	return changes
}

func dashboardHistoryPoint(snapshot dashboardSnapshot) domain.DashboardHistoryPoint {
	takenAt, _ := time.Parse(time.RFC3339Nano, snapshot.TakenAt)
	point := domain.DashboardHistoryPoint{
		At:      takenAt,
		Clans:   len(snapshot.Clans),
		Players: len(snapshot.Players),
	}
	for _, clan := range snapshot.Clans {
		point.Balance += clan.Balance()
	}

	collector, _ := snapshot.Extensions["scCollectorV1"].(map[string]any)
	wars, _ := collector["wars"].([]any)
	onlinePlayers, _ := collector["onlinePlayers"].([]any)
	point.Wars = len(wars)
	point.OnlinePlayers = len(onlinePlayers)
	point.Kills = dashboardKillCount(collector["killsByType"])
	return point
}

// The collector has sent kill counts both as rows and as a map keyed by kill type.
func dashboardKillCount(killsByType any) float64 {
	var total float64
	switch counts := killsByType.(type) {
	case []any:
		for _, row := range counts {
			item, _ := row.(map[string]any)
			total += numberValue(item["count"])
		}
	case map[string]any:
		for _, count := range counts {
			total += numberValue(count)
		}
	}
	return total
}

func numberValue(value any) float64 {
	switch number := value.(type) {
	case float64:
		return number
	case json.Number:
		parsed, _ := number.Float64()
		return parsed
	default:
		return 0
	}
}

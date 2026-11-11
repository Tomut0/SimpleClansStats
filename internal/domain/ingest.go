package domain

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
)

type DashboardData struct {
	State              json.RawMessage         `json:"state"`
	Events             []json.RawMessage       `json:"events"`
	History            []DashboardHistoryPoint `json:"history"`
	Snapshots          []json.RawMessage       `json:"-"`
	ClanBalanceChanges []json.RawMessage       `json:"clanBalanceChanges"`
}

type DashboardHistoryPoint struct {
	At            time.Time `json:"at"`
	Clans         int       `json:"clans"`
	Players       int       `json:"players"`
	Wars          int       `json:"wars"`
	Balance       float64   `json:"balance"`
	OnlinePlayers int       `json:"onlinePlayers"`
	Kills         float64   `json:"kills"`
}

type IngestEvent struct {
	ID      string
	Payload string
}

type IngestSnapshot struct {
	TakenAt  time.Time `json:"takenAt"`
	ServerID string    `json:"serverId"`
	Clans    []struct {
		ExternalID string `json:"externalId"`
		Tag        string `json:"tag"`
		Name       string `json:"name"`
	} `json:"clans"`
	Players []struct {
		UUID uuid.UUID `json:"uuid"`
		Name string    `json:"name"`
	} `json:"players"`
}

func (s IngestSnapshot) Valid(serverID string) bool {
	if s.TakenAt.IsZero() || s.Clans == nil || s.Players == nil || (s.ServerID != "" && !strings.EqualFold(s.ServerID, serverID)) {
		return false
	}
	for _, clan := range s.Clans {
		if strings.TrimSpace(clan.ExternalID) == "" || strings.TrimSpace(clan.Tag) == "" || strings.TrimSpace(clan.Name) == "" {
			return false
		}
	}
	for _, player := range s.Players {
		if player.UUID == uuid.Nil || strings.TrimSpace(player.Name) == "" {
			return false
		}
	}
	return true
}

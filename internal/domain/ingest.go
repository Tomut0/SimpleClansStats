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
	Clans    []Clan    `json:"clans"`
	Players  []Player  `json:"players"`
}

func (s IngestSnapshot) Valid(serverID string) bool {
	if s.TakenAt.IsZero() || s.Clans == nil || s.Players == nil || (s.ServerID != "" && !strings.EqualFold(s.ServerID, serverID)) {
		return false
	}
	for _, clan := range s.Clans {
		if strings.TrimSpace(stringField(clan, "externalId")) == "" || strings.TrimSpace(stringField(clan, "tag")) == "" || strings.TrimSpace(stringField(clan, "name")) == "" {
			return false
		}
	}
	for _, player := range s.Players {
		id, err := uuid.Parse(stringField(player, "uuid"))
		if err != nil || id == uuid.Nil || strings.TrimSpace(stringField(player, "name")) == "" {
			return false
		}
	}
	return true
}

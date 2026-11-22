package domain

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"sort"
)

type SearchResult struct {
	Type string          `json:"type"`
	Name string          `json:"name"`
	Data json.RawMessage `json:"data"`
}

type Clan map[string]any
type Player map[string]any

func (clan Clan) IdentityKey() string {
	for _, key := range []string{"externalId", "tag", "name"} {
		if value := stringField(clan, key); value != "" {
			return key + ":" + value
		}
	}
	return ""
}

func (clan Clan) Balance() float64 {
	switch value := clan["balance"].(type) {
	case float64:
		return value
	case json.Number:
		result, _ := value.Float64()
		return result
	default:
		return 0
	}
}

type SearchMatch struct {
	Type string
	ID   string
	Name string
}

type SearchData struct {
	Matches []SearchMatch
	Clans   []json.RawMessage
	Players []json.RawMessage
}

func ResolveSearchResults(matches []SearchMatch, clans, players []json.RawMessage) ([]SearchResult, error) {
	lookup, err := buildSearchLookup(clans, players)
	if err != nil {
		return nil, err
	}
	results := make([]SearchResult, 0, len(matches))
	for _, match := range matches {
		var fields map[string]any
		if match.Type == "player" {
			fields = lookup.playersByID[match.ID]
		} else {
			fields = lookup.clansByID[match.ID]
		}
		if fields == nil {
			continue
		}

		// Other results may still need the original fields to resolve their relationships.
		fields = maps.Clone(fields)
		data, err := searchDetails(match.Type, fields, lookup)
		if err != nil {
			return nil, err
		}
		results = append(results, SearchResult{Type: match.Type, Name: match.Name, Data: data})
	}
	return results, nil
}

type searchLookup struct {
	playersByID   map[string]Player
	playersByClan map[string][]Player
	clansByID     map[string]Clan
}

var playerIDFields = []string{"uuid", "externalId", "id"}

func buildSearchLookup(clans, players []json.RawMessage) (searchLookup, error) {
	lookup := searchLookup{
		playersByID:   make(map[string]Player),
		playersByClan: make(map[string][]Player),
		clansByID:     make(map[string]Clan, len(clans)),
	}
	for _, raw := range players {
		var player Player
		if err := json.Unmarshal(raw, &player); err != nil {
			return lookup, err
		}
		for _, key := range playerIDFields {
			if value := player[key]; value != nil {
				lookup.playersByID[fmt.Sprint(value)] = player
			}
		}
		if clanID := stringField(player, "clanExternalId"); clanID != "" {
			lookup.playersByClan[clanID] = append(lookup.playersByClan[clanID], player)
		}
	}
	for _, raw := range clans {
		var clan Clan
		if err := json.Unmarshal(raw, &clan); err != nil {
			return lookup, err
		}
		lookup.clansByID[stringField(clan, "externalId")] = clan
	}
	return lookup, nil
}

func searchDetails(kind string, fields map[string]any, lookup searchLookup) (json.RawMessage, error) {
	if kind == "clan" {
		clan := Clan(fields)
		clan["members"] = clanMembers(clan, lookup.playersByID, lookup.playersByClan)
		clan["allies"] = clanRelations(clan["allyExternalIds"], lookup.clansByID)
		clan["rivals"] = clanRelations(clan["rivalExternalIds"], lookup.clansByID)
		for _, key := range []string{"externalId", "leaderExternalId", "memberExternalIds", "allyExternalIds", "rivalExternalIds"} {
			delete(clan, key)
		}
	} else {
		player := Player(fields)
		for _, key := range []string{"uuid", "id", "externalId", "playerId", "clanExternalId", "pastClanExternalIds"} {
			delete(player, key)
		}
	}
	return json.Marshal(fields)
}

type searchClanMember struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

func clanMembers(clan Clan, playersByID map[string]Player, playersByClan map[string][]Player) []searchClanMember {
	leaderID := fmt.Sprint(clan["leaderExternalId"])
	memberIDs, _ := clan["memberExternalIds"].([]any)
	members := make([]searchClanMember, 0, len(memberIDs))
	for _, value := range memberIDs {
		id := fmt.Sprint(value)
		if player := playersByID[id]; player != nil {
			members = append(members, describeClanMember(player, id == leaderID))
		}
	}

	// Some snapshots only record membership on the player.
	if len(members) == 0 {
		for _, player := range playersByClan[stringField(clan, "externalId")] {
			isLeader := slices.ContainsFunc(playerIDFields, func(field string) bool {
				return fmt.Sprint(player[field]) == leaderID
			})
			members = append(members, describeClanMember(player, isLeader))
		}
	}

	sort.SliceStable(members, func(i, j int) bool {
		return members[i].Role == "leader" && members[j].Role != "leader"
	})
	return members
}

func describeClanMember(player Player, isLeader bool) searchClanMember {
	role := stringField(player, "role")
	switch {
	case isLeader:
		role = "leader"
	case role == "":
		role = "member"
	}
	return searchClanMember{Name: stringField(player, "name"), Role: role}
}

func clanRelations(ids any, clansByID map[string]Clan) []string {
	values, _ := ids.([]any)
	relations := make([]string, 0, len(values))
	for _, value := range values {
		if tag := stringField(clansByID[fmt.Sprint(value)], "tag"); tag != "" {
			relations = append(relations, tag)
		}
	}
	return relations
}

func stringField(fields map[string]any, key string) string {
	value, _ := fields[key].(string)
	return value
}

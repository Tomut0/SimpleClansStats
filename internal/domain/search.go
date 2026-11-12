package domain

import (
	"encoding/json"
	"fmt"
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
		var fields map[string]any = lookup.clansByID[match.ID]
		if match.Type == "player" {
			fields = lookup.playersByID[match.ID]
		}
		if fields == nil {
			continue
		}
		encoded, err := json.Marshal(fields)
		if err != nil {
			return nil, err
		}
		fields = map[string]any{}
		if err := json.Unmarshal(encoded, &fields); err != nil {
			return nil, err
		}
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

func buildSearchLookup(clans, players []json.RawMessage) (searchLookup, error) {
	lookup := searchLookup{playersByID: map[string]Player{}, playersByClan: map[string][]Player{}, clansByID: map[string]Clan{}}
	for _, raw := range players {
		var player Player
		if err := json.Unmarshal(raw, &player); err != nil {
			return lookup, err
		}
		for _, key := range []string{"uuid", "externalId", "id"} {
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

func clanMembers(clan Clan, playersByID map[string]Player, playersByClan map[string][]Player) []map[string]string {
	members, matched := []map[string]string{}, 0
	if ids, ok := clan["memberExternalIds"].([]any); ok {
		for _, value := range ids {
			id := fmt.Sprint(value)
			player := playersByID[id]
			if player == nil {
				continue
			}
			matched++
			role := stringField(player, "role")
			if role == "" {
				role = "member"
			}
			if id == fmt.Sprint(clan["leaderExternalId"]) {
				role = "leader"
			}
			members = append(members, map[string]string{"name": stringField(player, "name"), "role": role})
		}
	}
	if matched == 0 {
		for _, player := range playersByClan[stringField(clan, "externalId")] {
			role := stringField(player, "role")
			if role == "" {
				role = "member"
			}
			for _, key := range []string{"uuid", "externalId", "id"} {
				if fmt.Sprint(player[key]) == fmt.Sprint(clan["leaderExternalId"]) {
					role = "leader"
				}
			}
			members = append(members, map[string]string{"name": stringField(player, "name"), "role": role})
		}
	}
	sort.SliceStable(members, func(i, j int) bool { return members[i]["role"] == "leader" && members[j]["role"] != "leader" })
	return members
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

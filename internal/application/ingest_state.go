package application

import (
	"bytes"
	"encoding/json"
	"errors"
	"sort"
	"strings"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
)

func NormalizeIngestState(payload []byte) ([]byte, []domain.IngestEvent, error) {
	state, rawEvents, err := canonicalIngestState(payload)
	if err != nil {
		return nil, nil, err
	}
	if rawEvents == nil {
		return state, nil, nil
	}
	events, ok := rawEvents.([]any)
	if !ok {
		return nil, nil, errors.New("events must be an array")
	}
	result := make([]domain.IngestEvent, 0, len(events))
	for _, event := range events {
		object, _ := event.(map[string]any)
		id, _ := object["id"].(string)
		if strings.TrimSpace(id) == "" {
			return nil, nil, errors.New("event ID is required")
		}
		data, err := json.Marshal(object)
		if err != nil {
			return nil, nil, err
		}
		result = append(result, domain.IngestEvent{ID: id, Payload: string(data)})
	}
	return state, result, nil
}

func canonicalIngestState(payload []byte) ([]byte, any, error) {
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.UseNumber()
	var state map[string]any
	if err := decoder.Decode(&state); err != nil {
		return nil, nil, err
	}
	delete(state, "takenAt")
	delete(state, "serverId")

	// SCCollector puts its v1 widget data and activity events in this extension.
	extensions, _ := state["extensions"].(map[string]any)
	collector, _ := extensions["scCollectorV1"].(map[string]any)
	events := collector["events"]
	delete(collector, "events")
	if len(collector) == 0 && collector != nil {
		delete(extensions, "scCollectorV1")
	}
	if len(extensions) == 0 && extensions != nil {
		delete(state, "extensions")
	}

	sortStateList(state, "clans", "memberExternalIds", "allyExternalIds", "rivalExternalIds")
	sortStateList(state, "players", "pastClanExternalIds")
	sortStateList(collector, "onlinePlayers")
	sortStateList(collector, "wars", "clanExternalIds")
	sortStateList(collector, "pendingInvites")
	sortStateList(collector, "claims")

	canonical, err := json.Marshal(state)
	return canonical, events, err
}

// Only known unordered collections are sorted. Other extension arrays retain their order.
func sortStateList(object map[string]any, key string, childLists ...string) {
	items, ok := object[key].([]any)
	if !ok {
		return
	}
	values := make([]string, 0, len(items))
	for _, item := range items {
		fields, _ := item.(map[string]any)
		for _, child := range childLists {
			sortStateList(fields, child)
		}
		encoded, _ := json.Marshal(item) // Items were decoded from JSON above.
		values = append(values, string(encoded))
	}
	sort.Strings(values)
	sorted := make([]json.RawMessage, len(values))
	for i, value := range values {
		sorted[i] = json.RawMessage(value)
	}
	object[key] = sorted
}

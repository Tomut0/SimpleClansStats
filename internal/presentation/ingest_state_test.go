package presentation

import (
	"bytes"
	"testing"
)

func TestCanonicalIngestState(t *testing.T) {
	first := `{"takenAt":"first","clans":[{"tag":"A","memberExternalIds":["2","1"]},{"tag":"B"}],
		"extensions":{"scCollectorV1":{"events":[{"id":"a"}],"onlinePlayers":[{"name":"x"},{"name":"y"}]},"ordered":[1,2]}}`
	reordered := `{"clans":[{"tag":"B"},{"memberExternalIds":["1","2"],"tag":"A"}],"takenAt":"second",
		"extensions":{"ordered":[1,2],"scCollectorV1":{"onlinePlayers":[{"name":"y"},{"name":"x"}],"events":[{"id":"b"}]}}}`
	a, _, err := canonicalIngestState([]byte(first))
	if err != nil {
		t.Fatal(err)
	}
	b, events, err := ingestState([]byte(reordered))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(a, b) || len(events) != 1 || events[0].id != "b" {
		t.Fatal("an unordered list or event changed the snapshot state")
	}
	changed, _, err := canonicalIngestState(bytes.Replace([]byte(reordered), []byte("[1,2]"), []byte("[2,1]"), 1))
	if err != nil || bytes.Equal(b, changed) {
		t.Fatalf("an ordered extension array must remain ordered (%v)", err)
	}
}

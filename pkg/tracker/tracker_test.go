package tracker

import (
	"testing"
)

func TestNewTracker(t *testing.T) {
	id := "testid"
	subject := "abc123"
	tpl := NewPayload(pl)
	tr := NewTracker(tpl, id, subject)

	if tr.Id != id {
		t.Errorf("Tracker has unexpected id value: got %v want %v", tr.Id, id)
	}
}

package aggregate

import (
	"github.com/google/uuid"
	"goes/core/event"
	"testing"
)

type TestEvent struct {}

func TestNewAggregate(t *testing.T) {
	id := uuid.New()
	aggregate := NewAggregate(id)
	if aggregate.Id() != id {
		t.Errorf("expected id to be %s but was %s", id, aggregate.Id())
	}
	size := len(aggregate.GetChanges())
	if size != 0 {
		t.Errorf("expected changes to be empty was %d", size)
	}
	originalVersion := aggregate.OriginalVersion()
	if originalVersion != -1 {
		t.Errorf("expected original version to be -1 was %d", originalVersion)
	}
	currentVersion := aggregate.CurrentVersion()
	if currentVersion != -1 {
		t.Errorf("expected current version to be -1 was %d", currentVersion)
	}
}

func TestAggregate_IncrementVersion(t *testing.T) {
	aggregate := NewAggregate(uuid.New())
	aggregate.IncrementVersion()
	originalVersion := aggregate.OriginalVersion()
	if originalVersion != 0 {
		t.Errorf("expected original version to be 0 was %d", originalVersion)
	}
	currentVersion := aggregate.CurrentVersion()
	if currentVersion != 0 {
		t.Errorf("expected current version to be 0 was %d", currentVersion)
	}
}

func TestAggregate_ClearChanges(t *testing.T) {
	id := uuid.New()
	aggregate := NewAggregate(id)
	aggregate.changes = []event.IEvent{
		event.NewEvent(uuid.New(), id, TestEvent{}),
	}
	aggregate.ClearChanges()
	size := len(aggregate.GetChanges())
	if size != 0 {
		t.Errorf("expected changes to be empty was %d", size)
	}
}

func TestAggregate_TrackChange(t *testing.T) {
	id := uuid.New()
	aggregate := NewAggregate(id)
	aggregate.TrackChange(event.NewEvent(uuid.New(), id, TestEvent{}))
	size := len(aggregate.GetChanges())
	if size != 1 {
		t.Errorf("expected one change to be tracked but was %d", size)
	}
}
package core

import (
	"github.com/google/uuid"
	"github.com/tilau2328/goes/core/event"
	"testing"
)

type TestEvent struct{}

func TestLoad(t *testing.T) {
	aggregateId1 := uuid.New()
	aggregateId2 := uuid.New()
	event1 := event.NewEvent(uuid.New(), aggregateId1, TestEvent{})
	event2 := event.NewEvent(uuid.New(), aggregateId1, TestEvent{})
	event3 := event.NewEvent(uuid.New(), aggregateId2, TestEvent{})
	store := &Store{events: map[string][]event.IEvent{
		aggregateId1.String(): {event1, event2},
		aggregateId2.String(): {event3}},
	}
	events := store.Load(aggregateId1)
	size := len(events)
	if size < 2 {
		t.Errorf("expected to have 2 events for aggregateId1 instead there are %d", size)
	}
	if events[0] != event1 {
		t.Errorf("expected event %T to be %T", events[0], event1)
	}
	if events[1] != event2 {
		t.Errorf("expected event %T to be %T", events[1], event2)
	}
	events = store.Load(aggregateId2)
	size = len(events)
	if size < 1 {
		t.Errorf("expected to have 2 events for aggregateId1 instead there are %d", size)
	}
	if events[0] != event1 {
		t.Errorf("expected event %T to be %T", events[0], event3)
	}
}

func TestStore(t *testing.T) {
	aggregateId := uuid.New()
	event1 := event.NewEvent(uuid.New(), aggregateId, TestEvent{})
	event2 := event.NewEvent(uuid.New(), aggregateId, TestEvent{})
	store := &Store{events: map[string][]event.IEvent{aggregateId.String(): {event1}}}
	err := store.Store([]event.IEvent{event2})
	if err != nil {
		t.Error(err)
	}
	events := store.events[aggregateId.String()]
	size := len(events)
	if size < 2 {
		t.Errorf("expected to have 2 events instead there are %d", size)
	}
	if events[0] != event1 {
		t.Errorf("expected %T to be %T", events[0], event1)
	}
	res2 := events[1]
	if res2 != event2 {
		t.Errorf("expected %T to be %T", events[1], event2)
	}
}

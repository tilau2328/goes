package repository

import (
	"github.com/google/uuid"
	"goes/core"
	"goes/core/event"
	"goes/domain/aggregate"
	"testing"
)

type TestEvent struct{}
type TestEventHandler struct{ event event.IEvent }

func (t *TestEventHandler) Handle(event event.IEvent) error {
 	t.event = event
	return nil
}
func testFunc(event.IEvent, bool) {}

func TestLoad(t *testing.T) {
	bus := event.NewBus()
	store := core.NewStore(bus)
	factory := aggregate.NewFactory()
	err := factory.Register((*aggregate.Aggregate)(nil), func(id uuid.UUID) aggregate.IAggregate {
		return aggregate.NewAggregate(id)
	})
	if err != nil {
		t.Error(err)
	}
	aggregateId := uuid.New()
	repository := NewRepository(store, bus, factory)
	a, err := repository.Load("aggregate.Aggregate", aggregateId)
	if err != nil {
		t.Error(err)
	}
	if a.Id() != aggregateId {
		t.Errorf("expected id to be %s but was %s", aggregateId, a.Id())
	}
}

func TestSave(t *testing.T) {
	bus := event.NewBus()
	store := core.NewStore(bus)
	factory := aggregate.NewFactory()
	err := factory.Register((*aggregate.Aggregate)(nil), func(id uuid.UUID) aggregate.IAggregate {
		return aggregate.NewAggregate(id)
	})
	if err != nil {
		t.Error(err)
	}
	handler := &TestEventHandler{}
	bus.RegisterHandler((*TestEvent)(nil), handler)
	aggregateId := uuid.New()
	a := aggregate.NewAggregate(aggregateId)
	a.RegisterHandler((*TestEvent)(nil), testFunc)
	e := event.NewEvent(uuid.New(), aggregateId, TestEvent{})
	a.TrackChange(e)
	repository := NewRepository(store, bus, factory)
	err = repository.Save(a)
	if err != nil {
		t.Error(err)
	}
	events := store.Load(aggregateId)
	if events[0] != e {
		t.Errorf("expected the only store event to be %T but was %T", e, events[0])
	}
	if handler.event != e {
		t.Errorf("expected event %T to have been handled but was %T", e, handler.event)
	}
}

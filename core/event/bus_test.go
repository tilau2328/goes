package event

import (
	"github.com/google/uuid"
	"testing"
)

func TestNewEventBus(t *testing.T) {
	size := len(NewBus().Handlers())
	if size != 0 {
		t.Errorf("expected handlers map to start empty but has %d", size)
	}
}

func TestBus_RegisterHandler(t *testing.T) {
	bus := NewBus()
	handler := TestEventHandler{}
	const handlerTypeName = "event.TestEvent"
	bus.RegisterHandler((*TestEvent)(nil), &handler)
	size := len(bus.Handlers())
	if size != 1 {
		t.Errorf("expected handlers map to have one handler registered but has %d", size)
	}
	if bus.Handler(handlerTypeName) == nil {
		t.Errorf("expected handler %s could not be found", handlerTypeName)
	}
	expectedId := uuid.New()
	expectedEvent := TestEvent{}
	expectedAggregateId := uuid.New()
	event := NewEvent(expectedId, expectedAggregateId, expectedEvent)
	result, err := bus.Handle(event)
	if err != nil {
		t.Error(err)
	}
	if handler.event != event {
		t.Errorf("expected event to have be %T but was %T times", event, handler.event)
	}
	if result != ExpectedHandlerResult {
		t.Errorf("expected result to be %s but was %s", ExpectedHandlerResult, result)
	}
}

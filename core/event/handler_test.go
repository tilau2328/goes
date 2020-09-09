package event

import (
	"github.com/google/uuid"
	"testing"
)

var ExpectedHandlerResult = "test"

type TestEventHandler struct {
	event IEvent
}

func (t *TestEventHandler) Handle(event IEvent) (interface{}, error) {
	t.event = event
	return ExpectedHandlerResult, nil
}

func (t *Handler) Handle(event IEvent) (interface{}, error) {
	return t.next.Handle(event)
}

func TestChainedHandleEvent(t *testing.T) {
	handler := TestEventHandler{}
	message := TestEvent{"test"}
	chainHandler := Handler{next: &handler}
	event := NewEvent(uuid.New(), uuid.New(), message)
	result, err := chainHandler.Handle(event)
	if err != nil {
		t.Error(err)
	}
	if handler.event != event {
		t.Errorf("expected event to be %T but is %T", event, handler.event)
	}
	if result != ExpectedHandlerResult {
		t.Errorf("expected result to be %s but was %s", ExpectedHandlerResult, result)
	}
}

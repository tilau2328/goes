package event

import (
	"github.com/google/uuid"
	"testing"
)

type TestEventHandler struct {
	event IEvent
}

func (t *TestEventHandler) Handle(event IEvent) error {
	t.event = event
	return nil
}

func (t *Handler) Handle(event IEvent) error {
	return t.next.Handle(event)
}

func TestChainedHandleEvent(t *testing.T) {
	handler := TestEventHandler{}
	message := TestEvent{"test"}
	chainHandler := Handler{next: &handler}
	event := NewEvent(uuid.New(), uuid.New(), message)
	err := chainHandler.Handle(event)
	if err != nil {
		t.Error(err)
	}
	if handler.event != event {
		t.Errorf("expected event to be %T but is %T", event, handler.event)
	}
}
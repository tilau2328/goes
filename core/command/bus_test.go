package command

import (
	"github.com/google/uuid"
	"testing"
)

func TestNewCommandBus(t *testing.T) {
	size := len(NewBus().Handlers())
	if size != 0 {
		t.Errorf("expected handlers map to start empty but has %d", size)
	}
}

func TestBus_RegisterHandler(t *testing.T) {
	bus := NewBus()
	handler := TestCommandHandler{}
	const handlerTypeName = "command.TestCommand"
	bus.RegisterHandler((*TestCommand)(nil), &handler)
	size := len(bus.Handlers())
	if size != 1 {
		t.Errorf("expected handlers map to have one handler registered but has %d", size)
	}
	if bus.Handler(handlerTypeName) == nil {
		t.Errorf("expected handler %s could not be found", handlerTypeName)
	}
	expectedId := uuid.New()
	expectedCommand := TestCommand{}
	expectedAggregateId := uuid.New()
	command := NewCommand(expectedId, expectedAggregateId, expectedCommand)
	result, err := bus.Handle(command)
	if err != nil {
		t.Error(err)
	}
	if result != ExpectedHandlerResult {
		t.Errorf("expected result to be %s but was %s", ExpectedHandlerResult, result)
	}
	if handler.command != command {
		t.Errorf("expected command to be %T but was %T times", command, handler.command)
	}
}

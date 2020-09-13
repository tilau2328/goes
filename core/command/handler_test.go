package command

import (
	"github.com/google/uuid"
	"testing"
)

var ExpectedHandlerResult = "test"

type TestCommandHandler struct {
	command ICommand
}

func (t *TestCommandHandler) Handle(command ICommand) (interface{}, error) {
	t.command = command
	return ExpectedHandlerResult, nil
}

func (t *Handler) Handle(command ICommand) (interface{}, error) {
	return t.next.Handle(command)
}

func TestChainedHandleCommand(t *testing.T) {
	handler := TestCommandHandler{}
	message := TestCommand{"test"}
	chainHandler := Handler{next: &handler}
	command := NewCommand(uuid.New(), uuid.New(), message)
	result, err := chainHandler.Handle(command)
	if err != nil {
		t.Error(err)
	}
	if handler.command != command {
		t.Errorf("expected Command to be %T but was %T", command, handler.command)
	}
	if result != ExpectedHandlerResult {
		t.Errorf("expected result to be %s but was %s", ExpectedHandlerResult, result)
	}
}

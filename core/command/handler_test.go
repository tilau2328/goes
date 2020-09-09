package command

import (
	"github.com/google/uuid"
	"testing"
)

type TestCommandHandler struct {
	command ICommand
}

func (t *TestCommandHandler) Handle(command ICommand) error {
	t.command = command
	return nil
}

func (t *Handler) Handle(command ICommand) error {
	return t.next.Handle(command)
}

func TestChainedHandleCommand(t *testing.T) {
	handler := TestCommandHandler{}
	message := TestCommand{"test"}
	chainHandler := Handler{next: &handler}
	command := NewCommand(uuid.New(), uuid.New(), message)
	err := chainHandler.Handle(command)
	if err != nil {
		t.Error(err)
	}
	if handler.command != command {
		t.Errorf("expected command to be %T but is %T", command, handler.command)
	}
}

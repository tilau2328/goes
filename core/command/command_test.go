package command

import (
	"github.com/google/uuid"
	"testing"
)

type TestCommand struct { test string }

func TestNewCommand(t *testing.T) {
	const expectedCommandType = "command.TestCommand"
	expectedMessage := TestCommand{"test"}
	expectedAggregateId := uuid.New()
	expectedId := uuid.New()

	result := NewCommand(expectedId, expectedAggregateId, expectedMessage)
	aggregateId := result.AggregateId()
	if aggregateId != expectedAggregateId {
		t.Errorf("expected aggregateId to be %s but is %s", expectedAggregateId, aggregateId)
	}

	id := result.Id()
	if id != expectedId {
		t.Errorf("expected id to be %s but is %s", expectedAggregateId, id)
	}

	commandType := result.Type()
	if commandType != expectedCommandType {
		t.Errorf("expected command type to be %s but is %s", expectedCommandType, commandType)
	}

	resultMessage := result.Message()
	if resultMessage != expectedMessage {
		t.Errorf("expected message to be %s but is %s", expectedMessage, resultMessage)
	}
}
package event

import (
	"github.com/google/uuid"
	"testing"
)

type TestEvent struct { test string }

func TestNewEvent(t *testing.T) {
	const expectedEventType = "command.TestEvent"
	expectedMessage := TestEvent{"test"}
	expectedAggregateId := uuid.New()
	expectedId := uuid.New()

	result := NewEvent(expectedId, expectedAggregateId, expectedMessage)
	aggregateId := result.AggregateId()
	if aggregateId != expectedAggregateId {
		t.Errorf("expected aggregateId to be %s but is %s", expectedAggregateId, aggregateId)
	}

	id := result.Id()
	if id != expectedId {
		t.Errorf("expected id to be %s but is %s", expectedAggregateId, id)
	}

	commandType := result.Type()
	if commandType != expectedEventType {
		t.Errorf("expected command type to be %s but is %s", expectedEventType, commandType)
	}

	resultMessage := result.Message()
	if resultMessage != expectedMessage {
		t.Errorf("expected message to be %s but is %s", expectedMessage, resultMessage)
	}
}
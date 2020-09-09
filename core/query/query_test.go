package query

import (
	"github.com/google/uuid"
	"testing"
)

type TestQuery struct { test string }

func TestNewQuery(t *testing.T) {
	const expectedQueryType = "query.TestQuery"
	expectedMessage := TestQuery{"test"}
	expectedAggregateId := uuid.New()
	expectedId := uuid.New()

	result := NewQuery(expectedId, expectedAggregateId, expectedMessage)
	aggregateId := result.AggregateId()
	if aggregateId != expectedAggregateId {
		t.Errorf("expected aggregateId to be %s but is %s", expectedAggregateId, aggregateId)
	}

	id := result.Id()
	if id != expectedId {
		t.Errorf("expected id to be %s but is %s", expectedAggregateId, id)
	}

	queryType := result.Type()
	if queryType != expectedQueryType {
		t.Errorf("expected query type to be %s but is %s", expectedQueryType, queryType)
	}

	resultMessage := result.Message()
	if resultMessage != expectedMessage {
		t.Errorf("expected message to be %s but is %s", expectedMessage, resultMessage)
	}
}
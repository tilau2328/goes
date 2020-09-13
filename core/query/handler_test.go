package query

import (
	"github.com/google/uuid"
	"testing"
)

var ExpectedHandlerResult = "test"

type TestQueryHandler struct {
	query IQuery
}

func (t *TestQueryHandler) Handle(query IQuery) (interface{}, error) {
	t.query = query
	return ExpectedHandlerResult, nil
}

func (t *Handler) Handle(query IQuery) (interface{}, error) {
	return t.next.Handle(query)
}

func TestChainedHandleQuery(t *testing.T) {
	handler := TestQueryHandler{}
	message := TestQuery{"test"}
	chainHandler := Handler{next: &handler}
	query := NewQuery(uuid.New(), uuid.New(), message)
	result, err := chainHandler.Handle(query)
	if err != nil {
		t.Error(err)
	}
	if handler.query != query {
		t.Errorf("expected Query to be %T but is %T", query, handler.query)
	}
	if result != ExpectedHandlerResult {
		t.Errorf("expected result to be %s but was %s", ExpectedHandlerResult, result)
	}
}

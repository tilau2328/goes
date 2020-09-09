package query

import (
	"github.com/google/uuid"
	"testing"
)

type TestQueryHandler struct {
	query IQuery
}

func (t *TestQueryHandler) Handle(query IQuery) error {
	t.query = query
	return nil
}

func (t *Handler) Handle(query IQuery) error {
	return t.next.Handle(query)
}

func TestChainedHandleQuery(t *testing.T) {
	handler := TestQueryHandler{}
	message := TestQuery{"test"}
	chainHandler := Handler{next: &handler}
	query := NewQuery(uuid.New(), uuid.New(), message)
	err := chainHandler.Handle(query)
	if err != nil {
		t.Error(err)
	}
	if handler.query != query {
		t.Errorf("expected query to be %T but is %T", query, handler.query)
	}
}

package query

import (
	"github.com/google/uuid"
	"testing"
)

func TestNewQueryBus(t *testing.T) {
	size := len(NewBus().Handlers())
	if size != 0 {
		t.Errorf("expected handlers map to start empty but has %d", size)
	}
}

func TestBus_RegisterHandler(t *testing.T) {
	bus := NewBus()
	handler := TestQueryHandler{}
	const handlerTypeName = "query.TestQuery"
	bus.RegisterHandler((*TestQuery)(nil), &handler)
	size := len(bus.Handlers())
	if size != 1 {
		t.Errorf("expected handlers map to have one handler registered but has %d", size)
	}
	if bus.Handler(handlerTypeName) == nil {
		t.Errorf("expected handler %s could not be found", handlerTypeName)
	}
	expectedId := uuid.New()
	expectedQuery := TestQuery{}
	expectedAggregateId := uuid.New()
	query := NewQuery(expectedId, expectedAggregateId, expectedQuery)
	err := bus.Handle(query)
	if err != nil {
		t.Error(err)
	}
	if handler.query != query {
		t.Errorf("expected query to have be %T but was %T times", query, handler.query)
	}
}

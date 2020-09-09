package query

import (
	"github.com/google/uuid"
	"github.com/tilau2328/goes"
)

type IQuery interface {
	AggregateId() uuid.UUID
	Id() uuid.UUID
	Type() string
	Message() interface{}
}

type Query struct {
	aggregate uuid.UUID
	id        uuid.UUID
	query     interface{}
}

func NewQuery(id uuid.UUID, aggregate uuid.UUID, query interface{}) *Query {
	return &Query{
		id:        id,
		query:   query,
		aggregate: aggregate,
	}
}

func (c *Query) Id() uuid.UUID {
	return c.id
}

func (c *Query) Type() string {
	return goes.MessageType(c.query)
}

func (c *Query) Message() interface{} {
	return c.query
}

func (c *Query) AggregateId() uuid.UUID {
	return c.aggregate
}

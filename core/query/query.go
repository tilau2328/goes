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
	Aggregate uuid.UUID   `json:"aggregateId"`
	QueryId   uuid.UUID   `json:"queryId"`
	Query     interface{} `json:"message"`
}

func NewQuery(id uuid.UUID, aggregate uuid.UUID, query interface{}) *Query {
	return &Query{aggregate, id, query}
}

func (c *Query) Id() uuid.UUID {
	return c.QueryId
}

func (c *Query) Type() string {
	return goes.MessageType(c.Query)
}

func (c *Query) Message() interface{} {
	return c.Query
}

func (c *Query) AggregateId() uuid.UUID {
	return c.Aggregate
}

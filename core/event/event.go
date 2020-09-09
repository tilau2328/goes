package event

import (
	"github.com/google/uuid"
	"github.com/tilau2328/goes"
)

type IEvent interface {
	AggregateId() uuid.UUID
	Id() uuid.UUID
	Type() string
	Message() interface{}
}

type Event struct {
	aggregate uuid.UUID
	id        uuid.UUID
	event     interface{}
}

func NewEvent(id uuid.UUID, aggregateId uuid.UUID, event interface{}) *Event {
	return &Event{
		id:        id,
		event:     event,
		aggregate: aggregateId,
	}
}

func (c *Event) Id() uuid.UUID {
	return c.id
}

func (c *Event) Type() string {
	return goes.MessageType(c.event)
}

func (c *Event) Message() interface{} {
	return c.event
}

func (c *Event) AggregateId() uuid.UUID {
	return c.aggregate
}

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
	Aggregate uuid.UUID   `json:"aggregateId"`
	EventId   uuid.UUID   `json:"eventId"`
	Event     interface{} `json:"message"`
}

func NewEvent(id uuid.UUID, aggregateId uuid.UUID, event interface{}) *Event {
	return &Event{aggregateId, id, event}
}

func (c *Event) Id() uuid.UUID {
	return c.EventId
}

func (c *Event) Type() string {
	return goes.MessageType(c.Event)
}

func (c *Event) Message() interface{} {
	return c.Event
}

func (c *Event) AggregateId() uuid.UUID {
	return c.Aggregate
}

package store

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/tilau2328/goes/core/event"
)

type IEventStore interface {
	Store([]event.IEvent) error
	Load(id uuid.UUID) []event.IEvent
}

type EventStore struct {
	bus    event.IEventBus
	events map[string][]event.IEvent
}

func NewStore(bus event.IEventBus) *EventStore {
	return &EventStore{bus, make(map[string][]event.IEvent)}
}

func (s *EventStore) Store(events []event.IEvent) error {
	for _, e := range events {
		id := e.AggregateId().String()
		if s.events[id] == nil {
			s.events[id] = []event.IEvent{e}
		} else {
			s.events[id] = append(s.events[id], e)
		}
		_, err := s.bus.Handle(e)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}

func (s *EventStore) Load(id uuid.UUID) []event.IEvent {
	return s.events[id.String()]
}

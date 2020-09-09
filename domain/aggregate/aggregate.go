package aggregate

import (
	"fmt"
	"github.com/google/uuid"
	"goes"
	"goes/core/event"
)

type IAggregate interface {
	Id() uuid.UUID
	ClearChanges()
	IncrementVersion()
	CurrentVersion() int32
	OriginalVersion() int32
	TrackChange(event.IEvent)
	GetChanges() []event.IEvent
	Apply(event.IEvent, bool) error
}

type Aggregate struct {
	id      uuid.UUID
	version int32
	changes []event.IEvent
	handlers map[string]func(event.IEvent, bool)
}

func NewAggregate(id uuid.UUID) *Aggregate {
	return &Aggregate{
		handlers: make(map[string]func(event.IEvent, bool)),
		changes: []event.IEvent{},
		id:      id,
		version: -1,
	}
}

func (a *Aggregate) Id() uuid.UUID {
	return a.id
}

func (a *Aggregate) OriginalVersion() int32 {
	return a.version
}

func (a *Aggregate) CurrentVersion() int32 {
	return a.version + int32(len(a.changes))
}

func (a *Aggregate) IncrementVersion() {
	a.version++
}

func (a *Aggregate) TrackChange(event event.IEvent) {
	a.changes = append(a.changes, event)
}

func (a *Aggregate) GetChanges() []event.IEvent {
	return a.changes
}

func (a *Aggregate) ClearChanges() {
	a.changes = []event.IEvent{}
}

func (a *Aggregate) Apply(event event.IEvent, isNew bool) error {
	eventType := event.Type()
	handler := a.handlers[eventType]
	if handler == nil {
		return fmt.Errorf("no handler found for event type %s", eventType)
	}
	handler(event, isNew)
	return nil
}

func (a *Aggregate) RegisterHandler(eventType interface{}, handler func(event.IEvent, bool)) {
	a.handlers[goes.MessageType(eventType)] = handler
}

package aggregate

import (
	"fmt"
	"github.com/google/uuid"
	"goes"
)

type IAggregateFactory interface {
	Aggregate(typeName string, id uuid.UUID) IAggregate
	Register(IAggregate, func(uuid.UUID) IAggregate) error
}

type Factory struct {
	delegates map[string]func(uuid.UUID) IAggregate
}

func NewFactory() *Factory {
	return &Factory{delegates: make(map[string]func(uuid.UUID) IAggregate)}
}

func (t *Factory) Register(aggregate IAggregate, delegate func(uuid uuid.UUID) IAggregate) error {
	typeName := goes.MessageType(aggregate)
	if _, ok := t.delegates[typeName]; ok {
		return fmt.Errorf("factory delegate already registered for type: \"%s\"", typeName)
	}
	t.delegates[typeName] = delegate
	return nil
}

func (t *Factory) Aggregate(typeName string, id uuid.UUID) IAggregate {
	if f, ok := t.delegates[typeName]; ok {
		return f(id)
	}
	return nil
}
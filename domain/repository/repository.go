package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/tilau2328/goes"
	"github.com/tilau2328/goes/core/event"
	"github.com/tilau2328/goes/core/store"
	"github.com/tilau2328/goes/domain/aggregate"
)

type IRepository interface {
	Save(aggregate.IAggregate) error
	Load(interface{}, uuid.UUID) (aggregate.IAggregate, error)
}

type Repository struct {
	store   store.IEventStore
	bus     event.IEventBus
	factory aggregate.IAggregateFactory
}

func NewRepository(store store.IEventStore, bus event.IEventBus, factory aggregate.IAggregateFactory) *Repository {
	return &Repository{store, bus, factory}
}

func (r *Repository) Save(aggregate aggregate.IAggregate) error {
	err := r.store.Store(aggregate.GetChanges())
	if err != nil {
		return err
	}
	aggregate.ClearChanges()
	return nil
}

func (r *Repository) Load(aggregateType interface{}, id uuid.UUID) (aggregate.IAggregate, error) {
	a := r.factory.Aggregate(goes.MessageType(aggregateType), id)
	if a == nil {
		return nil, fmt.Errorf("unable to find aggregate for type %s", aggregateType)
	}
	var err error
	for _, e := range r.store.Load(id) {
		err = a.Apply(e, false)
		if err != nil {
			return nil, err
		}
	}
	return a, nil
}

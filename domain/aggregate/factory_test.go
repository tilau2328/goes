package aggregate

import (
	"github.com/google/uuid"
	"goes/core/event"
	"testing"
)

func (*Aggregate) Apply(event.IEvent, bool) {}
func delegate(id uuid.UUID) IAggregate      { return NewAggregate(id) }

func TestFactory_Register(t *testing.T) {
	factory := NewFactory()
	err := factory.Register((*Aggregate)(nil), delegate)
	if err != nil {
		t.Error(err)
	}
	size := len(factory.delegates)
	if size != 1 {
		t.Errorf("expected delegates map to have one handler registered but has %d", size)
	}
}

func TestFactory_RegisterDuplicate(t *testing.T) {
	factory := NewFactory()
	err := factory.Register((*Aggregate)(nil), delegate)
	if err != nil {
		t.Errorf("failed to create register aggregate")
	}
	err = factory.Register((*Aggregate)(nil), delegate)
	if err == nil {
		t.Errorf("expected duplicate registration to fail")
	}
}

func TestFactory_Aggregate(t *testing.T) {
	const typeName = "test"
	factory := &Factory{ delegates: map[string]func(uuid.UUID) IAggregate{ typeName: delegate} }
	aggregate := factory.Aggregate(typeName, uuid.New())
	if aggregate == nil {
		t.Errorf("failed to create aggregate")
	}
}

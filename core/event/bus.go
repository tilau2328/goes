package event

import "github.com/tilau2328/goes"

type IEventBus interface {
	RegisterHandler(interface{}, IEventHandler)
	Handlers() map[string]IEventHandler
	Handler(string) IEventHandler
	Handle(IEvent) error
}

type Bus struct {
	handlers map[string]IEventHandler
}

func NewBus() *Bus {
	return &Bus{handlers: make(map[string]IEventHandler)}
}

func (b *Bus) RegisterHandler(messageType interface{}, handler IEventHandler) {
	b.handlers[goes.MessageType(messageType)] = handler
}

func (b *Bus) Handler(handler string) IEventHandler {
	return b.handlers[handler]
}

func (b *Bus) Handlers() map[string]IEventHandler {
	return b.handlers
}

func (b *Bus) Handle(event IEvent) error {
	return b.Handler(event.Type()).Handle(event)
}
package query

import "github.com/tilau2328/goes"

type IQueryBus interface {
	RegisterHandler(interface{}, IQueryHandler)
	Handlers() map[string]IQueryHandler
	Handler(string) IQueryHandler
	Handle(IQuery) (interface{}, error)
}

type Bus struct {
	handlers map[string]IQueryHandler
}

func NewBus() *Bus {
	return &Bus{handlers: make(map[string]IQueryHandler)}
}

func (b *Bus) RegisterHandler(messageType interface{}, handler IQueryHandler) {
	b.handlers[goes.MessageType(messageType)] = handler
}

func (b *Bus) Handler(handler string) IQueryHandler {
	return b.handlers[handler]
}

func (b *Bus) Handlers() map[string]IQueryHandler {
	return b.handlers
}

func (b *Bus) Handle(query IQuery) (interface{}, error) {
	return b.Handler(query.Type()).Handle(query)
}

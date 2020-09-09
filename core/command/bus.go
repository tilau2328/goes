package command

import "github.com/tilau2328/goes"

type ICommandBus interface {
	RegisterHandler(interface{}, ICommandHandler)
	Handlers() map[string]ICommandHandler
	Handler(string) ICommandHandler
	Handle(ICommand) (interface{}, error)
}

type Bus struct {
	handlers map[string]ICommandHandler
}

func NewBus() *Bus {
	return &Bus{handlers: make(map[string]ICommandHandler)}
}

func (b *Bus) RegisterHandler(messageType interface{}, handler ICommandHandler) {
	b.handlers[goes.MessageType(messageType)] = handler
}

func (b *Bus) Handler(handler string) ICommandHandler {
	return b.handlers[handler]
}

func (b *Bus) Handlers() map[string]ICommandHandler {
	return b.handlers
}

func (b *Bus) Handle(command ICommand) (interface{}, error) {
	return b.Handler(command.Type()).Handle(command)
}

package event

type IEventHandler interface {
	Handle(IEvent) error
}

type Handler struct {
	next IEventHandler
}
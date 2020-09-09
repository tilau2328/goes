package event

type IEventHandler interface {
	Handle(IEvent) (interface{}, error)
}

type Handler struct {
	next IEventHandler
}

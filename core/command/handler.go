package command

type ICommandHandler interface {
	Handle(ICommand) (interface{}, error)
}

type Handler struct {
	next ICommandHandler
}

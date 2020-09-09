package command

type ICommandHandler interface {
	Handle(ICommand) error
}

type Handler struct {
	next ICommandHandler
}
package query

type IQueryHandler interface {
	Handle(IQuery) error
}
type Handler struct {
	next IQueryHandler
}
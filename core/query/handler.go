package query

type IQueryHandler interface {
	Handle(IQuery) (interface{}, error)
}
type Handler struct {
	next IQueryHandler
}

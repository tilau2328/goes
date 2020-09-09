package domain

type IState interface {
	Write(string, interface{})
	Read(string) (interface{}, error)
}

type State struct {
	values map[string]interface{}
}

func NewState() *State {
	return &State{ values: make(map[string]interface{}) }
}

func (s *State) Write(key string, value interface{}) {
	s.values[key] = value
}

func (s *State) Read(key string) interface{} {
	return s.values[key]
}

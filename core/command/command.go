package command

import (
	"github.com/google/uuid"
	"github.com/tilau2328/goes"
)

type ICommand interface {
	AggregateId() uuid.UUID
	Id() uuid.UUID
	Type() string
	Message() interface{}
}

type Command struct {
	Aggregate uuid.UUID   `json:"aggregateId"`
	CommandId uuid.UUID   `json:"commandId"`
	Command   interface{} `json:"message"`
}

func NewCommand(id uuid.UUID, aggregate uuid.UUID, command interface{}) *Command {
	return &Command{aggregate, id, command}
}

func (c *Command) Id() uuid.UUID {
	return c.CommandId
}

func (c *Command) Type() string {
	return goes.MessageType(c.Command)
}

func (c *Command) Message() interface{} {
	return c.Command
}

func (c *Command) AggregateId() uuid.UUID {
	return c.Aggregate
}

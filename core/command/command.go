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
	aggregate uuid.UUID
	id        uuid.UUID
	command   interface{}
}

func NewCommand(id uuid.UUID, aggregate uuid.UUID, command interface{}) *Command {
	return &Command{
		id:        id,
		command:   command,
		aggregate: aggregate,
	}
}

func (c *Command) Id() uuid.UUID {
	return c.id
}

func (c *Command) Type() string {
	return goes.MessageType(c.command)
}

func (c *Command) Message() interface{} {
	return c.command
}

func (c *Command) AggregateId() uuid.UUID {
	return c.aggregate
}

package inmemory

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
)

// CommandBus is an in-memory implementation of the command.Bus interface.
type CommandBus struct {
	handlers map[command.Type]command.Handler
}

// NewCommandBus creates a new instance of CommandBus.
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

// Dispatch executes a command by finding the appropriate handler and invoking it.
func (b *CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	handler, ok := b.handlers[cmd.Type()]
	if !ok {
		return nil
	}

	return handler.Handle(ctx, cmd)
}

// Register adds a command handler for a specific command type.
func (b *CommandBus) Register(cmdType command.Type, handler command.Handler) {
	b.handlers[cmdType] = handler
}

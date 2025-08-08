package command

import "context"

// Bus is the interface for a command bus that dispatches commands.
type Bus interface {
	// Dispatch processes a command in the given context.
	Dispatch(context.Context, Command) error
	// Register registers a command handler for a specific command type.
	Register(Type, Handler)
}

//go:generate mockery --name=Bus --output=commandmocks --case=snake --outpkg=commandmocks

// Type represents the type of a command.
type Type string

// Command represents an application command that can be dispatched.
type Command interface {
	Type() Type
}

// Handler defines the expected behavior for a command handler.
type Handler interface {
	Handle(context.Context, Command) error
}

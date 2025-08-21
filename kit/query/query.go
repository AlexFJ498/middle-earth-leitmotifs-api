package query

import "context"

// Bus is the interface for a query bus that dispatches querys.
type Bus interface {
	// Ask processes a query in the given context.
	Ask(ctx context.Context, query Query) (any, error)
	// Register registers a query handler for a specific query type.
	Register(Type, Handler)
}

//go:generate mockery --name=Bus --output=querymocks --case=snake --outpkg=querymocks

// Type represents the type of a query.
type Type string

// Query represents an application query that can be dispatched.
type Query interface {
	Type() Type
}

// Handler defines the expected behavior for a query handler.
type Handler interface {
	Handle(context.Context, Query) (any, error)
}

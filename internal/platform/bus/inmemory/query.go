package inmemory

import (
	"context"
	"fmt"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
)

// QueryBus is an in-memory implementation of the query bus.
type QueryBus struct {
	handlers map[query.Type]query.Handler
}

// NewQueryBus creates a new instance of QueryBus.
func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[query.Type]query.Handler),
	}
}

// Ask processes a query and returns the result.
func (b *QueryBus) Ask(ctx context.Context, query query.Query) (any, error) {
	handler, ok := b.handlers[query.Type()]
	if !ok {
		return nil, fmt.Errorf("no handler registered for query type: %s", query.Type())
	}
	return handler.Handle(ctx, query)
}

// Register registers a query handler for a specific query type.
func (b *QueryBus) Register(queryType query.Type, handler query.Handler) {
	b.handlers[queryType] = handler
}

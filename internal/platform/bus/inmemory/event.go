package inmemory

import (
	"context"
	"sync"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/event"
)

type EventBus struct {
	handlers map[event.Type][]event.Handler
	mu       sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[event.Type][]event.Handler),
	}
}

func (b *EventBus) Publish(ctx context.Context, events []event.Event) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, e := range events {
		if handlers, ok := b.handlers[e.Type()]; ok {
			for _, h := range handlers {
				if err := h.Handle(ctx, e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (b *EventBus) Subscribe(eventType event.Type, handler event.Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()

	subscribersForType, ok := b.handlers[eventType]
	if !ok {
		b.handlers[eventType] = []event.Handler{handler}
		return
	}

	b.handlers[eventType] = append(subscribersForType, handler)
}

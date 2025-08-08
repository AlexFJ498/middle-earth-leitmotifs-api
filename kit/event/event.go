package event

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Bus is the interface for an event bus that dispatches events.
type Bus interface {
	Publish(context.Context, []Event) error
	Subscribe(Type, Handler)
}

//go:generate mockery --name=Bus --output=eventmocks --case=snake --outpkg=eventmocks

// Handler defines the expected behavior for an event handler.
type Handler interface {
	Handle(context.Context, Event) error
}

// Type represents the type of an event.
type Type string

// Event represents an application event that can be published.
type Event interface {
	ID() string
	AggregateID() string
	OccurredOn() time.Time
	Type() Type
}

// BaseEvent is a base implementation of the Event interface.
type BaseEvent struct {
	eventID     string
	aggregateID string
	occurredOn  time.Time
}

// NewBaseEvent creates a new BaseEvent instance.
func NewBaseEvent(aggregateID string) BaseEvent {
	return BaseEvent{
		eventID:     uuid.New().String(),
		aggregateID: aggregateID,
		occurredOn:  time.Now(),
	}
}

func (e BaseEvent) ID() string {
	return e.eventID
}

func (e BaseEvent) AggregateID() string {
	return e.aggregateID
}

func (e BaseEvent) OccurredOn() time.Time {
	return e.occurredOn
}

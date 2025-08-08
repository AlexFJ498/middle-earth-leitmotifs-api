package creating

import (
	"context"
	"errors"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/increasing"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/event"
)

// At the moment, this is not called. It shows how an inmemory event bus can be used to handle events.
type IncreaseUsersCounterOnUserCreated struct {
	increaserService increasing.UserCounterIncreaserService
}

func NewIncreaseUsersCounterOnUserCreated(increaserService increasing.UserCounterIncreaserService) IncreaseUsersCounterOnUserCreated {
	return IncreaseUsersCounterOnUserCreated{
		increaserService: increaserService,
	}
}

func (e IncreaseUsersCounterOnUserCreated) Handle(ctx context.Context, event event.Event) error {
	userCreatedEvent, ok := event.(domain.UserCreatedEvent)
	if !ok {
		return errors.New("event is not of type UserCreatedEvent")
	}

	return e.increaserService.Increase(userCreatedEvent.ID())
}

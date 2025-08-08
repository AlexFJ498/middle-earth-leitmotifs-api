package domain

import "github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/event"

const UserCreatedEventType = "events.user.created"

type UserCreatedEvent struct {
	event.BaseEvent
	id    string
	name  string
	email string
}

func NewUserCreatedEvent(id, name, email string) UserCreatedEvent {
	return UserCreatedEvent{
		id:        id,
		name:      name,
		email:     email,
		BaseEvent: event.NewBaseEvent(id),
	}
}

func (e UserCreatedEvent) Type() event.Type {
	return UserCreatedEventType
}

func (e UserCreatedEvent) UserID() string {
	return e.id
}

func (e UserCreatedEvent) UserName() string {
	return e.name
}

func (e UserCreatedEvent) UserEmail() string {
	return e.email
}

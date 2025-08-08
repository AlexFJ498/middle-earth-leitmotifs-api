package domain

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/event"
	"github.com/google/uuid"
)

var ErrInvalidUserID = errors.New("invalid user ID")
var ErrInvalidUserName = errors.New("invalid user name")
var ErrInvalidUserEmail = errors.New("invalid user email")

// UserID represents the unique identifier for a user.
type UserID struct {
	value string
}

// UserName represents the name of a user.
type UserName struct {
	value string
}

// Email represents the email of a user.
type UserEmail struct {
	value string
}

// NewUserID creates a new UserID instance.
func NewUserID(value string) (UserID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return UserID{}, fmt.Errorf("%w: %w", ErrInvalidUserID, err)
	}

	return UserID{
		value: v.String(),
	}, nil
}

// String returns the string representation of the UserID.
func (id UserID) String() string {
	return id.value
}

// NewUserName creates a new UserName instance.
func NewUserName(value string) (UserName, error) {
	if value == "" {
		return UserName{}, ErrInvalidUserName
	}

	return UserName{
		value: value,
	}, nil
}

// String returns the string representation of the UserName.
func (name UserName) String() string {
	return name.value
}

// NewUserEmail creates a new UserEmail instance.
func NewUserEmail(value string) (UserEmail, error) {
	if value == "" {
		return UserEmail{}, ErrInvalidUserEmail
	}

	// Validate email format
	if _, err := mail.ParseAddress(value); err != nil {
		return UserEmail{}, fmt.Errorf("%w: %s", ErrInvalidUserEmail, err)
	}

	return UserEmail{
		value: value,
	}, nil
}

// String returns the string representation of the UserEmail.
func (email UserEmail) String() string {
	return email.value
}

// UserRepository defines the interface for user persistence operations.
type UserRepository interface {
	Save(ctx context.Context, user User) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=UserRepository

// User represents a user in the system.
type User struct {
	id    UserID
	name  UserName
	email UserEmail

	events []event.Event
}

// NewUser creates a new User instance.
func NewUser(id, name, email string) (User, error) {
	idVO, err := NewUserID(id)
	if err != nil {
		return User{}, err
	}

	nameVO, err := NewUserName(name)
	if err != nil {
		return User{}, err
	}

	emailVO, err := NewUserEmail(email)
	if err != nil {
		return User{}, err
	}

	user := User{
		id:    idVO,
		name:  nameVO,
		email: emailVO,
	}

	user.Record(NewUserCreatedEvent(idVO.String(), nameVO.String(), emailVO.String()))

	return user, nil
}

// ID returns the user's ID.
func (u User) ID() UserID {
	return u.id
}

// Name returns the user's name.
func (u User) Name() UserName {
	return u.name
}

// Email returns the user's email.
func (u User) Email() UserEmail {
	return u.email
}

// Record adds an event to the user's event list.
func (u *User) Record(event event.Event) {
	u.events = append(u.events, event)
}

// PullEvents returns the events recorded for the user and clears the event list.
func (u *User) PullEvents() []event.Event {
	events := u.events
	// Clear events after returning to prevent re-publishing
	u.events = nil

	return events
}

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
var ErrInvalidUserPassword = errors.New("invalid user password")
var ErrUserAlreadyExists = errors.New("user already exists")
var ErrUserNotFound = errors.New("user not found")

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

// Password represents the hashed password of a user.
type UserPassword struct {
	value string
}

// UserIsAdmin represents the admin status of a user.
type UserIsAdmin struct {
	value bool
}

// NewUserID creates a new UserID instance.
func NewUserID() (UserID, error) {
	v, err := uuid.NewRandom()
	if err != nil {
		return UserID{}, fmt.Errorf("%w: %w", ErrInvalidUserID, err)
	}

	return UserID{
		value: v.String(),
	}, nil
}

func NewUserIDFromString(id string) (UserID, error) {
	if id == "" {
		return UserID{}, ErrInvalidUserID
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return UserID{}, ErrInvalidUserID
	}

	return UserID{
		value: id,
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
		return UserEmail{}, ErrInvalidUserEmail
	}

	return UserEmail{
		value: value,
	}, nil
}

// String returns the string representation of the UserEmail.
func (email UserEmail) String() string {
	return email.value
}

// NewUserPassword creates a new UserPassword instance.
func NewUserPassword(value string) (UserPassword, error) {
	if value == "" {
		return UserPassword{}, ErrInvalidUserPassword
	}

	return UserPassword{
		value: value,
	}, nil
}

// String returns the string representation of the UserPassword.
func (password UserPassword) String() string {
	return password.value
}

// NewUserIsAdmin creates a new UserIsAdmin instance.
func NewUserIsAdmin(value bool) (UserIsAdmin, error) {
	return UserIsAdmin{
		value: value,
	}, nil
}

func (isAdmin UserIsAdmin) Bool() bool {
	return isAdmin.value
}

// UserRepository defines the interface for user persistence operations.
type UserRepository interface {
	Save(ctx context.Context, user User) error
	Find(ctx context.Context, id UserID) (User, error)
	FindByEmail(ctx context.Context, email UserEmail) (User, error)
	FindAll(ctx context.Context) ([]User, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=UserRepository

// User represents a user in the system.
type User struct {
	id       UserID
	name     UserName
	email    UserEmail
	password UserPassword
	isAdmin  UserIsAdmin

	events []event.Event
}

// NewUser creates a new User instance.
func NewUser(name, email, password string, isAdmin bool) (User, error) {
	idVO, err := NewUserID()
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

	passwordVO, err := NewUserPassword(password)
	if err != nil {
		return User{}, err
	}

	isAdminVO, err := NewUserIsAdmin(isAdmin)
	if err != nil {
		return User{}, err
	}

	user := User{
		id:       idVO,
		name:     nameVO,
		email:    emailVO,
		password: passwordVO,
		isAdmin:  isAdminVO,
	}

	user.Record(NewUserCreatedEvent(idVO.String(), nameVO.String(), emailVO.String()))

	return user, nil
}

// NewUserWithID creates a new User instance with the given ID.
func NewUserWithID(id, name, email, password string, isAdmin bool) (User, error) {
	idVO, err := NewUserIDFromString(id)
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

	passwordVO, err := NewUserPassword(password)
	if err != nil {
		return User{}, err
	}

	isAdminVO, err := NewUserIsAdmin(isAdmin)
	if err != nil {
		return User{}, err
	}

	user := User{
		id:       idVO,
		name:     nameVO,
		email:    emailVO,
		password: passwordVO,
		isAdmin:  isAdminVO,
	}

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

// Password returns the user's password.
func (u User) Password() UserPassword {
	return u.password
}

// IsAdmin returns the user's admin status.
func (u User) IsAdmin() UserIsAdmin {
	return u.isAdmin
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

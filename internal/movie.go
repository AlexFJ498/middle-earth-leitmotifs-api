package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidMovieID = fmt.Errorf("invalid movie ID")
var ErrInvalidMovieName = fmt.Errorf("invalid movie name")
var ErrMovieNotFound = fmt.Errorf("movie not found")

type MovieID struct {
	value string
}

type MovieName struct {
	value string
}

func NewMovieID() (MovieID, error) {
	v, err := uuid.NewRandom()
	if err != nil {
		return MovieID{}, fmt.Errorf("%w: %w", ErrInvalidMovieID, err)
	}

	return MovieID{
		value: v.String(),
	}, nil
}

func NewMovieIDFromString(id string) (MovieID, error) {
	if id == "" {
		return MovieID{}, ErrInvalidMovieID
	}

	v, err := uuid.Parse(id)
	if err != nil {
		return MovieID{}, ErrInvalidMovieID
	}

	return MovieID{
		value: v.String(),
	}, nil
}

func (id MovieID) String() string {
	return id.value
}

func NewMovieName(value string) (MovieName, error) {
	if value == "" {
		return MovieName{}, ErrInvalidMovieName
	}

	return MovieName{
		value: value,
	}, nil
}

func (n MovieName) String() string {
	return n.value
}

type MovieRepository interface {
	Save(ctx context.Context, movie Movie) error
	Find(ctx context.Context, id MovieID) (Movie, error)
	FindAll(ctx context.Context) ([]Movie, error)
	Delete(ctx context.Context, id MovieID) error
	Update(ctx context.Context, movie Movie) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=MovieRepository

type Movie struct {
	id   MovieID
	name MovieName
}

func NewMovie(name string) (Movie, error) {
	idVO, err := NewMovieID()
	if err != nil {
		return Movie{}, err
	}

	nameVO, err := NewMovieName(name)
	if err != nil {
		return Movie{}, err
	}

	movie := Movie{
		id:   idVO,
		name: nameVO,
	}

	return movie, nil
}

func NewMovieWithID(id, name string) (Movie, error) {
	idVO, err := NewMovieIDFromString(id)
	if err != nil {
		return Movie{}, err
	}

	nameVO, err := NewMovieName(name)
	if err != nil {
		return Movie{}, err
	}

	movie := Movie{
		id:   idVO,
		name: nameVO,
	}

	return movie, nil
}

func (m Movie) ID() MovieID {
	return m.id
}

func (m Movie) Name() MovieName {
	return m.name
}

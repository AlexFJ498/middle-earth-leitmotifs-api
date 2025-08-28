package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidTrackID = fmt.Errorf("invalid track ID")
var ErrInvalidTrackName = fmt.Errorf("invalid track name")
var ErrTrackNotFound = fmt.Errorf("track not found")

type TrackID struct {
	value string
}

type TrackName struct {
	value string
}

func NewTrackID() (TrackID, error) {
	v, err := uuid.NewRandom()
	if err != nil {
		return TrackID{}, fmt.Errorf("%w: %w", ErrInvalidTrackID, err)
	}

	return TrackID{
		value: v.String(),
	}, nil
}

func NewTrackIDFromString(id string) (TrackID, error) {
	if id == "" {
		return TrackID{}, ErrInvalidTrackID
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return TrackID{}, ErrInvalidTrackID
	}

	return TrackID{
		value: id,
	}, nil
}

func (id TrackID) String() string {
	return id.value
}

func NewTrackName(value string) (TrackName, error) {
	if value == "" {
		return TrackName{}, ErrInvalidTrackName
	}

	return TrackName{
		value: value,
	}, nil
}

func (n TrackName) String() string {
	return n.value
}

type TrackRepository interface {
	Save(ctx context.Context, track Track) error
	Find(ctx context.Context, id TrackID) (Track, error)
	FindAll(ctx context.Context) ([]Track, error)
	Delete(ctx context.Context, id TrackID) error
	Update(ctx context.Context, track Track) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=TrackRepository

type Track struct {
	id      TrackID
	name    TrackName
	movieID MovieID
}

func NewTrack(name string, movieID string) (Track, error) {
	idVO, err := NewTrackID()
	if err != nil {
		return Track{}, err
	}

	nameVO, err := NewTrackName(name)
	if err != nil {
		return Track{}, err
	}

	movieIDVO, err := NewMovieIDFromString(movieID)
	if err != nil {
		return Track{}, err
	}

	track := Track{
		id:      idVO,
		name:    nameVO,
		movieID: movieIDVO,
	}

	return track, nil
}

func NewTrackWithID(id, name string, movieID string) (Track, error) {
	idVO, err := NewTrackIDFromString(id)
	if err != nil {
		return Track{}, err
	}

	nameVO, err := NewTrackName(name)
	if err != nil {
		return Track{}, err
	}

	movieIDVO, err := NewMovieIDFromString(movieID)
	if err != nil {
		return Track{}, err
	}

	track := Track{
		id:      idVO,
		name:    nameVO,
		movieID: movieIDVO,
	}

	return track, nil
}

func (t Track) ID() TrackID {
	return t.id
}

func (t Track) Name() TrackName {
	return t.name
}

func (t Track) MovieID() MovieID {
	return t.movieID
}

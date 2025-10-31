package domain

import (
	"context"
	"fmt"
)

var ErrInvalidStartSecond = fmt.Errorf("invalid start second")
var ErrInvalidEndSecond = fmt.Errorf("invalid end second")
var ErrTrackThemeNotFound = fmt.Errorf("track theme not found")

type StartSecond struct {
	value int
}

type EndSecond struct {
	value int
}

type isVariant struct {
	value bool
}

func NewStartSecond(value int) (StartSecond, error) {
	if value < 0 {
		return StartSecond{}, ErrInvalidStartSecond
	}

	return StartSecond{value: value}, nil
}

func (s StartSecond) Int() int {
	return s.value
}

func NewEndSecond(value int) (EndSecond, error) {
	if value < 0 {
		return EndSecond{}, ErrInvalidEndSecond
	}

	return EndSecond{value: value}, nil
}

func (e EndSecond) Int() int {
	return e.value
}

func NewIsVariant(value bool) isVariant {
	return isVariant{value: value}
}

func (i isVariant) Bool() bool {
	return i.value
}

type TrackThemeRepository interface {
	Save(ctx context.Context, trackTheme TrackTheme) error
	Find(ctx context.Context, trackID TrackID, themeID ThemeID, startSecond StartSecond) (TrackTheme, error)
	FindByTrack(ctx context.Context, trackID TrackID) ([]TrackTheme, error)
	Delete(ctx context.Context, trackID TrackID, themeID ThemeID, startSecond StartSecond) error
	Update(ctx context.Context, trackTheme TrackTheme) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=TrackThemeRepository

type TrackTheme struct {
	trackID     TrackID
	themeID     ThemeID
	startSecond StartSecond
	endSecond   EndSecond
	isVariant   isVariant
}

func NewTrackTheme(trackID, themeID string, startSecond, endSecond int, isVariant bool) (TrackTheme, error) {
	trackIDVO, err := NewTrackIDFromString(trackID)
	if err != nil {
		return TrackTheme{}, err
	}

	themeIDVO, err := NewThemeIDFromString(themeID)
	if err != nil {
		return TrackTheme{}, err
	}

	startSecondVO, err := NewStartSecond(startSecond)
	if err != nil {
		return TrackTheme{}, err
	}

	endSecondVO, err := NewEndSecond(endSecond)
	if err != nil {
		return TrackTheme{}, err
	}

	isVariantVO := NewIsVariant(isVariant)

	return TrackTheme{
		trackID:     trackIDVO,
		themeID:     themeIDVO,
		startSecond: startSecondVO,
		endSecond:   endSecondVO,
		isVariant:   isVariantVO,
	}, nil
}

func (tt TrackTheme) TrackID() TrackID {
	return tt.trackID
}

func (tt TrackTheme) ThemeID() ThemeID {
	return tt.themeID
}

func (tt TrackTheme) StartSecond() StartSecond {
	return tt.startSecond
}

func (tt TrackTheme) EndSecond() EndSecond {
	return tt.endSecond
}

func (tt TrackTheme) IsVariant() isVariant {
	return tt.isVariant
}

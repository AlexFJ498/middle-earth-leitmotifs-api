package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidThemeID = fmt.Errorf("invalid theme ID")
var ErrInvalidThemeName = fmt.Errorf("invalid theme name")
var ErrThemeNotFound = fmt.Errorf("theme not found")

type ThemeID struct {
	value string
}

type ThemeName struct {
	value string
}

func NewThemeID() (ThemeID, error) {
	v, err := uuid.NewRandom()
	if err != nil {
		return ThemeID{}, ErrInvalidThemeID
	}
	return ThemeID{value: v.String()}, nil
}

func NewThemeIDFromString(id string) (ThemeID, error) {
	if id == "" {
		return ThemeID{}, ErrInvalidThemeID
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return ThemeID{}, ErrInvalidThemeID
	}
	return ThemeID{value: id}, nil
}

func (id ThemeID) String() string {
	return id.value
}

func NewThemeName(value string) (ThemeName, error) {
	if value == "" {
		return ThemeName{}, ErrInvalidThemeName
	}
	return ThemeName{value: value}, nil
}

func (name ThemeName) String() string {
	return name.value
}

type ThemeRepository interface {
	Save(ctx context.Context, theme Theme) error
	Find(ctx context.Context, id ThemeID) (Theme, error)
	FindAll(ctx context.Context) ([]Theme, error)
	Delete(ctx context.Context, id ThemeID) error
	Update(ctx context.Context, theme Theme) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=ThemeRepository

type Theme struct {
	id         ThemeID
	name       ThemeName
	firstHeard TrackID
	groupID    GroupID
	categoryID *CategoryID // Optional
}

func NewTheme(name, firstHeard, groupID string, categoryID *string) (Theme, error) {
	idVO, err := NewThemeID()
	if err != nil {
		return Theme{}, err
	}

	nameVO, err := NewThemeName(name)
	if err != nil {
		return Theme{}, err
	}

	firstHeardVO, err := NewTrackIDFromString(firstHeard)
	if err != nil {
		return Theme{}, err
	}

	groupIDVO, err := NewGroupIDFromString(groupID)
	if err != nil {
		return Theme{}, err
	}

	var categoryIDVO *CategoryID
	if categoryID != nil {
		var err error
		categoryValue, err := NewCategoryIDFromString(*categoryID)
		if err != nil {
			return Theme{}, err
		}
		categoryIDVO = &categoryValue
	}

	return Theme{
		id:         idVO,
		name:       nameVO,
		firstHeard: firstHeardVO,
		groupID:    groupIDVO,
		categoryID: categoryIDVO,
	}, nil
}

func NewThemeWithID(id, name, firstHeard, groupID string, categoryID *string) (Theme, error) {
	idVO, err := NewThemeIDFromString(id)
	if err != nil {
		return Theme{}, err
	}

	nameVO, err := NewThemeName(name)
	if err != nil {
		return Theme{}, err
	}

	firstHeardVO, err := NewTrackIDFromString(firstHeard)
	if err != nil {
		return Theme{}, err
	}

	groupIDVO, err := NewGroupIDFromString(groupID)
	if err != nil {
		return Theme{}, err
	}

	var categoryIDVO *CategoryID
	if categoryID != nil {
		var err error
		categoryValue, err := NewCategoryIDFromString(*categoryID)
		if err != nil {
			return Theme{}, err
		}
		categoryIDVO = &categoryValue
	}

	return Theme{
		id:         idVO,
		name:       nameVO,
		firstHeard: firstHeardVO,
		groupID:    groupIDVO,
		categoryID: categoryIDVO,
	}, nil
}

func (t Theme) ID() ThemeID {
	return t.id
}

func (t Theme) Name() ThemeName {
	return t.name
}

func (t Theme) FirstHeard() TrackID {
	return t.firstHeard
}

func (t Theme) GroupID() GroupID {
	return t.groupID
}

func (t Theme) CategoryID() *CategoryID {
	return t.categoryID
}

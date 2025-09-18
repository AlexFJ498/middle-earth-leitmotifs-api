package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidGroupID = fmt.Errorf("invalid group ID")
var ErrInvalidGroupName = fmt.Errorf("invalid group name")
var ErrInvalidGroupDescription = fmt.Errorf("invalid description")
var ErrInvalidImageURL = fmt.Errorf("invalid image URL")
var ErrGroupNotFound = fmt.Errorf("group not found")

type GroupID struct {
	value string
}

type GroupName struct {
	value string
}

type GroupDescription struct {
	value string
}

type ImageURL struct {
	value string
}

func NewGroupID() (GroupID, error) {
	v, err := uuid.NewRandom()
	if err != nil {
		return GroupID{}, fmt.Errorf("%w: %w", ErrInvalidGroupID, err)
	}
	return GroupID{value: v.String()}, nil
}

func NewGroupIDFromString(id string) (GroupID, error) {
	if id == "" {
		return GroupID{}, ErrInvalidGroupID
	}

	v, err := uuid.Parse(id)
	if err != nil {
		return GroupID{}, ErrInvalidGroupID
	}

	return GroupID{value: v.String()}, nil
}

func (id GroupID) String() string {
	return id.value
}

func NewGroupName(value string) (GroupName, error) {
	if value == "" {
		return GroupName{}, ErrInvalidGroupName
	}
	return GroupName{value: value}, nil
}

func (n GroupName) String() string {
	return n.value
}

func NewGroupDescription(value string) (GroupDescription, error) {
	if value == "" {
		return GroupDescription{}, ErrInvalidGroupDescription
	}
	return GroupDescription{value: value}, nil
}

func (d GroupDescription) String() string {
	return d.value
}

func NewImageURL(value string) (ImageURL, error) {
	if value == "" {
		return ImageURL{}, ErrInvalidImageURL
	}
	return ImageURL{value: value}, nil
}

func (u ImageURL) String() string {
	return u.value
}

type GroupRepository interface {
	Save(ctx context.Context, group Group) error
	Find(ctx context.Context, id GroupID) (Group, error)
	FindAll(ctx context.Context) ([]Group, error)
	Delete(ctx context.Context, id GroupID) error
	Update(ctx context.Context, group Group) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=GroupRepository

type Group struct {
	id          GroupID
	name        GroupName
	description GroupDescription
	imageURL    ImageURL
}

func NewGroup(name, description, imageURL string) (Group, error) {
	idVO, err := NewGroupID()
	if err != nil {
		return Group{}, err
	}

	nameVO, err := NewGroupName(name)
	if err != nil {
		return Group{}, err
	}

	descriptionVO, err := NewGroupDescription(description)
	if err != nil {
		return Group{}, err
	}

	imageURLVO, err := NewImageURL(imageURL)
	if err != nil {
		return Group{}, err
	}

	group := Group{
		id:          idVO,
		name:        nameVO,
		description: descriptionVO,
		imageURL:    imageURLVO,
	}

	return group, nil
}

func NewGroupWithID(id, name, description, imageURL string) (Group, error) {
	idVO, err := NewGroupIDFromString(id)
	if err != nil {
		return Group{}, err
	}

	nameVO, err := NewGroupName(name)
	if err != nil {
		return Group{}, err
	}

	descriptionVO, err := NewGroupDescription(description)
	if err != nil {
		return Group{}, err
	}

	imageURLVO, err := NewImageURL(imageURL)
	if err != nil {
		return Group{}, err
	}

	group := Group{
		id:          idVO,
		name:        nameVO,
		description: descriptionVO,
		imageURL:    imageURLVO,
	}

	return group, nil
}

func (g Group) ID() GroupID {
	return g.id
}

func (g Group) Name() GroupName {
	return g.name
}

func (g Group) Description() GroupDescription {
	return g.description
}

func (g Group) ImageURL() ImageURL {
	return g.imageURL
}

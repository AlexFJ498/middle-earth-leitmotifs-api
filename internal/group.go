package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidGroupID = fmt.Errorf("invalid group ID")
var ErrInvalidGroupName = fmt.Errorf("invalid group name")
var ErrGroupNotFound = fmt.Errorf("group not found")

type GroupID struct {
	value string
}

type GroupName struct {
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

type GroupRepository interface {
	Save(ctx context.Context, group Group) error
	Find(ctx context.Context, id GroupID) (Group, error)
	FindAll(ctx context.Context) ([]Group, error)
	Delete(ctx context.Context, id GroupID) error
	Update(ctx context.Context, group Group) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=GroupRepository

type Group struct {
	id   GroupID
	name GroupName
}

func NewGroup(name string) (Group, error) {
	idVO, err := NewGroupID()
	if err != nil {
		return Group{}, err
	}

	nameVO, err := NewGroupName(name)
	if err != nil {
		return Group{}, err
	}

	group := Group{
		id:   idVO,
		name: nameVO,
	}

	return group, nil
}

func NewGroupWithID(id, name string) (Group, error) {
	idVO, err := NewGroupIDFromString(id)
	if err != nil {
		return Group{}, err
	}

	nameVO, err := NewGroupName(name)
	if err != nil {
		return Group{}, err
	}

	group := Group{
		id:   idVO,
		name: nameVO,
	}

	return group, nil
}

func (g Group) ID() GroupID {
	return g.id
}

func (g Group) Name() GroupName {
	return g.name
}

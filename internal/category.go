package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidCategoryID = fmt.Errorf("invalid category ID")
var ErrInvalidCategoryName = fmt.Errorf("invalid category name")
var ErrCategoryNotFound = fmt.Errorf("category not found")

type CategoryID struct {
	value string
}

type CategoryName struct {
	value string
}

func NewCategoryID() (CategoryID, error) {
	v, err := uuid.NewRandom()
	if err != nil {
		return CategoryID{}, fmt.Errorf("%w: %w", ErrInvalidCategoryID, err)
	}

	return CategoryID{
		value: v.String(),
	}, nil
}

func NewCategoryIDFromString(id string) (CategoryID, error) {
	if id == "" {
		return CategoryID{}, ErrInvalidCategoryID
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return CategoryID{}, ErrInvalidCategoryID
	}

	return CategoryID{
		value: id,
	}, nil
}

func (id CategoryID) String() string {
	return id.value
}

func NewCategoryName(value string) (CategoryName, error) {
	if value == "" {
		return CategoryName{}, ErrInvalidCategoryName
	}

	return CategoryName{
		value: value,
	}, nil
}

func (n CategoryName) String() string {
	return n.value
}

type CategoryRepository interface {
	Save(ctx context.Context, category Category) error
	Find(ctx context.Context, id CategoryID) (Category, error)
	FindAll(ctx context.Context) ([]Category, error)
	Delete(ctx context.Context, id CategoryID) error
	Update(ctx context.Context, category Category) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CategoryRepository

type Category struct {
	id   CategoryID
	name CategoryName
}

func NewCategory(name string) (Category, error) {
	idVO, err := NewCategoryID()
	if err != nil {
		return Category{}, err
	}

	nameVO, err := NewCategoryName(name)
	if err != nil {
		return Category{}, err
	}

	category := Category{
		id:   idVO,
		name: nameVO,
	}

	return category, nil
}

func NewCategoryWithID(id, name string) (Category, error) {
	idVO, err := NewCategoryIDFromString(id)
	if err != nil {
		return Category{}, err
	}

	nameVO, err := NewCategoryName(name)
	if err != nil {
		return Category{}, err
	}

	category := Category{
		id:   idVO,
		name: nameVO,
	}

	return category, nil
}

func (c Category) ID() CategoryID {
	return c.id
}

func (c Category) Name() CategoryName {
	return c.name
}

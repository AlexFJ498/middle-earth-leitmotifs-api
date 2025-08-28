package updating

import (
	"context"
	"errors"
	"testing"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	testID       = "123e4567-e89b-12d3-a456-426614174000"
	movieName    = "The Fellowship of the Ring"
	groupName    = "The Elves"
	categoryName = "The Mordor Accompaniments"

	domainMovieType    = "domain.Movie"
	domainGroupType    = "domain.Group"
	domainCategoryType = "domain.Category"

	repositoryErrorMsg = "repository error"
	invalidId          = "invalid-id"
)

func TestMovieServiceUpdateMovieRepositoryError(t *testing.T) {
	dto := dto.MovieUpdateRequest{
		Name: movieName,
	}

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainMovieType)).Return(errors.New(repositoryErrorMsg)).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	service := NewMovieService(movieRepositoryMock)

	err := service.UpdateMovie(context.Background(), testID, dto)
	assert.Error(t, err)
}

func TestMovieServiceUpdateMovieSuccess(t *testing.T) {
	dto := dto.MovieUpdateRequest{
		Name: movieName,
	}

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainMovieType)).Return(nil).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	service := NewMovieService(movieRepositoryMock)

	err := service.UpdateMovie(context.Background(), testID, dto)
	assert.NoError(t, err)
}

func TestMovieServiceUpdateMovieInvalidID(t *testing.T) {
	dto := dto.MovieUpdateRequest{
		Name: movieName,
	}

	movieRepositoryMock := new(storagemocks.MovieRepository)
	defer movieRepositoryMock.AssertExpectations(t)

	service := NewMovieService(movieRepositoryMock)

	err := service.UpdateMovie(context.Background(), invalidId, dto)
	assert.Error(t, err)
}

func TestGroupServiceUpdateGroupRepositoryError(t *testing.T) {
	dto := dto.GroupUpdateRequest{
		Name: groupName,
	}

	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainGroupType)).Return(errors.New(repositoryErrorMsg)).Once()
	defer groupRepositoryMock.AssertExpectations(t)

	service := NewGroupService(groupRepositoryMock)

	err := service.UpdateGroup(context.Background(), testID, dto)
	assert.Error(t, err)
}

func TestGroupServiceUpdateGroupSuccess(t *testing.T) {
	dto := dto.GroupUpdateRequest{
		Name: groupName,
	}

	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainGroupType)).Return(nil).Once()
	defer groupRepositoryMock.AssertExpectations(t)

	service := NewGroupService(groupRepositoryMock)

	err := service.UpdateGroup(context.Background(), testID, dto)
	assert.NoError(t, err)
}

func TestGroupServiceUpdateGroupInvalidID(t *testing.T) {
	dto := dto.GroupUpdateRequest{
		Name: groupName,
	}

	groupRepositoryMock := new(storagemocks.GroupRepository)
	defer groupRepositoryMock.AssertExpectations(t)

	service := NewGroupService(groupRepositoryMock)

	err := service.UpdateGroup(context.Background(), invalidId, dto)
	assert.Error(t, err)
}

func TestCategoryServiceUpdateCategoryRepositoryError(t *testing.T) {
	dto := dto.CategoryUpdateRequest{
		Name: categoryName,
	}

	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	categoryRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainCategoryType)).Return(errors.New(repositoryErrorMsg)).Once()
	defer categoryRepositoryMock.AssertExpectations(t)

	service := NewCategoryService(categoryRepositoryMock)

	err := service.UpdateCategory(context.Background(), testID, dto)
	assert.Error(t, err)
}

func TestCategoryServiceUpdateCategorySuccess(t *testing.T) {
	dto := dto.CategoryUpdateRequest{
		Name: categoryName,
	}

	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	categoryRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainCategoryType)).Return(nil).Once()
	defer categoryRepositoryMock.AssertExpectations(t)

	service := NewCategoryService(categoryRepositoryMock)

	err := service.UpdateCategory(context.Background(), testID, dto)
	assert.NoError(t, err)
}

func TestCategoryServiceUpdateCategoryInvalidID(t *testing.T) {
	dto := dto.CategoryUpdateRequest{
		Name: categoryName,
	}

	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	defer categoryRepositoryMock.AssertExpectations(t)

	service := NewCategoryService(categoryRepositoryMock)

	err := service.UpdateCategory(context.Background(), invalidId, dto)
	assert.Error(t, err)
}

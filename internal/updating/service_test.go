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
	testID           = "123e4567-e89b-12d3-a456-426614174000"
	movieName        = "The Fellowship of the Ring"
	updatedMovieName = "The Fellowship of the Ring - Updated"
	groupName        = "Fellowship of the Ring"
	updatedGroupName = "Fellowship of the Ring - Updated"
	domainMovieType  = "domain.Movie"
	domainGroupType  = "domain.Group"
)

func TestMovieServiceUpdateMovieRepositoryError(t *testing.T) {
	dto := dto.MovieUpdateRequest{
		Name: updatedMovieName,
	}

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainMovieType)).Return(errors.New("repository error")).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	service := NewMovieService(movieRepositoryMock)

	err := service.UpdateMovie(context.Background(), testID, dto)
	assert.Error(t, err)
}

func TestMovieServiceUpdateMovieSuccess(t *testing.T) {
	dto := dto.MovieUpdateRequest{
		Name: updatedMovieName,
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
		Name: updatedMovieName,
	}

	movieRepositoryMock := new(storagemocks.MovieRepository)
	defer movieRepositoryMock.AssertExpectations(t)

	service := NewMovieService(movieRepositoryMock)

	err := service.UpdateMovie(context.Background(), "invalid-id", dto)
	assert.Error(t, err)
}

func TestGroupServiceUpdateGroupRepositoryError(t *testing.T) {
	dto := dto.GroupUpdateRequest{
		Name: updatedGroupName,
	}

	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainGroupType)).Return(errors.New("repository error")).Once()
	defer groupRepositoryMock.AssertExpectations(t)

	service := NewGroupService(groupRepositoryMock)

	err := service.UpdateGroup(context.Background(), testID, dto)
	assert.Error(t, err)
}

func TestGroupServiceUpdateGroupSuccess(t *testing.T) {
	dto := dto.GroupUpdateRequest{
		Name: updatedGroupName,
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
		Name: updatedGroupName,
	}

	groupRepositoryMock := new(storagemocks.GroupRepository)
	defer groupRepositoryMock.AssertExpectations(t)

	service := NewGroupService(groupRepositoryMock)

	err := service.UpdateGroup(context.Background(), "invalid-id", dto)
	assert.Error(t, err)
}

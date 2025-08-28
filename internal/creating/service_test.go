package creating

import (
	"context"
	"errors"
	"testing"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/storage/storagemocks"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/event/eventmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	userName       = "Test User"
	userEmail      = "testuser@example.com"
	userPassword   = "password123"
	domainUserType = "domain.User"

	repositoryErrorMsg = "repository error"
)

func TestUserServiceCreateUserRepositoryError(t *testing.T) {
	dto := dto.UserCreateRequest{
		Name:     userName,
		Email:    userEmail,
		Password: userPassword,
	}

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType(domainUserType)).Return(errors.New(repositoryErrorMsg)).Once()
	defer userRepositoryMock.AssertExpectations(t)

	eventBusMock := new(eventmocks.Bus)
	defer eventBusMock.AssertExpectations(t)

	service := NewUserService(userRepositoryMock, eventBusMock)

	err := service.CreateUser(context.Background(), dto)
	assert.Error(t, err)
}

func TestUserServiceCreateUserSuccess(t *testing.T) {
	dto := dto.UserCreateRequest{
		Name:     userName,
		Email:    userEmail,
		Password: userPassword,
	}

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType(domainUserType)).Return(nil).Once()
	defer userRepositoryMock.AssertExpectations(t)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil).Once()
	defer eventBusMock.AssertExpectations(t)

	service := NewUserService(userRepositoryMock, eventBusMock)

	err := service.CreateUser(context.Background(), dto)
	assert.NoError(t, err)
}

func TestUserServiceCreateUserEventBusError(t *testing.T) {
	dto := dto.UserCreateRequest{
		Name:     userName,
		Email:    userEmail,
		Password: userPassword,
	}

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType(domainUserType)).Return(nil).Once()
	defer userRepositoryMock.AssertExpectations(t)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(errors.New("event bus error")).Once()
	defer eventBusMock.AssertExpectations(t)

	service := NewUserService(userRepositoryMock, eventBusMock)

	err := service.CreateUser(context.Background(), dto)
	assert.Error(t, err)
}

func TestMovieServiceCreateMovieRepositoryError(t *testing.T) {
	dto := dto.MovieCreateRequest{
		Name: "Test Movie",
	}

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Movie")).Return(errors.New(repositoryErrorMsg)).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	service := NewMovieService(movieRepositoryMock)

	err := service.CreateMovie(context.Background(), dto)
	assert.Error(t, err)
}

func TestMovieServiceCreateMovieSuccess(t *testing.T) {
	dto := dto.MovieCreateRequest{
		Name: "Test Movie",
	}

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Movie")).Return(nil).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	service := NewMovieService(movieRepositoryMock)

	err := service.CreateMovie(context.Background(), dto)
	assert.NoError(t, err)
}

func TestGroupServiceCreateGroupRepositoryError(t *testing.T) {
	dto := dto.GroupCreateRequest{
		Name: "Test Group",
	}

	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Group")).Return(errors.New(repositoryErrorMsg)).Once()
	defer groupRepositoryMock.AssertExpectations(t)

	service := NewGroupService(groupRepositoryMock)

	err := service.CreateGroup(context.Background(), dto)
	assert.Error(t, err)
}

func TestGroupServiceCreateGroupSuccess(t *testing.T) {
	dto := dto.GroupCreateRequest{
		Name: "Test Group",
	}

	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Group")).Return(nil).Once()
	defer groupRepositoryMock.AssertExpectations(t)

	service := NewGroupService(groupRepositoryMock)

	err := service.CreateGroup(context.Background(), dto)
	assert.NoError(t, err)
}

func TestCategoryServiceCreateCategoryRepositoryError(t *testing.T) {
	dto := dto.CategoryCreateRequest{
		Name: "Test Category",
	}

	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	categoryRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Category")).Return(errors.New(repositoryErrorMsg)).Once()
	defer categoryRepositoryMock.AssertExpectations(t)

	service := NewCategoryService(categoryRepositoryMock)

	err := service.CreateCategory(context.Background(), dto)
	assert.Error(t, err)
}

func TestCategoryServiceCreateCategorySuccess(t *testing.T) {
	dto := dto.CategoryCreateRequest{
		Name: "Test Category",
	}

	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	categoryRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Category")).Return(nil).Once()
	defer categoryRepositoryMock.AssertExpectations(t)

	service := NewCategoryService(categoryRepositoryMock)

	err := service.CreateCategory(context.Background(), dto)
	assert.NoError(t, err)
}

func TestTrackServiceCreateTrackRepositoryError(t *testing.T) {
	dto := dto.TrackCreateRequest{
		Name:    "Test Track",
		MovieID: "456e7890-e89b-12d3-a456-426614174111",
	}

	trackRepositoryMock := new(storagemocks.TrackRepository)
	trackRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Track")).Return(errors.New(repositoryErrorMsg)).Once()
	defer trackRepositoryMock.AssertExpectations(t)

	service := NewTrackService(trackRepositoryMock)

	err := service.CreateTrack(context.Background(), dto)
	assert.Error(t, err)
}

func TestTrackServiceCreateTrackSuccess(t *testing.T) {
	dto := dto.TrackCreateRequest{
		Name:    "Test Track",
		MovieID: "456e7890-e89b-12d3-a456-426614174111",
	}

	trackRepositoryMock := new(storagemocks.TrackRepository)
	trackRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Track")).Return(nil).Once()
	defer trackRepositoryMock.AssertExpectations(t)

	service := NewTrackService(trackRepositoryMock)

	err := service.CreateTrack(context.Background(), dto)
	assert.NoError(t, err)
}

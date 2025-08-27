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

	movieName       = "Test Movie"
	domainMovieType = "domain.Movie"

	groupName       = "Test Group"
	domainGroupType = "domain.Group"

	eventEventSliceType = "[]event.Event"
	repositoryErrorMsg  = "repository error"
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
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType(eventEventSliceType)).Return(nil).Once()
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
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType(eventEventSliceType)).Return(errors.New("event bus error")).Once()
	defer eventBusMock.AssertExpectations(t)

	service := NewUserService(userRepositoryMock, eventBusMock)

	err := service.CreateUser(context.Background(), dto)
	assert.Error(t, err)
}

func TestMovieServiceCreateMovieRepositoryError(t *testing.T) {
	dto := dto.MovieCreateRequest{
		Name: movieName,
	}

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType(domainMovieType)).Return(errors.New(repositoryErrorMsg)).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	service := NewMovieService(movieRepositoryMock)

	err := service.CreateMovie(context.Background(), dto)
	assert.Error(t, err)
}

func TestMovieServiceCreateMovieSuccess(t *testing.T) {
	dto := dto.MovieCreateRequest{
		Name: movieName,
	}

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType(domainMovieType)).Return(nil).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	service := NewMovieService(movieRepositoryMock)

	err := service.CreateMovie(context.Background(), dto)
	assert.NoError(t, err)
}

func TestGroupServiceCreateGroupRepositoryError(t *testing.T) {
	dto := dto.GroupCreateRequest{
		Name: groupName,
	}

	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType(domainGroupType)).Return(errors.New(repositoryErrorMsg)).Once()
	defer groupRepositoryMock.AssertExpectations(t)

	service := NewGroupService(groupRepositoryMock)

	err := service.CreateGroup(context.Background(), dto)
	assert.Error(t, err)
}

func TestGroupServiceCreateGroupSuccess(t *testing.T) {
	dto := dto.GroupCreateRequest{
		Name: groupName,
	}

	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType(domainGroupType)).Return(nil).Once()
	defer groupRepositoryMock.AssertExpectations(t)

	service := NewGroupService(groupRepositoryMock)

	err := service.CreateGroup(context.Background(), dto)
	assert.NoError(t, err)
}

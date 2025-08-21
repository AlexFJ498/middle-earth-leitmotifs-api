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

const userID = "123e4567-e89b-12d3-a456-426614174000"
const userName = "Test User"
const userEmail = "testuser@example.com"
const userPassword = "password123"

const domainUserType = "domain.User"
const eventEventSliceType = "[]event.Event"

func TestUserServiceCreateUserRepositoryError(t *testing.T) {
	dto := dto.UserCreateRequest{
		Name:     userName,
		Email:    userEmail,
		Password: userPassword,
	}

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType(domainUserType)).Return(errors.New("repository error")).Once()
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

package creating

import (
	"context"
	"errors"
	"testing"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/storage/storagemocks"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/event/eventmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const userID = "123e4567-e89b-12d3-a456-426614174000"
const userName = "Test User"
const userEmail = "testuser@example.com"

const domainUserType = "domain.User"

func TestUserServiceCreateUserRepositoryError(t *testing.T) {
	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType(domainUserType)).Return(errors.New("repository error")).Once()
	defer userRepositoryMock.AssertExpectations(t)

	eventBusMock := new(eventmocks.Bus)
	defer eventBusMock.AssertExpectations(t)

	service := NewUserService(userRepositoryMock, eventBusMock)

	err := service.CreateUser(context.Background(), userID, userName, userEmail)
	assert.Error(t, err)
}

func TestUserServiceCreateUserSuccess(t *testing.T) {
	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType(domainUserType)).Return(nil).Once()
	defer userRepositoryMock.AssertExpectations(t)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil).Once()
	defer eventBusMock.AssertExpectations(t)

	service := NewUserService(userRepositoryMock, eventBusMock)

	err := service.CreateUser(context.Background(), userID, userName, userEmail)
	assert.NoError(t, err)
}

func TestUserServiceCreateUserEventBusError(t *testing.T) {
	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType(domainUserType)).Return(nil).Once()
	defer userRepositoryMock.AssertExpectations(t)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(errors.New("event bus error")).Once()
	defer eventBusMock.AssertExpectations(t)

	service := NewUserService(userRepositoryMock, eventBusMock)

	err := service.CreateUser(context.Background(), userID, userName, userEmail)
	assert.Error(t, err)
}

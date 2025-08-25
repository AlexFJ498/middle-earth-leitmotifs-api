package listing

import (
	"context"
	"errors"
	"testing"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserServiceListUsersRepositoryError(t *testing.T) {
	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("FindAll", mock.Anything).Return(nil, errors.New("repository error")).Once()

	userService := NewUserService(userRepositoryMock)

	ctx := context.Background()
	_, err := userService.ListUsers(ctx)
	assert.Error(t, err)
}

func TestUserServiceListUsersSuccess(t *testing.T) {
	userRepositoryMock := new(storagemocks.UserRepository)
	users := []domain.User{}
	user1, err := domain.NewUser("John Doe", "john@example.com", "password123", false)
	assert.NoError(t, err)
	users = append(users, user1)
	user2, err := domain.NewUser("Jane Doe", "jane@example.com", "password456", false)
	assert.NoError(t, err)
	users = append(users, user2)
	userRepositoryMock.On("FindAll", mock.Anything).Return(users, nil).Once()

	userService := NewUserService(userRepositoryMock)

	usersDTO, err := userService.ListUsers(context.Background())
	assert.NoError(t, err)
	assert.Len(t, usersDTO, 2)
	assert.Equal(t, "John Doe", usersDTO[0].Name)
	assert.Equal(t, "john@example.com", usersDTO[0].Email)
	assert.Equal(t, "Jane Doe", usersDTO[1].Name)
	assert.Equal(t, "jane@example.com", usersDTO[1].Email)
}

package authenticating

import (
	"context"
	"testing"
	"time"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/auth"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
)

const (
	email  = "user@example.com"
	jwtKey = "some.jwt.token"
	exp    = 24 * time.Hour
)

func TestLoginServiceLoginUserRepositoryEmailError(t *testing.T) {
	email := "invalid-email"
	password := "password123"

	userRepositoryMock := new(storagemocks.UserRepository)
	service := NewLoginService(userRepositoryMock, []byte(jwtKey), exp)

	_, err := service.LoginUser(context.Background(), email, password)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrInvalidUserEmail, err)
}

func TestLoginServiceLoginUserRepositoryFindByEmailError(t *testing.T) {
	password := "password123"

	emailVO, err := domain.NewUserEmail(email)
	assert.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("FindByEmail", context.Background(), emailVO).Return(domain.User{}, domain.ErrUserNotFound)
	defer userRepositoryMock.AssertExpectations(t)

	service := NewLoginService(userRepositoryMock, []byte(jwtKey), exp)

	_, err = service.LoginUser(context.Background(), email, password)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrUserNotFound, err)
}

func TestLoginServiceLoginUserPasswordError(t *testing.T) {
	password := "wrongpassword"

	emailVO, err := domain.NewUserEmail(email)
	assert.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("FindByEmail", context.Background(), emailVO).Return(domain.User{}, nil)
	defer userRepositoryMock.AssertExpectations(t)

	service := NewLoginService(userRepositoryMock, []byte(jwtKey), exp)

	_, err = service.LoginUser(context.Background(), email, password)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrInvalidUserPassword, err)
}

func TestLoginServiceLoginUserSuccess(t *testing.T) {
	password := "password123"
	hashedPassword, _ := auth.HashPassword(password)
	user, _ := domain.NewUser("name", email, hashedPassword)

	emailVO, err := domain.NewUserEmail(email)
	assert.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("FindByEmail", context.Background(), emailVO).Return(user, nil)
	defer userRepositoryMock.AssertExpectations(t)

	service := NewLoginService(userRepositoryMock, []byte(jwtKey), exp)

	_, err = service.LoginUser(context.Background(), email, password)
	assert.NoError(t, err)
}

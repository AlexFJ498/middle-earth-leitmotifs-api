package sqldb

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const userID = "123e4567-e89b-12d3-a456-426614174000"
const name = "Test User"
const userEmail = "test@example.com"

func TestUserRepositorySaveRepositoryError(t *testing.T) {
	user, err := domain.NewUser(userID, name, userEmail)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO users (id, name, email) VALUES ($1, $2, $3)").
		WithArgs(userID, name, userEmail).
		WillReturnError(errors.New("database error"))

	repo := NewUserRepository(db, 1*time.Second)

	err = repo.Save(context.Background(), user)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestUserRepositorySaveSuccess(t *testing.T) {
	user, err := domain.NewUser(userID, name, userEmail)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO users (id, name, email) VALUES ($1, $2, $3)").
		WithArgs(userID, name, userEmail).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewUserRepository(db, 1*time.Second)

	err = repo.Save(context.Background(), user)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func TestUserRepositoryFindUserNotFound(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		"SELECT users.id, users.name, users.email FROM users WHERE id = ?").
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}))

	repo := NewUserRepository(db, 1*time.Second)

	userIDObj, err := domain.NewUserID(userID)
	require.NoError(t, err)

	foundUser, err := repo.Find(context.Background(), userIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.Equal(t, domain.User{}, foundUser)
}

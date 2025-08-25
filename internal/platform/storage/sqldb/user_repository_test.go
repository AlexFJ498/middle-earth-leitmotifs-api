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
const userPassword = "password123"

const querySelectUserByEmail = "SELECT users.id, users.name, users.email, users.password, users.is_admin FROM users WHERE email = $1"
const querySelectAllUsers = "SELECT users.id, users.name, users.email, users.password, users.is_admin FROM users"

func TestUserRepositorySaveRepositoryError(t *testing.T) {
	user, err := domain.NewUserWithID(userID, name, userEmail, userPassword, false)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectUserByEmail).
		WithArgs(userEmail).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "is_admin"}))

	sqlMock.ExpectExec(
		"INSERT INTO users (id, name, email, password, is_admin) VALUES ($1, $2, $3, $4, $5)").
		WithArgs(userID, name, userEmail, userPassword, false).
		WillReturnError(errors.New("database error"))

	repo := NewUserRepository(db, 1*time.Second)

	err = repo.Save(context.Background(), user)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestUserRepositorySaveSuccess(t *testing.T) {
	user, err := domain.NewUserWithID(userID, name, userEmail, userPassword, false)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectUserByEmail).
		WithArgs(userEmail).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "is_admin"}))

	sqlMock.ExpectExec(
		"INSERT INTO users (id, name, email, password, is_admin) VALUES ($1, $2, $3, $4, $5)").
		WithArgs(userID, name, userEmail, userPassword, false).
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
		"SELECT users.id, users.name, users.email, users.password, users.is_admin FROM users WHERE id = $1").
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "is_admin"}))

	repo := NewUserRepository(db, 1*time.Second)

	userIDObj, err := domain.NewUserIDFromString(userID)
	require.NoError(t, err)

	foundUser, err := repo.Find(context.Background(), userIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.Equal(t, domain.User{}, foundUser)
}

func TestUserRepositorySaveUserExistsError(t *testing.T) {
	user, err := domain.NewUserWithID(userID, name, userEmail, userPassword, false)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectUserByEmail).
		WithArgs(userEmail).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "is_admin"}).
			AddRow(userID, name, userEmail, userPassword, false))

	repo := NewUserRepository(db, 1*time.Second)

	err = repo.Save(context.Background(), user)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)

	assert.Equal(t, domain.ErrUserAlreadyExists, err)
}

func TestUserRepositoryFindByEmailNotFound(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectUserByEmail).
		WithArgs(userEmail).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "is_admin"}))

	repo := NewUserRepository(db, 1*time.Second)

	emailObj, err := domain.NewUserEmail(userEmail)
	require.NoError(t, err)

	foundUser, err := repo.FindByEmail(context.Background(), emailObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.Equal(t, domain.User{}, foundUser)
}

func TestUserRepositoryFindByEmailSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectUserByEmail).
		WithArgs(userEmail).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "is_admin"}).
			AddRow(userID, name, userEmail, userPassword, false))

	repo := NewUserRepository(db, 1*time.Second)

	emailObj, err := domain.NewUserEmail(userEmail)
	require.NoError(t, err)

	_, err = repo.FindByEmail(context.Background(), emailObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func TestUserRepositoryFindAllSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		querySelectAllUsers).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "is_admin"}).
			AddRow(userID, name, userEmail, userPassword, false).
			AddRow("223e4567-e89b-12d3-a456-426614174001", "Another User", "another.user@example.com", "password123", false))

	repo := NewUserRepository(db, 1*time.Second)

	users, err := repo.FindAll(context.Background())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestUserRepositoryFindAllEmpty(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		querySelectAllUsers).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "is_admin"}))

	repo := NewUserRepository(db, 1*time.Second)

	users, err := repo.FindAll(context.Background())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Len(t, users, 0)
}

func TestUserRepositoryFindAllQueryError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		querySelectAllUsers).
		WillReturnError(errors.New("query error"))

	repo := NewUserRepository(db, 1*time.Second)

	users, err := repo.FindAll(context.Background())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.Len(t, users, 0)
}

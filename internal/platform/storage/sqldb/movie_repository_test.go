package sqldb

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const movieID = "123e4567-e89b-12d3-a456-426614174000"
const movieName = "The Lord of the Rings"
const querySelectAllMovies = "SELECT movies.id, movies.name FROM movies"

func TestMovieRepositorySaveRepositoryError(t *testing.T) {
	movie, err := domain.NewMovieWithID(movieID, movieName)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO movies (id, name) VALUES ($1, $2)").
		WithArgs(movieID, movieName).
		WillReturnError(errors.New("database error"))

	repo := NewMovieRepository(db, 1*time.Second)

	err = repo.Save(context.Background(), movie)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestMovieRepositorySaveSuccess(t *testing.T) {
	movie, err := domain.NewMovieWithID(movieID, movieName)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO movies (id, name) VALUES ($1, $2)").
		WithArgs(movieID, movieName).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewMovieRepository(db, 1*time.Second)

	err = repo.Save(context.Background(), movie)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func TestMovieRepositoryFindNotFound(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery("SELECT movies.id, movies.name FROM movies WHERE id = $1").
		WithArgs(movieID).
		WillReturnError(sql.ErrNoRows)

	repo := NewMovieRepository(db, 1*time.Second)

	movieIDObj, err := domain.NewMovieIDFromString(movieID)
	require.NoError(t, err)

	_, err = repo.Find(context.Background(), movieIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestMovieRepositoryFindSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery("SELECT movies.id, movies.name FROM movies WHERE id = $1").
		WithArgs(movieID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(movieID, movieName))

	repo := NewMovieRepository(db, 1*time.Second)

	movieIDObj, err := domain.NewMovieIDFromString(movieID)
	require.NoError(t, err)

	movie, err := repo.Find(context.Background(), movieIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Equal(t, movieIDObj, movie.ID())
	assert.Equal(t, movieName, movie.Name().String())
}

func TestMovieRepositoryFindAllSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectAllMovies).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(movieID, movieName).
			AddRow("223e4567-e89b-12d3-a456-426614174001", "The Hobbit"))

	repo := NewMovieRepository(db, 1*time.Second)

	movies, err := repo.FindAll(context.Background())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Len(t, movies, 2)
	assert.Equal(t, movieID, movies[0].ID().String())
	assert.Equal(t, movieName, movies[0].Name().String())
	assert.Equal(t, "223e4567-e89b-12d3-a456-426614174001", movies[1].ID().String())
	assert.Equal(t, "The Hobbit", movies[1].Name().String())
}

func TestMovieRepositoryFindAllEmpty(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectAllMovies).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))

	repo := NewMovieRepository(db, 1*time.Second)

	movies, err := repo.FindAll(context.Background())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Len(t, movies, 0)
}

func TestMovieRepositoryFindAllQueryError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectAllMovies).
		WillReturnError(errors.New("query error"))

	repo := NewMovieRepository(db, 1*time.Second)

	movies, err := repo.FindAll(context.Background())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.Nil(t, movies)
}

func TestMovieRepositoryDeleteError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("DELETE FROM movies WHERE id = $1").
		WithArgs(movieID).
		WillReturnError(errors.New("delete error"))

	repo := NewMovieRepository(db, 1*time.Second)

	movieIDObj, err := domain.NewMovieIDFromString(movieID)
	require.NoError(t, err)

	err = repo.Delete(context.Background(), movieIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestMovieRepositoryDeleteSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("DELETE FROM movies WHERE id = $1").
		WithArgs(movieID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewMovieRepository(db, 1*time.Second)

	movieIDObj, err := domain.NewMovieIDFromString(movieID)
	require.NoError(t, err)

	err = repo.Delete(context.Background(), movieIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func TestMovieRepositoryUpdateError(t *testing.T) {
	movie, err := domain.NewMovieWithID(movieID, movieName)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("UPDATE movies SET id = $1, name = $2 WHERE id = $3").
		WithArgs(movieID, movieName, movieID).
		WillReturnError(errors.New("update error"))

	repo := NewMovieRepository(db, 1*time.Second)

	err = repo.Update(context.Background(), movie)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestMovieRepositoryUpdateSuccess(t *testing.T) {
	movie, err := domain.NewMovieWithID(movieID, movieName)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("UPDATE movies SET id = $1, name = $2 WHERE id = $3").
		WithArgs(movieID, movieName, movieID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewMovieRepository(db, 1*time.Second)

	err = repo.Update(context.Background(), movie)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

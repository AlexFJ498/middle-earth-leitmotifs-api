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

const (
	trackID            = "123e4567-e89b-12d3-a456-426614174000"
	trackName          = "The Shire"
	trackMovieID       = "456e7890-e89b-12d3-a456-426614174111"
	connectionErrorMsg = "connection error"
	selectQuery        = "SELECT tracks.id, tracks.name, tracks.movie_id, tracks.spotify_url FROM tracks WHERE id = $1"
	deleteQuery        = "DELETE FROM tracks WHERE id = $1"
	selectAllQuery     = "SELECT tracks.id, tracks.name, tracks.movie_id, tracks.spotify_url FROM tracks"
)

func TestTrackRepositorySaveError(t *testing.T) {
	track, err := domain.NewTrackWithID(trackID, trackName, trackMovieID, nil)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("INSERT INTO tracks (id, name, movie_id, spotify_url) VALUES ($1, $2, $3, $4)").
		WithArgs(trackID, trackName, trackMovieID, nil).
		WillReturnError(errors.New(connectionErrorMsg))

	repo := NewTrackRepository(db, 1*time.Second)

	err = repo.Save(context.Background(), track)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestTrackRepositorySaveSuccess(t *testing.T) {
	track, err := domain.NewTrackWithID(trackID, trackName, trackMovieID, nil)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("INSERT INTO tracks (id, name, movie_id, spotify_url) VALUES ($1, $2, $3, $4)").
		WithArgs(trackID, trackName, trackMovieID, nil).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewTrackRepository(db, 1*time.Second)

	err = repo.Save(context.Background(), track)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func TestTrackRepositoryFindError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(selectQuery).
		WithArgs(trackID).
		WillReturnError(errors.New(connectionErrorMsg))

	repo := NewTrackRepository(db, 1*time.Second)

	trackID, err := domain.NewTrackIDFromString(trackID)
	require.NoError(t, err)
	_, err = repo.Find(context.Background(), trackID)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestTrackRepositoryFindNotFound(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(selectQuery).
		WithArgs(trackID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "movie_id", "spotify_url"}))

	repo := NewTrackRepository(db, 1*time.Second)

	trackIDVO, err := domain.NewTrackIDFromString(trackID)
	require.NoError(t, err)
	_, err = repo.Find(context.Background(), trackIDVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.ErrorIs(t, err, domain.ErrTrackNotFound)
}

func TestTrackRepositoryFindSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(selectQuery).
		WithArgs(trackID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "movie_id", "spotify_url"}).
			AddRow(trackID, trackName, trackMovieID, nil))

	repo := NewTrackRepository(db, 1*time.Second)

	trackIDVO, err := domain.NewTrackIDFromString(trackID)
	require.NoError(t, err)
	track, err := repo.Find(context.Background(), trackIDVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Equal(t, trackID, track.ID().String())
	assert.Equal(t, trackName, track.Name().String())
	assert.Equal(t, trackMovieID, track.MovieID().String())
}

func TestTrackRepositoryFindAllError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(selectAllQuery).
		WillReturnError(errors.New(connectionErrorMsg))

	repo := NewTrackRepository(db, 1*time.Second)

	_, err = repo.FindAll(context.Background())

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestTrackRepositoryFindAllSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(selectAllQuery).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "movie_id", "spotify_url"}).
			AddRow(trackID, trackName, trackMovieID, nil).
			AddRow("789e1011-e89b-12d3-a456-426614174222", "Concerning Hobbits", trackMovieID, nil))

	repo := NewTrackRepository(db, 1*time.Second)

	tracks, err := repo.FindAll(context.Background())

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Len(t, tracks, 2)
	assert.Equal(t, trackID, tracks[0].ID().String())
	assert.Equal(t, trackName, tracks[0].Name().String())
	assert.Equal(t, trackMovieID, tracks[0].MovieID().String())
}

func TestTrackRepositoryDeleteError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(deleteQuery).
		WithArgs(trackID).
		WillReturnError(errors.New(connectionErrorMsg))

	repo := NewTrackRepository(db, 1*time.Second)

	trackIDVO, err := domain.NewTrackIDFromString(trackID)
	require.NoError(t, err)
	err = repo.Delete(context.Background(), trackIDVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestTrackRepositoryDeleteNotFound(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(deleteQuery).
		WithArgs(trackID).
		WillReturnResult(sqlmock.NewResult(0, 0))

	repo := NewTrackRepository(db, 1*time.Second)

	trackIDVO, err := domain.NewTrackIDFromString(trackID)
	require.NoError(t, err)
	err = repo.Delete(context.Background(), trackIDVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.ErrorIs(t, err, domain.ErrTrackNotFound)
}

func TestTrackRepositoryDeleteSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(deleteQuery).
		WithArgs(trackID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewTrackRepository(db, 1*time.Second)

	trackIDVO, err := domain.NewTrackIDFromString(trackID)
	require.NoError(t, err)
	err = repo.Delete(context.Background(), trackIDVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func TestTrackRepositoryUpdateError(t *testing.T) {
	track, err := domain.NewTrackWithID(trackID, trackName, trackMovieID, nil)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("UPDATE tracks SET id = $1, name = $2, movie_id = $3, spotify_url = $4 WHERE id = $5").
		WithArgs(trackID, trackName, trackMovieID, nil, trackID).
		WillReturnError(errors.New("update error"))

	repo := NewTrackRepository(db, 1*time.Second)

	err = repo.Update(context.Background(), track)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestTrackRepositoryUpdateSuccess(t *testing.T) {
	track, err := domain.NewTrackWithID(trackID, trackName, trackMovieID, nil)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("UPDATE tracks SET id = $1, name = $2, movie_id = $3, spotify_url = $4 WHERE id = $5").
		WithArgs(trackID, trackName, trackMovieID, nil, trackID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewTrackRepository(db, 1*time.Second)

	err = repo.Update(context.Background(), track)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func TestTrackRepositoryFindByMovieError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery("SELECT tracks.id, tracks.name, tracks.movie_id, tracks.spotify_url FROM tracks WHERE movie_id = $1").
		WillReturnError(errors.New(connectionErrorMsg))

	repo := NewTrackRepository(db, 1*time.Second)

	movieIDVO, err := domain.NewMovieIDFromString(trackMovieID)
	require.NoError(t, err)

	_, err = repo.FindByMovie(context.Background(), movieIDVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestTrackRepositoryFindByMovieSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery("SELECT tracks.id, tracks.name, tracks.movie_id, tracks.spotify_url FROM tracks WHERE movie_id = $1").
		WithArgs(trackMovieID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "movie_id", "spotify_url"}).
			AddRow(trackID, trackName, trackMovieID, nil).
			AddRow("789e1011-e89b-12d3-a456-426614174222", "Concerning Hobbits", trackMovieID, nil))

	repo := NewTrackRepository(db, 1*time.Second)

	movieIDVO, err := domain.NewMovieIDFromString(trackMovieID)
	require.NoError(t, err)

	_, err = repo.FindByMovie(context.Background(), movieIDVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

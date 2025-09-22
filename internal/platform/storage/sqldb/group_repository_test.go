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

const groupID = "123e4567-e89b-12d3-a456-426614174000"
const groupName = "Fellowship of the Ring"
const groupDescription = "A group formed to destroy the One Ring"
const groupImageURL = "http://example.com/image.jpg"
const querySelectAllGroups = "SELECT groups.id, groups.name, groups.description, groups.image_url FROM groups ORDER BY created_at ASC"

func TestGroupRepositorySaveRepositoryError(t *testing.T) {
	group, err := domain.NewGroupWithID(groupID, groupName, groupDescription, groupImageURL)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO groups (id, name, description, image_url) VALUES ($1, $2, $3, $4)").
		WithArgs(groupID, groupName, groupDescription, groupImageURL).
		WillReturnError(errors.New("database error"))

	repo := NewGroupRepository(db, 1*time.Second)

	err = repo.Save(context.Background(), group)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestGroupRepositorySaveSuccess(t *testing.T) {
	group, err := domain.NewGroupWithID(groupID, groupName, groupDescription, groupImageURL)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO groups (id, name, description, image_url) VALUES ($1, $2, $3, $4)").
		WithArgs(groupID, groupName, groupDescription, groupImageURL).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewGroupRepository(db, 1*time.Second)

	err = repo.Save(context.Background(), group)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func TestGroupRepositoryFindNotFound(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery("SELECT groups.id, groups.name, groups.description, groups.image_url FROM groups WHERE id = $1").
		WithArgs(groupID).
		WillReturnError(sql.ErrNoRows)

	repo := NewGroupRepository(db, 1*time.Second)

	groupIDObj, err := domain.NewGroupIDFromString(groupID)
	require.NoError(t, err)

	_, err = repo.Find(context.Background(), groupIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestGroupRepositoryFindSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery("SELECT groups.id, groups.name, groups.description, groups.image_url FROM groups WHERE id = $1").
		WithArgs(groupID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "image_url"}).AddRow(groupID, groupName, groupDescription, groupImageURL))

	repo := NewGroupRepository(db, 1*time.Second)

	groupIDObj, err := domain.NewGroupIDFromString(groupID)
	require.NoError(t, err)

	group, err := repo.Find(context.Background(), groupIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Equal(t, groupIDObj, group.ID())
	assert.Equal(t, groupName, group.Name().String())
}

func TestGroupRepositoryFindAllSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectAllGroups).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "image_url"}).
			AddRow(groupID, groupName, groupDescription, groupImageURL).
			AddRow("223e4567-e89b-12d3-a456-426614174001", "Company of the Ring", "Description", "http://example.com/image.jpg"))

	repo := NewGroupRepository(db, 1*time.Second)

	groups, err := repo.FindAll(context.Background())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Len(t, groups, 2)
	assert.Equal(t, groupID, groups[0].ID().String())
	assert.Equal(t, groupName, groups[0].Name().String())
	assert.Equal(t, "223e4567-e89b-12d3-a456-426614174001", groups[1].ID().String())
	assert.Equal(t, "Company of the Ring", groups[1].Name().String())
}

func TestGroupRepositoryFindAllEmpty(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectAllGroups).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "image_url"}))

	repo := NewGroupRepository(db, 1*time.Second)

	groups, err := repo.FindAll(context.Background())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Len(t, groups, 0)
}

func TestGroupRepositoryFindAllQueryError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectAllGroups).
		WillReturnError(errors.New("query error"))

	repo := NewGroupRepository(db, 1*time.Second)

	groups, err := repo.FindAll(context.Background())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.Nil(t, groups)
}

func TestGroupRepositoryDeleteError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("DELETE FROM groups WHERE id = $1").
		WithArgs(groupID).
		WillReturnError(errors.New("delete error"))

	repo := NewGroupRepository(db, 1*time.Second)

	groupIDObj, err := domain.NewGroupIDFromString(groupID)
	require.NoError(t, err)

	err = repo.Delete(context.Background(), groupIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestGroupRepositoryDeleteSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("DELETE FROM groups WHERE id = $1").
		WithArgs(groupID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewGroupRepository(db, 1*time.Second)

	groupIDObj, err := domain.NewGroupIDFromString(groupID)
	require.NoError(t, err)

	err = repo.Delete(context.Background(), groupIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func TestGroupRepositoryUpdateError(t *testing.T) {
	group, err := domain.NewGroupWithID(groupID, groupName, groupDescription, groupImageURL)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("UPDATE groups SET id = $1, name = $2, description = $3, image_url = $4 WHERE id = $5").
		WithArgs(groupID, groupName, groupDescription, groupImageURL, groupID).
		WillReturnError(errors.New("update error"))

	repo := NewGroupRepository(db, 1*time.Second)

	err = repo.Update(context.Background(), group)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestGroupRepositoryUpdateSuccess(t *testing.T) {
	group, err := domain.NewGroupWithID(groupID, groupName, groupDescription, groupImageURL)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("UPDATE groups SET id = $1, name = $2, description = $3, image_url = $4 WHERE id = $5").
		WithArgs(groupID, groupName, groupDescription, groupImageURL, groupID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewGroupRepository(db, 1*time.Second)

	err = repo.Update(context.Background(), group)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

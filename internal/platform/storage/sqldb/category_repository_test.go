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

const categoryID = "123e4567-e89b-12d3-a456-426614174000"
const categoryName = "Fantasy"
const querySelectAllCategories = "SELECT categories.id, categories.name FROM categories"

func TestCategoryRepositorySaveRepositoryError(t *testing.T) {
	category, err := domain.NewCategoryWithID(categoryID, categoryName)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO categories (id, name) VALUES ($1, $2)").
		WithArgs(categoryID, categoryName).
		WillReturnError(errors.New("database error"))

	repo := NewCategoryRepository(db, 1*time.Second)

	err = repo.Save(context.Background(), category)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestCategoryRepositorySaveSuccess(t *testing.T) {
	category, err := domain.NewCategoryWithID(categoryID, categoryName)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO categories (id, name) VALUES ($1, $2)").
		WithArgs(categoryID, categoryName).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewCategoryRepository(db, 1*time.Second)

	err = repo.Save(context.Background(), category)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func TestCategoryRepositoryFindNotFound(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery("SELECT categories.id, categories.name FROM categories WHERE id = $1").
		WithArgs(categoryID).
		WillReturnError(sql.ErrNoRows)

	repo := NewCategoryRepository(db, 1*time.Second)

	categoryIDObj, err := domain.NewCategoryIDFromString(categoryID)
	require.NoError(t, err)

	_, err = repo.Find(context.Background(), categoryIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestCategoryRepositoryFindSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery("SELECT categories.id, categories.name FROM categories WHERE id = $1").
		WithArgs(categoryID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(categoryID, categoryName))

	repo := NewCategoryRepository(db, 1*time.Second)

	categoryIDObj, err := domain.NewCategoryIDFromString(categoryID)
	require.NoError(t, err)

	category, err := repo.Find(context.Background(), categoryIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Equal(t, categoryIDObj, category.ID())
	assert.Equal(t, categoryName, category.Name().String())
}

func TestCategoryRepositoryFindAllSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectAllCategories).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(categoryID, categoryName).
			AddRow("223e4567-e89b-12d3-a456-426614174001", "Action"))

	repo := NewCategoryRepository(db, 1*time.Second)

	categories, err := repo.FindAll(context.Background())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Len(t, categories, 2)
	assert.Equal(t, categoryID, categories[0].ID().String())
	assert.Equal(t, categoryName, categories[0].Name().String())
	assert.Equal(t, "223e4567-e89b-12d3-a456-426614174001", categories[1].ID().String())
	assert.Equal(t, "Action", categories[1].Name().String())
}

func TestCategoryRepositoryFindAllEmpty(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectAllCategories).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))

	repo := NewCategoryRepository(db, 1*time.Second)

	categories, err := repo.FindAll(context.Background())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Len(t, categories, 0)
}

func TestCategoryRepositoryFindAllQueryError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(querySelectAllCategories).
		WillReturnError(errors.New("query error"))

	repo := NewCategoryRepository(db, 1*time.Second)

	categories, err := repo.FindAll(context.Background())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.Nil(t, categories)
}

func TestCategoryRepositoryDeleteError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("DELETE FROM categories WHERE id = $1").
		WithArgs(categoryID).
		WillReturnError(errors.New("delete error"))

	repo := NewCategoryRepository(db, 1*time.Second)

	categoryIDObj, err := domain.NewCategoryIDFromString(categoryID)
	require.NoError(t, err)

	err = repo.Delete(context.Background(), categoryIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestCategoryRepositoryDeleteSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("DELETE FROM categories WHERE id = $1").
		WithArgs(categoryID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewCategoryRepository(db, 1*time.Second)

	categoryIDObj, err := domain.NewCategoryIDFromString(categoryID)
	require.NoError(t, err)

	err = repo.Delete(context.Background(), categoryIDObj)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func TestCategoryRepositoryUpdateError(t *testing.T) {
	category, err := domain.NewCategoryWithID(categoryID, categoryName)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("UPDATE categories SET id = $1, name = $2 WHERE id = $3").
		WithArgs(categoryID, categoryName, categoryID).
		WillReturnError(errors.New("update error"))

	repo := NewCategoryRepository(db, 1*time.Second)

	err = repo.Update(context.Background(), category)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestCategoryRepositoryUpdateSuccess(t *testing.T) {
	category, err := domain.NewCategoryWithID(categoryID, categoryName)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("UPDATE categories SET id = $1, name = $2 WHERE id = $3").
		WithArgs(categoryID, categoryName, categoryID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewCategoryRepository(db, 1*time.Second)

	err = repo.Update(context.Background(), category)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

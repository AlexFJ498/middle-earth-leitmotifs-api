package deleting

import (
	"context"
	"fmt"
	"testing"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	uuidStr          = "123e4567-e89b-12d3-a456-426614174000"
	databaseErrorMsg = "database error"
)

func TestMovieServiceDeleteMovieRepositoryError(t *testing.T) {
	movieIDObj, err := domain.NewMovieIDFromString(uuidStr)
	require.NoError(t, err)

	mockRepo := new(storagemocks.MovieRepository)
	mockRepo.On("Delete", mock.Anything, movieIDObj).Return(fmt.Errorf("%s", databaseErrorMsg))

	service := NewMovieService(mockRepo)

	err = service.DeleteMovie(context.Background(), movieIDObj)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), databaseErrorMsg)

	mockRepo.AssertExpectations(t)
}

func TestMovieServiceDeleteMovieSuccess(t *testing.T) {
	movieIDObj, err := domain.NewMovieIDFromString(uuidStr)
	require.NoError(t, err)

	mockRepo := new(storagemocks.MovieRepository)
	mockRepo.On("Delete", mock.Anything, movieIDObj).Return(nil)

	service := NewMovieService(mockRepo)

	err = service.DeleteMovie(context.Background(), movieIDObj)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGroupServiceDeleteGroupRepositoryError(t *testing.T) {
	groupIDObj, err := domain.NewGroupIDFromString(uuidStr)
	require.NoError(t, err)

	mockRepo := new(storagemocks.GroupRepository)
	mockRepo.On("Delete", mock.Anything, groupIDObj).Return(fmt.Errorf("%s", databaseErrorMsg))

	service := NewGroupService(mockRepo)

	err = service.DeleteGroup(context.Background(), groupIDObj)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database error")

	mockRepo.AssertExpectations(t)
}

func TestGroupServiceDeleteGroupSuccess(t *testing.T) {
	groupIDObj, err := domain.NewGroupIDFromString(uuidStr)
	require.NoError(t, err)

	mockRepo := new(storagemocks.GroupRepository)
	mockRepo.On("Delete", mock.Anything, groupIDObj).Return(nil)

	service := NewGroupService(mockRepo)

	err = service.DeleteGroup(context.Background(), groupIDObj)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCategoryServiceDeleteCategoryRepositoryError(t *testing.T) {
	categoryIDObj, err := domain.NewCategoryIDFromString(uuidStr)
	require.NoError(t, err)

	mockRepo := new(storagemocks.CategoryRepository)
	mockRepo.On("Delete", mock.Anything, categoryIDObj).Return(fmt.Errorf("%s", databaseErrorMsg))

	service := NewCategoryService(mockRepo)

	err = service.DeleteCategory(context.Background(), categoryIDObj)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), databaseErrorMsg)

	mockRepo.AssertExpectations(t)
}

func TestCategoryServiceDeleteCategorySuccess(t *testing.T) {
	categoryIDObj, err := domain.NewCategoryIDFromString(uuidStr)
	require.NoError(t, err)

	mockRepo := new(storagemocks.CategoryRepository)
	mockRepo.On("Delete", mock.Anything, categoryIDObj).Return(nil)

	service := NewCategoryService(mockRepo)

	err = service.DeleteCategory(context.Background(), categoryIDObj)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestTrackServiceDeleteTrackRepositoryError(t *testing.T) {
	trackIDObj, err := domain.NewTrackIDFromString(uuidStr)
	require.NoError(t, err)

	mockRepo := new(storagemocks.TrackRepository)
	mockRepo.On("Delete", mock.Anything, trackIDObj).Return(fmt.Errorf("%s", databaseErrorMsg))

	service := NewTrackService(mockRepo)

	err = service.DeleteTrack(context.Background(), trackIDObj)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), databaseErrorMsg)

	mockRepo.AssertExpectations(t)
}

func TestTrackServiceDeleteTrackSuccess(t *testing.T) {
	trackIDObj, err := domain.NewTrackIDFromString(uuidStr)
	require.NoError(t, err)

	mockRepo := new(storagemocks.TrackRepository)
	mockRepo.On("Delete", mock.Anything, trackIDObj).Return(nil)

	service := NewTrackService(mockRepo)

	err = service.DeleteTrack(context.Background(), trackIDObj)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestThemeServiceDeleteThemeRepositoryError(t *testing.T) {
	themeIDObj, err := domain.NewThemeIDFromString(uuidStr)
	require.NoError(t, err)

	mockRepo := new(storagemocks.ThemeRepository)
	mockRepo.On("Delete", mock.Anything, themeIDObj).Return(fmt.Errorf("%s", databaseErrorMsg))

	service := NewThemeService(mockRepo)

	err = service.DeleteTheme(context.Background(), themeIDObj)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), databaseErrorMsg)

	mockRepo.AssertExpectations(t)
}

func TestThemeServiceDeleteThemeSuccess(t *testing.T) {
	themeIDObj, err := domain.NewThemeIDFromString(uuidStr)
	require.NoError(t, err)

	mockRepo := new(storagemocks.ThemeRepository)
	mockRepo.On("Delete", mock.Anything, themeIDObj).Return(nil)

	service := NewThemeService(mockRepo)

	err = service.DeleteTheme(context.Background(), themeIDObj)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

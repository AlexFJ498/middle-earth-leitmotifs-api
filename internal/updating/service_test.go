package updating

import (
	"context"
	"errors"
	"testing"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	testID           = "123e4567-e89b-12d3-a456-426614174000"
	movieName        = "The Fellowship of the Ring"
	groupName        = "The Elves"
	groupDescription = "Description"
	groupImageURL    = "http://example.com/image.jpg"
	categoryName     = "The Mordor Accompaniments"
	trackName        = "The Three Hunters"
	themeName        = "The History of the Ring"
	themeDescription = "Description"

	domainMovieType    = "domain.Movie"
	domainGroupType    = "domain.Group"
	domainCategoryType = "domain.Category"
	domainTrackType    = "domain.Track"
	domainThemeType    = "domain.Theme"

	repositoryErrorMsg = "repository error"
	invalidId          = "invalid-id"
)

var categoryID = "456e7890-e89b-12d3-a456-426614174114"

func TestMovieServiceUpdateMovieRepositoryError(t *testing.T) {
	dto := dto.MovieUpdateRequest{
		Name: movieName,
	}

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainMovieType)).Return(errors.New(repositoryErrorMsg)).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	service := NewMovieService(movieRepositoryMock)

	err := service.UpdateMovie(context.Background(), testID, dto)
	assert.Error(t, err)
}

func TestMovieServiceUpdateMovieSuccess(t *testing.T) {
	dto := dto.MovieUpdateRequest{
		Name: movieName,
	}

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainMovieType)).Return(nil).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	service := NewMovieService(movieRepositoryMock)

	err := service.UpdateMovie(context.Background(), testID, dto)
	assert.NoError(t, err)
}

func TestMovieServiceUpdateMovieInvalidID(t *testing.T) {
	dto := dto.MovieUpdateRequest{
		Name: movieName,
	}

	movieRepositoryMock := new(storagemocks.MovieRepository)
	defer movieRepositoryMock.AssertExpectations(t)

	service := NewMovieService(movieRepositoryMock)

	err := service.UpdateMovie(context.Background(), invalidId, dto)
	assert.Error(t, err)
}

func TestGroupServiceUpdateGroupRepositoryError(t *testing.T) {
	dto := dto.GroupUpdateRequest{
		Name:        groupName,
		Description: groupDescription,
		ImageURL:    groupImageURL,
	}

	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainGroupType)).Return(errors.New(repositoryErrorMsg)).Once()
	defer groupRepositoryMock.AssertExpectations(t)

	service := NewGroupService(groupRepositoryMock)

	err := service.UpdateGroup(context.Background(), testID, dto)
	assert.Error(t, err)
}

func TestGroupServiceUpdateGroupSuccess(t *testing.T) {
	dto := dto.GroupUpdateRequest{
		Name:        groupName,
		Description: groupDescription,
		ImageURL:    groupImageURL,
	}

	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainGroupType)).Return(nil).Once()
	defer groupRepositoryMock.AssertExpectations(t)

	service := NewGroupService(groupRepositoryMock)

	err := service.UpdateGroup(context.Background(), testID, dto)
	assert.NoError(t, err)
}

func TestGroupServiceUpdateGroupInvalidID(t *testing.T) {
	dto := dto.GroupUpdateRequest{
		Name:        groupName,
		Description: groupDescription,
		ImageURL:    groupImageURL,
	}

	groupRepositoryMock := new(storagemocks.GroupRepository)
	defer groupRepositoryMock.AssertExpectations(t)

	service := NewGroupService(groupRepositoryMock)

	err := service.UpdateGroup(context.Background(), invalidId, dto)
	assert.Error(t, err)
}

func TestCategoryServiceUpdateCategoryRepositoryError(t *testing.T) {
	dto := dto.CategoryUpdateRequest{
		Name: categoryName,
	}

	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	categoryRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainCategoryType)).Return(errors.New(repositoryErrorMsg)).Once()
	defer categoryRepositoryMock.AssertExpectations(t)

	service := NewCategoryService(categoryRepositoryMock)

	err := service.UpdateCategory(context.Background(), testID, dto)
	assert.Error(t, err)
}

func TestCategoryServiceUpdateCategorySuccess(t *testing.T) {
	dto := dto.CategoryUpdateRequest{
		Name: categoryName,
	}

	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	categoryRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainCategoryType)).Return(nil).Once()
	defer categoryRepositoryMock.AssertExpectations(t)

	service := NewCategoryService(categoryRepositoryMock)

	err := service.UpdateCategory(context.Background(), testID, dto)
	assert.NoError(t, err)
}

func TestCategoryServiceUpdateCategoryInvalidID(t *testing.T) {
	dto := dto.CategoryUpdateRequest{
		Name: categoryName,
	}

	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	defer categoryRepositoryMock.AssertExpectations(t)

	service := NewCategoryService(categoryRepositoryMock)

	err := service.UpdateCategory(context.Background(), invalidId, dto)
	assert.Error(t, err)
}

func TestTrackServiceUpdateTrackRepositoryError(t *testing.T) {
	dto := dto.TrackUpdateRequest{
		Name:    trackName,
		MovieID: testID,
	}

	trackRepositoryMock := new(storagemocks.TrackRepository)
	trackRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainTrackType)).Return(errors.New(repositoryErrorMsg)).Once()
	defer trackRepositoryMock.AssertExpectations(t)

	service := NewTrackService(trackRepositoryMock)

	err := service.UpdateTrack(context.Background(), testID, dto)
	assert.Error(t, err)
}

func TestTrackServiceUpdateTrackSuccess(t *testing.T) {
	dto := dto.TrackUpdateRequest{
		Name:    trackName,
		MovieID: testID,
	}

	trackRepositoryMock := new(storagemocks.TrackRepository)
	trackRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainTrackType)).Return(nil).Once()
	defer trackRepositoryMock.AssertExpectations(t)

	service := NewTrackService(trackRepositoryMock)

	err := service.UpdateTrack(context.Background(), testID, dto)
	assert.NoError(t, err)
}

func TestTrackServiceUpdateTrackInvalidID(t *testing.T) {
	dto := dto.TrackUpdateRequest{
		Name:    trackName,
		MovieID: testID,
	}

	trackRepositoryMock := new(storagemocks.TrackRepository)
	defer trackRepositoryMock.AssertExpectations(t)

	service := NewTrackService(trackRepositoryMock)

	err := service.UpdateTrack(context.Background(), invalidId, dto)
	assert.Error(t, err)
}

func TestThemeServiceUpdateThemeRepositoryError(t *testing.T) {
	dto := dto.ThemeUpdateRequest{
		Name:        themeName,
		FirstHeard:  testID,
		GroupID:     testID,
		Description: themeDescription,
		CategoryID:  &categoryID,
	}

	themeRepositoryMock := new(storagemocks.ThemeRepository)
	themeRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainThemeType)).Return(errors.New(repositoryErrorMsg)).Once()
	defer themeRepositoryMock.AssertExpectations(t)

	service := NewThemeService(themeRepositoryMock)

	err := service.UpdateTheme(context.Background(), testID, dto)
	assert.Error(t, err)
}

func TestThemeServiceUpdateThemeSuccess(t *testing.T) {
	dto := dto.ThemeUpdateRequest{
		Name:        themeName,
		FirstHeard:  testID,
		GroupID:     testID,
		Description: themeDescription,
		CategoryID:  &categoryID,
	}

	themeRepositoryMock := new(storagemocks.ThemeRepository)
	themeRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType(domainThemeType)).Return(nil).Once()
	defer themeRepositoryMock.AssertExpectations(t)

	service := NewThemeService(themeRepositoryMock)

	err := service.UpdateTheme(context.Background(), testID, dto)
	assert.NoError(t, err)
}

func TestThemeServiceUpdateThemeInvalidID(t *testing.T) {
	dto := dto.ThemeUpdateRequest{
		Name:        themeName,
		FirstHeard:  testID,
		GroupID:     testID,
		Description: themeDescription,
		CategoryID:  &categoryID,
	}

	themeRepositoryMock := new(storagemocks.ThemeRepository)
	defer themeRepositoryMock.AssertExpectations(t)

	service := NewThemeService(themeRepositoryMock)

	err := service.UpdateTheme(context.Background(), invalidId, dto)
	assert.Error(t, err)
}

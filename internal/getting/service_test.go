package getting

import (
	"context"
	"errors"
	"testing"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const repositoryErrorMsg = "repository error"
const exampleUUID = "28712a35-04dd-4200-9316-4d6a1e399128"
const movieName = "The Fellowship of the Ring"
const trackName = "The Three Hunters"

func TestMovieServiceGetMovieRepositoryError(t *testing.T) {
	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Movie{}, errors.New(repositoryErrorMsg))
	defer movieRepositoryMock.AssertExpectations(t)

	movieService := NewMovieService(movieRepositoryMock)

	_, err := movieService.GetMovie(context.Background(), exampleUUID)
	assert.Error(t, err)
	assert.Equal(t, repositoryErrorMsg, err.Error())
}

func TestMovieServiceGetMovieNotFound(t *testing.T) {
	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Movie{}, domain.ErrMovieNotFound)
	defer movieRepositoryMock.AssertExpectations(t)

	movieService := NewMovieService(movieRepositoryMock)

	_, err := movieService.GetMovie(context.Background(), exampleUUID)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrMovieNotFound, err)
}

func TestMovieServiceGetMovieInvalidID(t *testing.T) {
	movieRepositoryMock := new(storagemocks.MovieRepository)
	defer movieRepositoryMock.AssertExpectations(t)
	movieService := NewMovieService(movieRepositoryMock)

	_, err := movieService.GetMovie(context.Background(), "invalid-movie-uuid")
	assert.Error(t, err)
	assert.Equal(t, "invalid movie ID", err.Error())
}

func TestMovieServiceGetMovieSuccess(t *testing.T) {
	movieRepositoryMock := new(storagemocks.MovieRepository)
	movie, err := domain.NewMovie(movieName)
	assert.NoError(t, err)

	movieRepositoryMock.On("Find", mock.Anything, movie.ID()).Return(movie, nil).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	movieService := NewMovieService(movieRepositoryMock)

	result, err := movieService.GetMovie(context.Background(), movie.ID().String())
	assert.NoError(t, err)
	assert.Equal(t, movieName, result.Name)
}

func TestGroupServiceGetGroupRepositoryError(t *testing.T) {
	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Group{}, errors.New(repositoryErrorMsg))
	defer groupRepositoryMock.AssertExpectations(t)

	groupService := NewGroupService(groupRepositoryMock)

	_, err := groupService.GetGroup(context.Background(), exampleUUID)
	assert.Error(t, err)
	assert.Equal(t, repositoryErrorMsg, err.Error())
}

func TestGroupServiceGetGroupNotFound(t *testing.T) {
	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Group{}, domain.ErrGroupNotFound)
	defer groupRepositoryMock.AssertExpectations(t)

	groupService := NewGroupService(groupRepositoryMock)

	_, err := groupService.GetGroup(context.Background(), exampleUUID)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrGroupNotFound, err)
}

func TestGroupServiceGetGroupInvalidID(t *testing.T) {
	groupRepositoryMock := new(storagemocks.GroupRepository)
	defer groupRepositoryMock.AssertExpectations(t)

	groupService := NewGroupService(groupRepositoryMock)

	_, err := groupService.GetGroup(context.Background(), "invalid-group-uuid")
	assert.Error(t, err)
	assert.Equal(t, "invalid group ID", err.Error())
}

func TestGroupServiceGetGroupSuccess(t *testing.T) {
	groupRepositoryMock := new(storagemocks.GroupRepository)
	group, err := domain.NewGroup("Orchestral", "Description of Orchestral", "http://example.com/orchestral.jpg")
	assert.NoError(t, err)

	groupRepositoryMock.On("Find", mock.Anything, group.ID()).Return(group, nil).Once()
	defer groupRepositoryMock.AssertExpectations(t)

	groupService := NewGroupService(groupRepositoryMock)

	result, err := groupService.GetGroup(context.Background(), group.ID().String())
	assert.NoError(t, err)
	assert.Equal(t, "Orchestral", result.Name)
}

func TestCategoryServiceGetCategoryRepositoryError(t *testing.T) {
	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	categoryRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Category{}, errors.New(repositoryErrorMsg))
	defer categoryRepositoryMock.AssertExpectations(t)

	categoryService := NewCategoryService(categoryRepositoryMock)

	_, err := categoryService.GetCategory(context.Background(), exampleUUID)
	assert.Error(t, err)
	assert.Equal(t, repositoryErrorMsg, err.Error())
}

func TestCategoryServiceGetCategoryNotFound(t *testing.T) {
	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	categoryRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Category{}, domain.ErrCategoryNotFound)
	defer categoryRepositoryMock.AssertExpectations(t)

	categoryService := NewCategoryService(categoryRepositoryMock)

	_, err := categoryService.GetCategory(context.Background(), exampleUUID)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrCategoryNotFound, err)
}

func TestCategoryServiceGetCategoryInvalidID(t *testing.T) {
	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	defer categoryRepositoryMock.AssertExpectations(t)

	categoryService := NewCategoryService(categoryRepositoryMock)

	_, err := categoryService.GetCategory(context.Background(), "invalid-category-uuid")
	assert.Error(t, err)
	assert.Equal(t, "invalid category ID", err.Error())
}

func TestCategoryServiceGetCategorySuccess(t *testing.T) {
	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	category, err := domain.NewCategory("Main Theme")
	assert.NoError(t, err)

	categoryRepositoryMock.On("Find", mock.Anything, category.ID()).Return(category, nil).Once()
	defer categoryRepositoryMock.AssertExpectations(t)

	categoryService := NewCategoryService(categoryRepositoryMock)

	result, err := categoryService.GetCategory(context.Background(), category.ID().String())
	assert.NoError(t, err)
	assert.Equal(t, "Main Theme", result.Name)
}

func TestTrackServiceGetTrackRepositoryError(t *testing.T) {
	trackRepositoryMock := new(storagemocks.TrackRepository)
	trackRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Track{}, errors.New(repositoryErrorMsg))
	defer trackRepositoryMock.AssertExpectations(t)

	movieService := NewMovieService(new(storagemocks.MovieRepository))

	trackService := NewTrackService(trackRepositoryMock, movieService)

	_, err := trackService.GetTrack(context.Background(), exampleUUID)
	assert.Error(t, err)
	assert.Equal(t, repositoryErrorMsg, err.Error())
}

func TestTrackServiceGetTrackNotFound(t *testing.T) {
	trackRepositoryMock := new(storagemocks.TrackRepository)
	trackRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Track{}, domain.ErrTrackNotFound)
	defer trackRepositoryMock.AssertExpectations(t)

	movieService := NewMovieService(new(storagemocks.MovieRepository))
	trackService := NewTrackService(trackRepositoryMock, movieService)

	_, err := trackService.GetTrack(context.Background(), exampleUUID)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrTrackNotFound, err)
}

func TestTrackServiceGetTrackInvalidID(t *testing.T) {
	trackRepositoryMock := new(storagemocks.TrackRepository)
	defer trackRepositoryMock.AssertExpectations(t)

	movieService := NewMovieService(new(storagemocks.MovieRepository))
	trackService := NewTrackService(trackRepositoryMock, movieService)

	_, err := trackService.GetTrack(context.Background(), "invalid-track-uuid")
	assert.Error(t, err)
	assert.Equal(t, "invalid track ID", err.Error())
}

func TestTrackServiceGetTrackMovieRepositoryError(t *testing.T) {
	trackRepositoryMock := new(storagemocks.TrackRepository)
	track, err := domain.NewTrack(trackName, exampleUUID, nil)
	assert.NoError(t, err)

	trackRepositoryMock.On("Find", mock.Anything, track.ID()).Return(track, nil).Once()
	defer trackRepositoryMock.AssertExpectations(t)

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Movie{}, errors.New(repositoryErrorMsg)).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	movieService := NewMovieService(movieRepositoryMock)
	trackService := NewTrackService(trackRepositoryMock, movieService)

	_, err = trackService.GetTrack(context.Background(), track.ID().String())
	assert.Error(t, err)
	assert.Equal(t, repositoryErrorMsg, err.Error())
}

func TestTrackServiceGetTrackSuccess(t *testing.T) {
	trackRepositoryMock := new(storagemocks.TrackRepository)
	track, err := domain.NewTrack(trackName, exampleUUID, nil)
	assert.NoError(t, err)

	trackRepositoryMock.On("Find", mock.Anything, track.ID()).Return(track, nil).Once()
	defer trackRepositoryMock.AssertExpectations(t)

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movie, err := domain.NewMovieWithID(exampleUUID, movieName)
	assert.NoError(t, err)
	movieRepositoryMock.On("Find", mock.Anything, movie.ID()).Return(movie, nil).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	movieService := NewMovieService(movieRepositoryMock)
	trackService := NewTrackService(trackRepositoryMock, movieService)

	result, err := trackService.GetTrack(context.Background(), track.ID().String())
	assert.NoError(t, err)
	assert.Equal(t, trackName, result.Name)
	assert.Equal(t, movieName, result.Movie.Name)
}

func TestThemeServiceGetThemeRepositoryError(t *testing.T) {
	themeRepositoryMock := new(storagemocks.ThemeRepository)
	themeRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Theme{}, errors.New(repositoryErrorMsg))
	defer themeRepositoryMock.AssertExpectations(t)

	trackService := NewTrackService(new(storagemocks.TrackRepository), NewMovieService(new(storagemocks.MovieRepository)))
	groupService := NewGroupService(new(storagemocks.GroupRepository))
	categoryService := NewCategoryService(new(storagemocks.CategoryRepository))
	themeService := NewThemeService(themeRepositoryMock, trackService, groupService, categoryService)

	_, err := themeService.GetTheme(context.Background(), exampleUUID)
	assert.Error(t, err)
	assert.Equal(t, repositoryErrorMsg, err.Error())
}

func TestThemeServiceGetThemeNotFound(t *testing.T) {
	themeRepositoryMock := new(storagemocks.ThemeRepository)
	themeRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Theme{}, domain.ErrThemeNotFound)
	defer themeRepositoryMock.AssertExpectations(t)

	trackService := NewTrackService(new(storagemocks.TrackRepository), NewMovieService(new(storagemocks.MovieRepository)))
	groupService := NewGroupService(new(storagemocks.GroupRepository))
	categoryService := NewCategoryService(new(storagemocks.CategoryRepository))
	themeService := NewThemeService(themeRepositoryMock, trackService, groupService, categoryService)

	_, err := themeService.GetTheme(context.Background(), exampleUUID)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrThemeNotFound, err)
}

func TestThemeServiceGetThemeInvalidID(t *testing.T) {
	themeRepositoryMock := new(storagemocks.ThemeRepository)
	defer themeRepositoryMock.AssertExpectations(t)

	trackService := NewTrackService(new(storagemocks.TrackRepository), NewMovieService(new(storagemocks.MovieRepository)))
	groupService := NewGroupService(new(storagemocks.GroupRepository))
	categoryService := NewCategoryService(new(storagemocks.CategoryRepository))
	themeService := NewThemeService(themeRepositoryMock, trackService, groupService, categoryService)

	_, err := themeService.GetTheme(context.Background(), "invalid-theme-uuid")
	assert.Error(t, err)
	assert.Equal(t, "invalid theme ID", err.Error())
}

func TestThemeServiceGetThemeSuccess(t *testing.T) {
	categoryID := "28712a35-04dd-4200-9316-4d6a1e399123"
	themeRepositoryMock := new(storagemocks.ThemeRepository)
	theme, err := domain.NewTheme("The Bridge of Khazad-dûm", exampleUUID, "28712a35-04dd-4200-9316-4d6a1e399122", "Description", &categoryID)
	assert.NoError(t, err)
	themeRepositoryMock.On("Find", mock.Anything, theme.ID()).Return(theme, nil).Once()
	defer themeRepositoryMock.AssertExpectations(t)

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieService := NewMovieService(movieRepositoryMock)

	track, err := domain.NewTrack(trackName, exampleUUID, nil)
	assert.NoError(t, err)

	trackRepositoryMock := new(storagemocks.TrackRepository)
	trackRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(track, nil).Once()
	trackService := NewTrackService(trackRepositoryMock, movieService)
	defer trackRepositoryMock.AssertExpectations(t)

	movieRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Movie{}, nil).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Group{}, nil).Once()
	groupService := NewGroupService(groupRepositoryMock)
	defer groupRepositoryMock.AssertExpectations(t)

	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	categoryRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Category{}, nil).Once()
	categoryService := NewCategoryService(categoryRepositoryMock)
	defer categoryRepositoryMock.AssertExpectations(t)

	themeService := NewThemeService(themeRepositoryMock, trackService, groupService, categoryService)

	result, err := themeService.GetTheme(context.Background(), theme.ID().String())
	assert.NoError(t, err)
	assert.Equal(t, "The Bridge of Khazad-dûm", result.Name)
	assert.Equal(t, trackName, result.FirstHeard.Name)
}

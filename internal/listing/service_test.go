package listing

import (
	"context"
	"errors"
	"testing"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/getting"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const repositoryErrorMsg = "repository error"

func TestUserServiceListUsersRepositoryError(t *testing.T) {
	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("FindAll", mock.Anything).Return(nil, errors.New(repositoryErrorMsg)).Once()
	defer userRepositoryMock.AssertExpectations(t)

	userService := NewUserService(userRepositoryMock)

	ctx := context.Background()
	_, err := userService.ListUsers(ctx)
	assert.Error(t, err)
}

func TestUserServiceListUsersSuccess(t *testing.T) {
	userRepositoryMock := new(storagemocks.UserRepository)
	users := []domain.User{}
	user1, err := domain.NewUser("John Doe", "john@example.com", "password123", false)
	assert.NoError(t, err)
	users = append(users, user1)
	user2, err := domain.NewUser("Jane Doe", "jane@example.com", "password456", false)
	assert.NoError(t, err)
	users = append(users, user2)
	userRepositoryMock.On("FindAll", mock.Anything).Return(users, nil).Once()
	defer userRepositoryMock.AssertExpectations(t)

	userService := NewUserService(userRepositoryMock)

	usersDTO, err := userService.ListUsers(context.Background())
	assert.NoError(t, err)
	assert.Len(t, usersDTO, 2)
	assert.Equal(t, "John Doe", usersDTO[0].Name)
	assert.Equal(t, "john@example.com", usersDTO[0].Email)
	assert.Equal(t, "Jane Doe", usersDTO[1].Name)
	assert.Equal(t, "jane@example.com", usersDTO[1].Email)
}

func TestMovieServiceListMoviesRepositoryError(t *testing.T) {
	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("FindAll", mock.Anything).Return(nil, errors.New(repositoryErrorMsg)).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	movieService := NewMovieService(movieRepositoryMock)

	ctx := context.Background()
	_, err := movieService.ListMovies(ctx)
	assert.Error(t, err)
}

func TestMovieServiceListMoviesSuccess(t *testing.T) {
	movieRepositoryMock := new(storagemocks.MovieRepository)
	movies := []domain.Movie{}
	movie1, err := domain.NewMovie("The Fellowship of the Ring")
	assert.NoError(t, err)
	movies = append(movies, movie1)
	movie2, err := domain.NewMovie("The Two Towers")
	assert.NoError(t, err)
	movies = append(movies, movie2)
	movieRepositoryMock.On("FindAll", mock.Anything).Return(movies, nil).Once()
	defer movieRepositoryMock.AssertExpectations(t)

	movieService := NewMovieService(movieRepositoryMock)

	moviesDTO, err := movieService.ListMovies(context.Background())
	assert.NoError(t, err)
	assert.Len(t, moviesDTO, 2)
	assert.Equal(t, "The Fellowship of the Ring", moviesDTO[0].Name)
	assert.Equal(t, "The Two Towers", moviesDTO[1].Name)
}

func TestGroupServiceListGroupsRepositoryError(t *testing.T) {
	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("FindAll", mock.Anything).Return(nil, errors.New(repositoryErrorMsg)).Once()
	defer groupRepositoryMock.AssertExpectations(t)

	groupService := NewGroupService(groupRepositoryMock)

	ctx := context.Background()
	_, err := groupService.ListGroups(ctx)
	assert.Error(t, err)
}

func TestGroupServiceListGroupsSuccess(t *testing.T) {
	groupRepositoryMock := new(storagemocks.GroupRepository)
	groups := []domain.Group{}
	group1, err := domain.NewGroup("The Elves", "Description of The Elves", "http://example.com/elves.jpg")
	assert.NoError(t, err)
	groups = append(groups, group1)
	group2, err := domain.NewGroup("Rohan", "Description of Rohan", "http://example.com/rohan.jpg")
	assert.NoError(t, err)
	groups = append(groups, group2)
	groupRepositoryMock.On("FindAll", mock.Anything).Return(groups, nil).Once()
	defer groupRepositoryMock.AssertExpectations(t)

	groupService := NewGroupService(groupRepositoryMock)

	groupsDTO, err := groupService.ListGroups(context.Background())
	assert.NoError(t, err)
	assert.Len(t, groupsDTO, 2)
	assert.Equal(t, "The Elves", groupsDTO[0].Name)
	assert.Equal(t, "Rohan", groupsDTO[1].Name)
}

func TestCategoryServiceListCategoriesRepositoryError(t *testing.T) {
	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	categoryRepositoryMock.On("FindAll", mock.Anything).Return(nil, errors.New("repository error")).Once()
	defer categoryRepositoryMock.AssertExpectations(t)

	categoryService := NewCategoryService(categoryRepositoryMock)

	ctx := context.Background()
	_, err := categoryService.ListCategories(ctx)
	assert.Error(t, err)
}

func TestCategoryServiceListCategoriesSuccess(t *testing.T) {
	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	categories := []domain.Category{}
	category1, err := domain.NewCategory("The Mordor Accompaniments")
	assert.NoError(t, err)
	categories = append(categories, category1)
	category2, err := domain.NewCategory("The Hobbit Accompaniments")
	assert.NoError(t, err)
	categories = append(categories, category2)
	categoryRepositoryMock.On("FindAll", mock.Anything).Return(categories, nil).Once()
	defer categoryRepositoryMock.AssertExpectations(t)

	categoryService := NewCategoryService(categoryRepositoryMock)

	categoriesDTO, err := categoryService.ListCategories(context.Background())
	assert.NoError(t, err)
	assert.Len(t, categoriesDTO, 2)
	assert.Equal(t, "The Mordor Accompaniments", categoriesDTO[0].Name)
	assert.Equal(t, "The Hobbit Accompaniments", categoriesDTO[1].Name)
}

func TestTrackServiceListTracksRepositoryError(t *testing.T) {
	trackRepositoryMock := new(storagemocks.TrackRepository)
	trackRepositoryMock.On("FindAll", mock.Anything).Return(nil, errors.New(repositoryErrorMsg)).Once()
	defer trackRepositoryMock.AssertExpectations(t)

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieService := NewMovieService(movieRepositoryMock)
	gettingMovieService := getting.NewMovieService(movieRepositoryMock)

	trackService := NewTrackService(trackRepositoryMock, movieService, gettingMovieService)

	ctx := context.Background()
	_, err := trackService.ListTracks(ctx)
	assert.Error(t, err)
}

func TestTrackServiceListTracksSuccess(t *testing.T) {
	trackRepositoryMock := new(storagemocks.TrackRepository)
	tracks := []domain.Track{}
	track1, err := domain.NewTrack("The Three Hunters", "b6c9d5ae-bf3b-419e-ba8f-09c8ce39d9bc", nil)
	assert.NoError(t, err)
	tracks = append(tracks, track1)
	track2, err := domain.NewTrack("The Shire", "28712a55-04dd-4200-9316-4d6a1e399128", nil)
	assert.NoError(t, err)
	tracks = append(tracks, track2)
	trackRepositoryMock.On("FindAll", mock.Anything).Return(tracks, nil).Once()
	defer trackRepositoryMock.AssertExpectations(t)

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Movie{}, nil).Twice()
	movieService := NewMovieService(movieRepositoryMock)
	gettingMovieService := getting.NewMovieService(movieRepositoryMock)

	trackService := NewTrackService(trackRepositoryMock, movieService, gettingMovieService)

	tracksDTO, err := trackService.ListTracks(context.Background())
	assert.NoError(t, err)
	assert.Len(t, tracksDTO, 2)
	assert.Equal(t, "The Three Hunters", tracksDTO[0].Name)
	assert.Equal(t, "The Shire", tracksDTO[1].Name)
}

func TestThemeServiceListThemesRepositoryError(t *testing.T) {
	themeRepositoryMock := new(storagemocks.ThemeRepository)
	themeRepositoryMock.On("FindAll", mock.Anything).Return(nil, errors.New(repositoryErrorMsg)).Once()
	defer themeRepositoryMock.AssertExpectations(t)

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieService := NewMovieService(movieRepositoryMock)
	gettingMovieService := getting.NewMovieService(movieRepositoryMock)

	trackRepositoryMock := new(storagemocks.TrackRepository)
	trackService := NewTrackService(trackRepositoryMock, movieService, gettingMovieService)
	gettingTrackService := getting.NewTrackService(trackRepositoryMock, gettingMovieService)

	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupService := NewGroupService(groupRepositoryMock)
	gettingGroupService := getting.NewGroupService(groupRepositoryMock)

	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	categoryService := NewCategoryService(categoryRepositoryMock)
	gettingCategoryService := getting.NewCategoryService(categoryRepositoryMock)

	themeService := NewThemeService(themeRepositoryMock, trackService, groupService, categoryService, gettingGroupService, gettingTrackService, gettingCategoryService)

	ctx := context.Background()
	_, err := themeService.ListThemes(ctx)
	assert.Error(t, err)
}

func TestThemeServiceListThemesSuccess(t *testing.T) {
	categoryID1 := "40929ca6-ed89-4548-a1d9-54b604ea50b5"
	categoryID2 := "40929ca6-ed89-4548-a1d9-54b604ea50b6"
	themeRepositoryMock := new(storagemocks.ThemeRepository)
	themes := []domain.Theme{}
	theme1, err := domain.NewTheme("The History of the Ring", "6a4f86e4-4fef-4151-9c60-e467007dd213", "40929ca6-ed89-4548-a1d9-54b604ea50b2", "Description", &categoryID1)
	assert.NoError(t, err)
	themes = append(themes, theme1)
	theme2, err := domain.NewTheme("The Rohan Fanfare", "6a4f86e4-4fef-4151-9c60-e467007dd213", "40929ca6-ed89-4548-a1d9-54b604ea50b2", "Description", &categoryID2)
	assert.NoError(t, err)
	themes = append(themes, theme2)
	themeRepositoryMock.On("FindAll", mock.Anything).Return(themes, nil).Once()
	defer themeRepositoryMock.AssertExpectations(t)

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieService := NewMovieService(movieRepositoryMock)
	gettingMovieService := getting.NewMovieService(movieRepositoryMock)

	track, err := domain.NewTrack("Track", "28712a55-04dd-4200-9316-4d6a1e399128", nil)
	assert.NoError(t, err)

	trackRepositoryMock := new(storagemocks.TrackRepository)
	trackRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(track, nil).Twice()
	trackService := NewTrackService(trackRepositoryMock, movieService, gettingMovieService)
	defer trackRepositoryMock.AssertExpectations(t)

	movieRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Movie{}, nil).Twice()
	gettingTrackService := getting.NewTrackService(trackRepositoryMock, gettingMovieService)
	defer movieRepositoryMock.AssertExpectations(t)

	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Group{}, nil).Twice()
	groupService := NewGroupService(groupRepositoryMock)
	gettingGroupService := getting.NewGroupService(groupRepositoryMock)
	defer groupRepositoryMock.AssertExpectations(t)

	categoryRepositoryMock := new(storagemocks.CategoryRepository)
	categoryRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Category{}, nil).Twice()
	categoryService := NewCategoryService(categoryRepositoryMock)
	gettingCategoryService := getting.NewCategoryService(categoryRepositoryMock)
	defer categoryRepositoryMock.AssertExpectations(t)

	themeService := NewThemeService(themeRepositoryMock, trackService, groupService, categoryService, gettingGroupService, gettingTrackService, gettingCategoryService)

	themesDTO, err := themeService.ListThemes(context.Background())
	assert.NoError(t, err)
	assert.Len(t, themesDTO, 2)
	assert.Equal(t, "The History of the Ring", themesDTO[0].Name)
	assert.Equal(t, "The Rohan Fanfare", themesDTO[1].Name)
}

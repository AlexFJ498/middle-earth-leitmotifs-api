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

func getGettingMovieServiceMock(movieRepositoryMock *storagemocks.MovieRepository) getting.MovieService {
	return getting.NewMovieService(movieRepositoryMock)
}

func getGettingTrackServiceMock(trackRepositoryMock *storagemocks.TrackRepository, movieRepositoryMock *storagemocks.MovieRepository) getting.TrackService {
	gettingMovieService := getGettingMovieServiceMock(movieRepositoryMock)
	return getting.NewTrackService(trackRepositoryMock, gettingMovieService)
}

func getGettingGroupServiceMock(groupRepositoryMock *storagemocks.GroupRepository) getting.GroupService {
	return getting.NewGroupService(groupRepositoryMock)
}

func getGettingCategoryServiceMock(categoryRepositoryMock *storagemocks.CategoryRepository) getting.CategoryService {
	return getting.NewCategoryService(categoryRepositoryMock)
}

func getGettingThemeServiceMock(themeRepositoryMock *storagemocks.ThemeRepository, trackRepositoryMock *storagemocks.TrackRepository, groupRepositoryMock *storagemocks.GroupRepository, movieRepositoryMock *storagemocks.MovieRepository, categoryRepositoryMock *storagemocks.CategoryRepository) getting.ThemeService {
	return getting.NewThemeService(
		themeRepositoryMock,
		getGettingTrackServiceMock(trackRepositoryMock, movieRepositoryMock),
		getGettingGroupServiceMock(groupRepositoryMock),
		getGettingCategoryServiceMock(categoryRepositoryMock),
	)
}

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

func TestTrackServiceListTracksByMovieRepositoryError(t *testing.T) {
	trackRepositoryMock := new(storagemocks.TrackRepository)
	trackRepositoryMock.On("FindByMovie", mock.Anything, mock.Anything).Return(nil, errors.New(repositoryErrorMsg)).Once()
	defer trackRepositoryMock.AssertExpectations(t)

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieService := NewMovieService(movieRepositoryMock)
	gettingMovieService := getting.NewMovieService(movieRepositoryMock)

	trackService := NewTrackService(trackRepositoryMock, movieService, gettingMovieService)

	ctx := context.Background()
	_, err := trackService.ListTracksByMovie(ctx, "12345678-1234-1234-1234-123456789012")
	assert.Error(t, err)
}

func TestTrackServiceListTracksByMovieSuccess(t *testing.T) {
	trackRepositoryMock := new(storagemocks.TrackRepository)
	tracks := []domain.Track{}
	track1, err := domain.NewTrack("Track name", "b6c1d5ae-bf3b-419e-ba8f-09c8ce39d9bc", nil)
	assert.NoError(t, err)
	tracks = append(tracks, track1)
	track2, err := domain.NewTrack("Track 2 name", "22712a55-04dd-4200-9316-4d6a1e399128", nil)
	assert.NoError(t, err)
	tracks = append(tracks, track2)
	trackRepositoryMock.On("FindByMovie", mock.Anything, mock.Anything).Return(tracks, nil).Once()
	defer trackRepositoryMock.AssertExpectations(t)

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(domain.Movie{}, nil).Twice()
	movieService := NewMovieService(movieRepositoryMock)
	gettingMovieService := getting.NewMovieService(movieRepositoryMock)

	trackService := NewTrackService(trackRepositoryMock, movieService, gettingMovieService)

	tracksDTO, err := trackService.ListTracksByMovie(context.Background(), "12345678-1234-1234-1234-123456789012")
	assert.NoError(t, err)
	assert.Len(t, tracksDTO, 2)
	assert.Equal(t, "Track name", tracksDTO[0].Name)
	assert.Equal(t, "Track 2 name", tracksDTO[1].Name)
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
	trackID := "6a4f86e4-4fef-4151-9c60-e467007dd213"
	groupID := "40929ca6-ed89-4548-a1d9-54b604ea50b2"
	themeRepositoryMock := new(storagemocks.ThemeRepository)
	themes := []domain.Theme{}
	theme1, err := domain.NewTheme("The History of the Ring", trackID, groupID, "Description", 0, 1, &categoryID1)
	assert.NoError(t, err)
	themes = append(themes, theme1)
	theme2, err := domain.NewTheme("The Rohan Fanfare", trackID, groupID, "Description", 0, 1, &categoryID2)
	assert.NoError(t, err)
	themes = append(themes, theme2)
	themeRepositoryMock.On("FindAll", mock.Anything).Return(themes, nil).Once()
	defer themeRepositoryMock.AssertExpectations(t)

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieService := NewMovieService(movieRepositoryMock)
	gettingMovieService := getting.NewMovieService(movieRepositoryMock)

	track, err := domain.NewTrack("Track", "28712a55-04dd-4200-9316-4d6a1e399122", nil)
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

func TestTrackThemeServiceListTrackThemesRepositoryError(t *testing.T) {
	trackThemeRepositoryMock := new(storagemocks.TrackThemeRepository)
	trackThemeRepositoryMock.On("FindByTrack", mock.Anything, mock.Anything).Return(nil, errors.New(repositoryErrorMsg)).Once()
	defer trackThemeRepositoryMock.AssertExpectations(t)

	trackRepositoryMock := new(storagemocks.TrackRepository)

	themeRepositoryMock := new(storagemocks.ThemeRepository)
	groupRepositoryMock := new(storagemocks.GroupRepository)
	movieRepositoryMock := new(storagemocks.MovieRepository)
	categoryRepositoryMock := new(storagemocks.CategoryRepository)

	trackThemeService := NewTrackThemeService(trackThemeRepositoryMock, getGettingTrackServiceMock(trackRepositoryMock, movieRepositoryMock), getGettingThemeServiceMock(themeRepositoryMock, trackRepositoryMock, groupRepositoryMock, movieRepositoryMock, categoryRepositoryMock))

	ctx := context.Background()
	_, err := trackThemeService.ListTracksThemesByTrack(ctx, "28712a55-04dd-4200-9316-4d6a1e399128")
	assert.Error(t, err)
	assert.Equal(t, repositoryErrorMsg, err.Error())
}

func TestTrackThemeServiceListTrackThemesSuccess(t *testing.T) {
	trackID := "28712a55-04dd-4200-9316-4d6a1e399121"
	themeID := "6a4f86e4-4fef-4151-9c60-e467007dd213"
	trackTheme1, err := domain.NewTrackTheme(trackID, themeID, 0, 10, false)
	assert.NoError(t, err)

	trackThemes := []domain.TrackTheme{}
	trackThemes = append(trackThemes, trackTheme1)

	movie, err := domain.NewMovie("The Return of the King")
	assert.NoError(t, err)

	group, err := domain.NewGroup("Gondor", "Description of Gondor", "http://example.com/gondor.jpg")
	assert.NoError(t, err)

	track, err := domain.NewTrack("Track", trackID, nil)
	assert.NoError(t, err)

	theme, err := domain.NewTheme("Theme 1", themeID, "40929ca6-ed89-4548-a1d9-54b604ea50b2", "Description", 0, 1, nil)
	assert.NoError(t, err)

	trackThemeRepositoryMock := new(storagemocks.TrackThemeRepository)
	trackThemeRepositoryMock.On("FindByTrack", mock.Anything, mock.Anything).Return(trackThemes, nil).Once()
	defer trackThemeRepositoryMock.AssertExpectations(t)

	trackRepositoryMock := new(storagemocks.TrackRepository)
	trackRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(track, nil).Twice()
	defer trackRepositoryMock.AssertExpectations(t)

	movieRepositoryMock := new(storagemocks.MovieRepository)
	movieRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(movie, nil).Twice()
	defer movieRepositoryMock.AssertExpectations(t)

	themeRepositoryMock := new(storagemocks.ThemeRepository)
	themeRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(theme, nil).Once()
	defer themeRepositoryMock.AssertExpectations(t)

	groupRepositoryMock := new(storagemocks.GroupRepository)
	groupRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(group, nil).Once()
	defer groupRepositoryMock.AssertExpectations(t)

	categoryRepositoryMock := new(storagemocks.CategoryRepository)

	trackThemeService := NewTrackThemeService(trackThemeRepositoryMock, getGettingTrackServiceMock(trackRepositoryMock, movieRepositoryMock), getGettingThemeServiceMock(themeRepositoryMock, trackRepositoryMock, groupRepositoryMock, movieRepositoryMock, categoryRepositoryMock))

	themesDTO, err := trackThemeService.ListTracksThemesByTrack(context.Background(), trackID)
	assert.NoError(t, err)
	assert.Len(t, themesDTO, 1)
}

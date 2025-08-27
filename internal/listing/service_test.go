package listing

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

func TestUserServiceListUsersRepositoryError(t *testing.T) {
	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("FindAll", mock.Anything).Return(nil, errors.New(repositoryErrorMsg)).Once()

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

	groupService := NewGroupService(groupRepositoryMock)

	ctx := context.Background()
	_, err := groupService.ListGroups(ctx)
	assert.Error(t, err)
}

func TestGroupServiceListGroupsSuccess(t *testing.T) {
	groupRepositoryMock := new(storagemocks.GroupRepository)
	groups := []domain.Group{}
	group1, err := domain.NewGroup("Fellowship of the Ring")
	assert.NoError(t, err)
	groups = append(groups, group1)
	group2, err := domain.NewGroup("Company of the Ring")
	assert.NoError(t, err)
	groups = append(groups, group2)
	groupRepositoryMock.On("FindAll", mock.Anything).Return(groups, nil).Once()

	groupService := NewGroupService(groupRepositoryMock)

	groupsDTO, err := groupService.ListGroups(context.Background())
	assert.NoError(t, err)
	assert.Len(t, groupsDTO, 2)
	assert.Equal(t, "Fellowship of the Ring", groupsDTO[0].Name)
	assert.Equal(t, "Company of the Ring", groupsDTO[1].Name)
}

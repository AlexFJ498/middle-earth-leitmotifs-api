package listing

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
)

// UserService is the default implementation of the UserService interface
type UserService struct {
	userRepository domain.UserRepository
}

// NewUserService returns a new UserService instance.
func NewUserService(userRepository domain.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

// ListUsers implements the UserService interface for listing all users.
func (s UserService) ListUsers(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := s.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	userResponses := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		userResponses = append(userResponses, dto.NewUserResponse(
			user.ID().String(),
			user.Name().String(),
			user.Email().String(),
		))
	}

	return userResponses, nil
}

// MovieService is the service for managing movies.
type MovieService struct {
	movieRepository domain.MovieRepository
}

// NewMovieService returns a new MovieService instance.
func NewMovieService(movieRepository domain.MovieRepository) MovieService {
	return MovieService{
		movieRepository: movieRepository,
	}
}

// ListMovies implements the MovieService interface for listing all movies.
func (s MovieService) ListMovies(ctx context.Context) ([]dto.MovieResponse, error) {
	movies, err := s.movieRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	movieResponses := make([]dto.MovieResponse, 0, len(movies))
	for _, movie := range movies {
		movieResponses = append(movieResponses, dto.NewMovieResponse(
			movie.ID().String(),
			movie.Name().String(),
		))
	}

	return movieResponses, nil
}

// GroupService is the service for managing groups.
type GroupService struct {
	groupRepository domain.GroupRepository
}

// NewGroupService returns a new GroupService instance.
func NewGroupService(groupRepository domain.GroupRepository) GroupService {
	return GroupService{
		groupRepository: groupRepository,
	}
}

// ListGroups implements the GroupService interface for listing all groups.
func (s GroupService) ListGroups(ctx context.Context) ([]dto.GroupResponse, error) {
	groups, err := s.groupRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	groupResponses := make([]dto.GroupResponse, 0, len(groups))
	for _, group := range groups {
		groupResponses = append(groupResponses, dto.NewGroupResponse(
			group.ID().String(),
			group.Name().String(),
		))
	}

	return groupResponses, nil
}

// CategoryService is the service for managing categories.
type CategoryService struct {
	categoryRepository domain.CategoryRepository
}

// NewCategoryService returns a new CategoryService instance.
func NewCategoryService(categoryRepository domain.CategoryRepository) CategoryService {
	return CategoryService{
		categoryRepository: categoryRepository,
	}
}

// ListCategories implements the CategoryService interface for listing all categories.
func (s CategoryService) ListCategories(ctx context.Context) ([]dto.CategoryResponse, error) {
	categories, err := s.categoryRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	categoryResponses := make([]dto.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		categoryResponses = append(categoryResponses, dto.NewCategoryResponse(
			category.ID().String(),
			category.Name().String(),
		))
	}

	return categoryResponses, nil
}

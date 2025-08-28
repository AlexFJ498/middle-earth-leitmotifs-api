package creating

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/auth"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/event"
)

// UserService is the default implementation of the UserService interface
type UserService struct {
	userRepository domain.UserRepository
	eventBus       event.Bus
}

// NewUserService returns a new UserService instance.
func NewUserService(userRepository domain.UserRepository, eventBus event.Bus) UserService {
	return UserService{
		userRepository: userRepository,
		eventBus:       eventBus,
	}
}

// CreateUser implements the UserService interface for creating a new user.
func (s UserService) CreateUser(ctx context.Context, dto dto.UserCreateRequest) error {
	// Hash the user's password
	hashedPassword, err := auth.HashPassword(dto.Password)
	if err != nil {
		return err
	}

	// Create a new user object. The isAdmin field is always set to false.
	user, err := domain.NewUser(dto.Name, dto.Email, hashedPassword, false)
	if err != nil {
		return err
	}

	// Save the user to the database
	if err := s.userRepository.Save(ctx, user); err != nil {
		return err
	}

	// Publish user events
	return s.eventBus.Publish(ctx, user.PullEvents())
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

// CreateMovie implements the MovieService interface for creating a new movie.
func (s MovieService) CreateMovie(ctx context.Context, dto dto.MovieCreateRequest) error {
	movie, err := domain.NewMovie(dto.Name)
	if err != nil {
		return err
	}

	return s.movieRepository.Save(ctx, movie)
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

// CreateGroup implements the GroupService interface for creating a new group.
func (s GroupService) CreateGroup(ctx context.Context, dto dto.GroupCreateRequest) error {
	group, err := domain.NewGroup(dto.Name)
	if err != nil {
		return err
	}

	return s.groupRepository.Save(ctx, group)
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

// CreateCategory implements the CategoryService interface for creating a new category.
func (s CategoryService) CreateCategory(ctx context.Context, dto dto.CategoryCreateRequest) error {
	category, err := domain.NewCategory(dto.Name)
	if err != nil {
		return err
	}

	return s.categoryRepository.Save(ctx, category)
}

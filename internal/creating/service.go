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

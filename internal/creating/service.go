package creating

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/auth"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/event"
)

type UserService struct {
	userRepository domain.UserRepository
	eventBus       event.Bus
}

func NewUserService(userRepository domain.UserRepository, eventBus event.Bus) UserService {
	return UserService{
		userRepository: userRepository,
		eventBus:       eventBus,
	}
}

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

type MovieService struct {
	movieRepository domain.MovieRepository
}

func NewMovieService(movieRepository domain.MovieRepository) MovieService {
	return MovieService{
		movieRepository: movieRepository,
	}
}

func (s MovieService) CreateMovie(ctx context.Context, dto dto.MovieCreateRequest) error {
	movie, err := domain.NewMovie(dto.Name)
	if err != nil {
		return err
	}

	return s.movieRepository.Save(ctx, movie)
}

type GroupService struct {
	groupRepository domain.GroupRepository
}

func NewGroupService(groupRepository domain.GroupRepository) GroupService {
	return GroupService{
		groupRepository: groupRepository,
	}
}

func (s GroupService) CreateGroup(ctx context.Context, dto dto.GroupCreateRequest) error {
	group, err := domain.NewGroup(dto.Name, dto.Description, dto.ImageURL)
	if err != nil {
		return err
	}

	return s.groupRepository.Save(ctx, group)
}

type CategoryService struct {
	categoryRepository domain.CategoryRepository
}

func NewCategoryService(categoryRepository domain.CategoryRepository) CategoryService {
	return CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s CategoryService) CreateCategory(ctx context.Context, dto dto.CategoryCreateRequest) error {
	category, err := domain.NewCategory(dto.Name)
	if err != nil {
		return err
	}

	return s.categoryRepository.Save(ctx, category)
}

type TrackService struct {
	trackRepository domain.TrackRepository
}

func NewTrackService(trackRepository domain.TrackRepository) TrackService {
	return TrackService{
		trackRepository: trackRepository,
	}
}

func (s TrackService) CreateTrack(ctx context.Context, dto dto.TrackCreateRequest) error {
	track, err := domain.NewTrack(dto.Name, dto.MovieID, dto.SpotifyURL)
	if err != nil {
		return err
	}

	return s.trackRepository.Save(ctx, track)
}

type ThemeService struct {
	themeRepository domain.ThemeRepository
}

func NewThemeService(themeRepository domain.ThemeRepository) ThemeService {
	return ThemeService{
		themeRepository: themeRepository,
	}
}

func (s ThemeService) CreateTheme(ctx context.Context, dto dto.ThemeCreateRequest) error {
	theme, err := domain.NewTheme(dto.Name, dto.FirstHeard, dto.GroupID, dto.Description, dto.FirstHeardStart, dto.FirstHeardEnd, dto.CategoryID)
	if err != nil {
		return err
	}

	return s.themeRepository.Save(ctx, theme)
}

type TrackThemeService struct {
	trackThemeRepository domain.TrackThemeRepository
}

func NewTrackThemeService(trackThemeRepository domain.TrackThemeRepository) TrackThemeService {
	return TrackThemeService{
		trackThemeRepository: trackThemeRepository,
	}
}

func (s TrackThemeService) CreateTrackTheme(ctx context.Context, dto dto.TrackThemeCreateRequest) error {
	trackTheme, err := domain.NewTrackTheme(dto.TrackID, dto.ThemeID, dto.StartSecond, dto.EndSecond, dto.IsVariant)
	if err != nil {
		return err
	}

	return s.trackThemeRepository.Save(ctx, trackTheme)
}

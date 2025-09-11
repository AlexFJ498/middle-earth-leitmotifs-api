package getting

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
)

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

// GetMovie implements the MovieService interface for getting a movie by ID.
func (s MovieService) GetMovie(ctx context.Context, id string) (dto.MovieResponse, error) {
	movieID, err := domain.NewMovieIDFromString(id)
	if err != nil {
		return dto.MovieResponse{}, err
	}

	movie, err := s.movieRepository.Find(ctx, movieID)
	if err != nil {
		return dto.MovieResponse{}, err
	}

	return dto.NewMovieResponse(movie), nil
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

// GetGroup implements the GroupService interface for getting a group by ID.
func (s GroupService) GetGroup(ctx context.Context, id string) (dto.GroupResponse, error) {
	groupID, err := domain.NewGroupIDFromString(id)
	if err != nil {
		return dto.GroupResponse{}, err
	}

	group, err := s.groupRepository.Find(ctx, groupID)
	if err != nil {
		return dto.GroupResponse{}, err
	}

	return dto.NewGroupResponse(group), nil
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

// GetCategory implements the CategoryService interface for getting a category by ID.
func (s CategoryService) GetCategory(ctx context.Context, id string) (dto.CategoryResponse, error) {
	categoryID, err := domain.NewCategoryIDFromString(id)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	category, err := s.categoryRepository.Find(ctx, categoryID)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.NewCategoryResponse(category), nil
}

// TrackService is the service for managing tracks.
type TrackService struct {
	trackRepository domain.TrackRepository
	MovieService    MovieService
}

// NewTrackService returns a new TrackService instance.
func NewTrackService(trackRepository domain.TrackRepository, movieService MovieService) TrackService {
	return TrackService{
		trackRepository: trackRepository,
		MovieService:    movieService,
	}
}

// GetTrack implements the TrackService interface for getting a track by ID.
func (s TrackService) GetTrack(ctx context.Context, id string) (dto.TrackResponse, error) {
	trackID, err := domain.NewTrackIDFromString(id)
	if err != nil {
		return dto.TrackResponse{}, err
	}

	track, err := s.trackRepository.Find(ctx, trackID)
	if err != nil {
		return dto.TrackResponse{}, err
	}

	// Obtain the movie related to the track.
	movieDTO, err := s.MovieService.GetMovie(ctx, track.MovieID().String())
	if err != nil {
		return dto.TrackResponse{}, err
	}

	return dto.NewTrackResponse(track, movieDTO), nil
}

// ThemeService is the service for managing themes.
type ThemeService struct {
	themeRepository domain.ThemeRepository
	trackService    TrackService
	groupService    GroupService
	categoryService CategoryService
}

// NewThemeService returns a new ThemeService instance.
func NewThemeService(themeRepository domain.ThemeRepository, trackService TrackService, groupService GroupService, categoryService CategoryService) ThemeService {
	return ThemeService{
		themeRepository: themeRepository,
		trackService:    trackService,
		groupService:    groupService,
		categoryService: categoryService,
	}
}

// GetTheme implements the ThemeService interface for getting a theme by ID.
func (s ThemeService) GetTheme(ctx context.Context, id string) (dto.ThemeResponse, error) {
	themeID, err := domain.NewThemeIDFromString(id)
	if err != nil {
		return dto.ThemeResponse{}, err
	}

	theme, err := s.themeRepository.Find(ctx, themeID)
	if err != nil {
		return dto.ThemeResponse{}, err
	}

	// Fetch related entities.
	groupDTO, err := s.groupService.GetGroup(ctx, theme.GroupID().String())
	if err != nil {
		return dto.ThemeResponse{}, err
	}

	trackDTO, err := s.trackService.GetTrack(ctx, theme.FirstHeard().String())
	if err != nil {
		return dto.ThemeResponse{}, err
	}

	var categoryDTO *dto.CategoryResponse
	if theme.CategoryID() != nil {
		categoryDTORes, err := s.categoryService.GetCategory(ctx, theme.CategoryID().String())
		if err != nil {
			return dto.ThemeResponse{}, err
		}
		categoryDTO = &categoryDTORes
	}

	return dto.NewThemeResponse(theme, trackDTO, groupDTO, categoryDTO), nil
}

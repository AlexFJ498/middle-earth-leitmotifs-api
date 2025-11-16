package getting

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
)

type MovieService struct {
	movieRepository domain.MovieRepository
}

func NewMovieService(movieRepository domain.MovieRepository) MovieService {
	return MovieService{
		movieRepository: movieRepository,
	}
}

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

type GroupService struct {
	groupRepository domain.GroupRepository
}

func NewGroupService(groupRepository domain.GroupRepository) GroupService {
	return GroupService{
		groupRepository: groupRepository,
	}
}

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

type CategoryService struct {
	categoryRepository domain.CategoryRepository
}

func NewCategoryService(categoryRepository domain.CategoryRepository) CategoryService {
	return CategoryService{
		categoryRepository: categoryRepository,
	}
}

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

type TrackService struct {
	trackRepository domain.TrackRepository
	MovieService    MovieService
}

func NewTrackService(trackRepository domain.TrackRepository, movieService MovieService) TrackService {
	return TrackService{
		trackRepository: trackRepository,
		MovieService:    movieService,
	}
}

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

type ThemeService struct {
	themeRepository domain.ThemeRepository
	trackService    TrackService
	groupService    GroupService
	categoryService CategoryService
}

func NewThemeService(themeRepository domain.ThemeRepository, trackService TrackService, groupService GroupService, categoryService CategoryService) ThemeService {
	return ThemeService{
		themeRepository: themeRepository,
		trackService:    trackService,
		groupService:    groupService,
		categoryService: categoryService,
	}
}

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

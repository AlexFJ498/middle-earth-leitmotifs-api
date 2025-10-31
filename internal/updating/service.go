package updating

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

func (s *MovieService) UpdateMovie(ctx context.Context, id string, dto dto.MovieUpdateRequest) error {
	movie, err := domain.NewMovieWithID(id, dto.Name)
	if err != nil {
		return err
	}
	return s.movieRepository.Update(ctx, movie)
}

type GroupService struct {
	groupRepository domain.GroupRepository
}

func NewGroupService(groupRepository domain.GroupRepository) GroupService {
	return GroupService{
		groupRepository: groupRepository,
	}
}

func (s *GroupService) UpdateGroup(ctx context.Context, id string, dto dto.GroupUpdateRequest) error {
	group, err := domain.NewGroupWithID(id, dto.Name, dto.Description, dto.ImageURL)
	if err != nil {
		return err
	}
	return s.groupRepository.Update(ctx, group)
}

type CategoryService struct {
	categoryRepository domain.CategoryRepository
}

func NewCategoryService(categoryRepository domain.CategoryRepository) CategoryService {
	return CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryService) UpdateCategory(ctx context.Context, id string, dto dto.CategoryUpdateRequest) error {
	category, err := domain.NewCategoryWithID(id, dto.Name)
	if err != nil {
		return err
	}
	return s.categoryRepository.Update(ctx, category)
}

type TrackService struct {
	trackRepository domain.TrackRepository
}

func NewTrackService(trackRepository domain.TrackRepository) TrackService {
	return TrackService{
		trackRepository: trackRepository,
	}
}

func (s *TrackService) UpdateTrack(ctx context.Context, id string, dto dto.TrackUpdateRequest) error {
	track, err := domain.NewTrackWithID(id, dto.Name, dto.MovieID, dto.SpotifyURL)
	if err != nil {
		return err
	}
	return s.trackRepository.Update(ctx, track)
}

type ThemeService struct {
	themeRepository domain.ThemeRepository
}

func NewThemeService(themeRepository domain.ThemeRepository) ThemeService {
	return ThemeService{
		themeRepository: themeRepository,
	}
}

func (s *ThemeService) UpdateTheme(ctx context.Context, id string, dto dto.ThemeUpdateRequest) error {
	theme, err := domain.NewThemeWithID(id, dto.Name, dto.FirstHeard, dto.GroupID, dto.Description, dto.FirstHeardStart, dto.FirstHeardEnd, dto.CategoryID)
	if err != nil {
		return err
	}
	return s.themeRepository.Update(ctx, theme)
}

type TrackThemeService struct {
	trackThemeRepository domain.TrackThemeRepository
}

func NewTrackThemeService(trackThemeRepository domain.TrackThemeRepository) TrackThemeService {
	return TrackThemeService{
		trackThemeRepository: trackThemeRepository,
	}
}

func (s *TrackThemeService) UpdateTrackTheme(ctx context.Context, dto dto.TrackThemeUpdateRequest) error {
	trackTheme, err := domain.NewTrackTheme(dto.TrackID, dto.ThemeID, dto.StartSecond, dto.EndSecond, dto.IsVariant)
	if err != nil {
		return err
	}
	return s.trackThemeRepository.Update(ctx, trackTheme)
}

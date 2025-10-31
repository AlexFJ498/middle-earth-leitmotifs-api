package deleting

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
)

type MovieService struct {
	movieRepository domain.MovieRepository
}

func NewMovieService(repo domain.MovieRepository) MovieService {
	return MovieService{
		movieRepository: repo,
	}
}

func (s *MovieService) DeleteMovie(ctx context.Context, id domain.MovieID) error {
	return s.movieRepository.Delete(ctx, id)
}

type GroupService struct {
	groupRepository domain.GroupRepository
}

func NewGroupService(repo domain.GroupRepository) GroupService {
	return GroupService{
		groupRepository: repo,
	}
}

func (s *GroupService) DeleteGroup(ctx context.Context, id domain.GroupID) error {
	return s.groupRepository.Delete(ctx, id)
}

type CategoryService struct {
	categoryRepository domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) CategoryService {
	return CategoryService{
		categoryRepository: repo,
	}
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id domain.CategoryID) error {
	return s.categoryRepository.Delete(ctx, id)
}

type TrackService struct {
	trackRepository domain.TrackRepository
}

func NewTrackService(repo domain.TrackRepository) TrackService {
	return TrackService{
		trackRepository: repo,
	}
}

func (s *TrackService) DeleteTrack(ctx context.Context, id domain.TrackID) error {
	return s.trackRepository.Delete(ctx, id)
}

type ThemeService struct {
	themeRepository domain.ThemeRepository
}

func NewThemeService(repo domain.ThemeRepository) ThemeService {
	return ThemeService{
		themeRepository: repo,
	}
}

func (s *ThemeService) DeleteTheme(ctx context.Context, id domain.ThemeID) error {
	return s.themeRepository.Delete(ctx, id)
}

type TrackThemeService struct {
	trackThemeRepository domain.TrackThemeRepository
}

func NewTrackThemeService(repo domain.TrackThemeRepository) TrackThemeService {
	return TrackThemeService{
		trackThemeRepository: repo,
	}
}

func (s *TrackThemeService) DeleteTrackTheme(ctx context.Context, trackID domain.TrackID, themeID domain.ThemeID, startSecond domain.StartSecond) error {
	return s.trackThemeRepository.Delete(ctx, trackID, themeID, startSecond)
}

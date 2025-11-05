package listing

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/getting"
)

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (s UserService) ListUsers(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := s.userRepository.FindAll(ctx)
	if err != nil {
		return []dto.UserResponse{}, err
	}

	userResponses := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		userResponses = append(userResponses, dto.NewUserResponse(user))
	}

	return userResponses, nil
}

type MovieService struct {
	movieRepository domain.MovieRepository
}

func NewMovieService(movieRepository domain.MovieRepository) MovieService {
	return MovieService{
		movieRepository: movieRepository,
	}
}

func (s MovieService) ListMovies(ctx context.Context) ([]dto.MovieResponse, error) {
	movies, err := s.movieRepository.FindAll(ctx)
	if err != nil {
		return []dto.MovieResponse{}, err
	}

	movieResponses := make([]dto.MovieResponse, 0, len(movies))
	for _, movie := range movies {
		movieResponses = append(movieResponses, dto.NewMovieResponse(movie))
	}

	return movieResponses, nil
}

type GroupService struct {
	groupRepository domain.GroupRepository
}

func NewGroupService(groupRepository domain.GroupRepository) GroupService {
	return GroupService{
		groupRepository: groupRepository,
	}
}

func (s GroupService) ListGroups(ctx context.Context) ([]dto.GroupResponse, error) {
	groups, err := s.groupRepository.FindAll(ctx)
	if err != nil {
		return []dto.GroupResponse{}, err
	}

	groupResponses := make([]dto.GroupResponse, 0, len(groups))
	for _, group := range groups {
		groupResponses = append(groupResponses, dto.NewGroupResponse(group))
	}

	return groupResponses, nil
}

type CategoryService struct {
	categoryRepository domain.CategoryRepository
}

func NewCategoryService(categoryRepository domain.CategoryRepository) CategoryService {
	return CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s CategoryService) ListCategories(ctx context.Context) ([]dto.CategoryResponse, error) {
	categories, err := s.categoryRepository.FindAll(ctx)
	if err != nil {
		return []dto.CategoryResponse{}, err
	}

	categoryResponses := make([]dto.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		categoryResponses = append(categoryResponses, dto.NewCategoryResponse(category))
	}

	return categoryResponses, nil
}

type TrackService struct {
	trackRepository     domain.TrackRepository
	MovieService        MovieService
	GettingMovieService getting.MovieService
}

func NewTrackService(trackRepository domain.TrackRepository, movieService MovieService, gettingMovieService getting.MovieService) TrackService {
	return TrackService{
		trackRepository:     trackRepository,
		MovieService:        movieService,
		GettingMovieService: gettingMovieService,
	}
}

func (s TrackService) ListTracks(ctx context.Context) ([]dto.TrackResponse, error) {
	tracks, err := s.trackRepository.FindAll(ctx)
	if err != nil {
		return []dto.TrackResponse{}, err
	}

	trackResponses := make([]dto.TrackResponse, 0, len(tracks))
	for _, track := range tracks {
		// Obtain the movie related to the track.
		movieDTO, err := s.GettingMovieService.GetMovie(ctx, track.MovieID().String())
		if err != nil {
			return []dto.TrackResponse{}, err
		}

		trackResponses = append(trackResponses, dto.NewTrackResponse(track, movieDTO))
	}

	return trackResponses, nil
}

func (s TrackService) ListTracksByMovie(ctx context.Context, movieID string) ([]dto.TrackResponse, error) {
	movieIDObj, err := domain.NewMovieIDFromString(movieID)
	if err != nil {
		return nil, err
	}

	tracks, err := s.trackRepository.FindByMovie(ctx, movieIDObj)
	if err != nil {
		return []dto.TrackResponse{}, err
	}

	trackResponses := make([]dto.TrackResponse, 0, len(tracks))
	for _, track := range tracks {
		movieDTO, err := s.GettingMovieService.GetMovie(ctx, track.MovieID().String())
		if err != nil {
			return []dto.TrackResponse{}, err
		}

		trackResponses = append(trackResponses, dto.NewTrackResponse(track, movieDTO))
	}

	return trackResponses, nil
}

type ThemeService struct {
	themeRepository        domain.ThemeRepository
	trackService           TrackService
	groupService           GroupService
	categoryService        CategoryService
	GettingGroupService    getting.GroupService
	GettingTrackService    getting.TrackService
	GettingCategoryService getting.CategoryService
}

func NewThemeService(themeRepository domain.ThemeRepository, trackService TrackService, groupService GroupService, categoryService CategoryService, gettingGroupService getting.GroupService, gettingTrackService getting.TrackService, gettingCategoryService getting.CategoryService) ThemeService {
	return ThemeService{
		themeRepository:        themeRepository,
		trackService:           trackService,
		groupService:           groupService,
		categoryService:        categoryService,
		GettingGroupService:    gettingGroupService,
		GettingTrackService:    gettingTrackService,
		GettingCategoryService: gettingCategoryService,
	}
}

func (s ThemeService) ListThemes(ctx context.Context) ([]dto.ThemeResponse, error) {
	themes, err := s.themeRepository.FindAll(ctx)
	if err != nil {
		return []dto.ThemeResponse{}, err
	}

	themeResponses := make([]dto.ThemeResponse, 0, len(themes))
	for _, theme := range themes {
		groupDTO, err := s.GettingGroupService.GetGroup(ctx, theme.GroupID().String())
		if err != nil {
			return nil, err
		}

		trackDTO, err := s.GettingTrackService.GetTrack(ctx, theme.FirstHeard().String())
		if err != nil {
			return nil, err
		}

		var categoryDTO *dto.CategoryResponse
		if theme.CategoryID() != nil {
			categoryDTORes, err := s.GettingCategoryService.GetCategory(ctx, theme.CategoryID().String())
			if err != nil {
				return nil, err
			}
			categoryDTO = &categoryDTORes
		}

		themeResponses = append(themeResponses, dto.NewThemeResponse(theme, trackDTO, groupDTO, categoryDTO))
	}

	return themeResponses, nil
}

func (s ThemeService) ListThemesByGroup(ctx context.Context, groupID string) ([]dto.ThemeResponse, error) {
	groupIDObj, err := domain.NewGroupIDFromString(groupID)
	if err != nil {
		return nil, err
	}

	themes, err := s.themeRepository.FindByGroup(ctx, groupIDObj)
	if err != nil {
		return []dto.ThemeResponse{}, err
	}

	themeResponses := make([]dto.ThemeResponse, 0, len(themes))
	for _, theme := range themes {
		groupDTO, err := s.GettingGroupService.GetGroup(ctx, theme.GroupID().String())
		if err != nil {
			return nil, err
		}

		trackDTO, err := s.GettingTrackService.GetTrack(ctx, theme.FirstHeard().String())
		if err != nil {
			return nil, err
		}

		var categoryDTO *dto.CategoryResponse
		if theme.CategoryID() != nil {
			categoryDTORes, err := s.GettingCategoryService.GetCategory(ctx, theme.CategoryID().String())
			if err != nil {
				return nil, err
			}
			categoryDTO = &categoryDTORes
		}

		themeResponses = append(themeResponses, dto.NewThemeResponse(theme, trackDTO, groupDTO, categoryDTO))
	}

	return themeResponses, nil
}

type TrackThemeService struct {
	trackThemeRepository domain.TrackThemeRepository
	GettingTrackService  getting.TrackService
	GettingThemeService  getting.ThemeService
}

func NewTrackThemeService(trackThemeRepository domain.TrackThemeRepository, gettingTrackService getting.TrackService, gettingThemeService getting.ThemeService) TrackThemeService {
	return TrackThemeService{
		trackThemeRepository: trackThemeRepository,
		GettingTrackService:  gettingTrackService,
		GettingThemeService:  gettingThemeService,
	}
}

func (s TrackThemeService) ListTracksThemesByTrack(ctx context.Context, trackID string) ([]dto.TrackThemeResponse, error) {
	trackIDObj, err := domain.NewTrackIDFromString(trackID)
	if err != nil {
		return nil, err
	}

	trackThemes, err := s.trackThemeRepository.FindByTrack(ctx, trackIDObj)
	if err != nil {
		return []dto.TrackThemeResponse{}, err
	}

	trackThemeResponses := make([]dto.TrackThemeResponse, 0, len(trackThemes))
	for _, trackTheme := range trackThemes {
		trackDTO, err := s.GettingTrackService.GetTrack(ctx, trackTheme.TrackID().String())
		if err != nil {
			return nil, err
		}

		themeDTO, err := s.GettingThemeService.GetTheme(ctx, trackTheme.ThemeID().String())
		if err != nil {
			return nil, err
		}

		trackThemeResponses = append(trackThemeResponses, dto.NewTrackThemeResponse(trackTheme, trackDTO, themeDTO))
	}
	return trackThemeResponses, nil
}

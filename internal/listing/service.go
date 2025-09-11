package listing

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/getting"
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
		return []dto.UserResponse{}, err
	}

	userResponses := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		userResponses = append(userResponses, dto.NewUserResponse(user))
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
		return []dto.MovieResponse{}, err
	}

	movieResponses := make([]dto.MovieResponse, 0, len(movies))
	for _, movie := range movies {
		movieResponses = append(movieResponses, dto.NewMovieResponse(movie))
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
		return []dto.GroupResponse{}, err
	}

	groupResponses := make([]dto.GroupResponse, 0, len(groups))
	for _, group := range groups {
		groupResponses = append(groupResponses, dto.NewGroupResponse(group))
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
		return []dto.CategoryResponse{}, err
	}

	categoryResponses := make([]dto.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		categoryResponses = append(categoryResponses, dto.NewCategoryResponse(category))
	}

	return categoryResponses, nil
}

// TrackService is the service for managing tracks.
type TrackService struct {
	trackRepository     domain.TrackRepository
	MovieService        MovieService
	GettingMovieService getting.MovieService
}

// NewTrackService returns a new TrackService instance.
func NewTrackService(trackRepository domain.TrackRepository, movieService MovieService, gettingMovieService getting.MovieService) TrackService {
	return TrackService{
		trackRepository:     trackRepository,
		MovieService:        movieService,
		GettingMovieService: gettingMovieService,
	}
}

// ListTracks implements the TrackService interface for listing all tracks.
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

// ThemeService is the service for managing themes.
type ThemeService struct {
	themeRepository        domain.ThemeRepository
	trackService           TrackService
	groupService           GroupService
	categoryService        CategoryService
	GettingGroupService    getting.GroupService
	GettingTrackService    getting.TrackService
	GettingCategoryService getting.CategoryService
}

// NewThemeService returns a new ThemeService instance.
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

// ListThemes implements the ThemeService interface for listing all themes.
func (s ThemeService) ListThemes(ctx context.Context) ([]dto.ThemeResponse, error) {
	themes, err := s.themeRepository.FindAll(ctx)
	if err != nil {
		return []dto.ThemeResponse{}, err
	}

	themeResponses := make([]dto.ThemeResponse, 0, len(themes))
	for _, theme := range themes {
		// Fetch related entities.

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

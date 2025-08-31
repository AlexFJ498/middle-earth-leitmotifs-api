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

// ListTracks implements the TrackService interface for listing all tracks.
func (s TrackService) ListTracks(ctx context.Context) ([]dto.TrackResponse, error) {
	tracks, err := s.trackRepository.FindAll(ctx)
	if err != nil {
		return []dto.TrackResponse{}, err
	}

	trackResponses := make([]dto.TrackResponse, 0, len(tracks))
	for _, track := range tracks {
		// Obtain the movie related to the track.
		movieDTO, err := s.MovieService.GetMovie(ctx, track.MovieID().String())
		if err != nil {
			return []dto.TrackResponse{}, err
		}

		trackResponses = append(trackResponses, dto.NewTrackResponse(track, movieDTO))
	}

	return trackResponses, nil
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

// ListThemes implements the ThemeService interface for listing all themes.
func (s ThemeService) ListThemes(ctx context.Context) ([]dto.ThemeResponse, error) {
	themes, err := s.themeRepository.FindAll(ctx)
	if err != nil {
		return []dto.ThemeResponse{}, err
	}

	themeResponses := make([]dto.ThemeResponse, 0, len(themes))
	for _, theme := range themes {
		// Fetch related entities.

		groupDTO, err := s.groupService.GetGroup(ctx, theme.GroupID().String())
		if err != nil {
			return nil, err
		}

		trackDTO, err := s.trackService.GetTrack(ctx, theme.FirstHeard().String())
		if err != nil {
			return nil, err
		}

		var categoryDTO *dto.CategoryResponse
		if theme.CategoryID() != nil {
			categoryDTORes, err := s.categoryService.GetCategory(ctx, theme.CategoryID().String())
			if err != nil {
				return nil, err
			}
			categoryDTO = &categoryDTORes
		}

		themeResponses = append(themeResponses, dto.NewThemeResponse(theme, trackDTO, groupDTO, categoryDTO))
	}

	return themeResponses, nil
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

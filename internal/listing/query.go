package listing

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
)

const (
	UsersQueryType               = "query.listing.users"
	MoviesQueryType              = "query.listing.movies"
	GroupsQueryType              = "query.listing.groups"
	CategoriesQueryType          = "query.listing.categories"
	TracksQueryType              = "query.listing.tracks"
	TracksByMovieQueryType       = "query.listing.tracks.by_movie"
	ThemesQueryType              = "query.listing.themes"
	ThemesByGroupQueryType       = "query.listing.themes.by_group"
	TracksThemesByTrackQueryType = "query.listing.track_themes.by_track"
)

type UsersQuery struct{}

func NewUsersQuery() UsersQuery {
	return UsersQuery{}
}

func (q UsersQuery) Type() query.Type {
	return UsersQueryType
}

type UsersQueryHandler struct {
	userService UserService
}

func NewUsersQueryHandler(userService UserService) UsersQueryHandler {
	return UsersQueryHandler{
		userService: userService,
	}
}

func (h UsersQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	_, ok := query.(UsersQuery)
	if !ok {
		return nil, nil
	}

	return h.userService.ListUsers(ctx)
}

type MoviesQuery struct{}

func NewMoviesQuery() MoviesQuery {
	return MoviesQuery{}
}

func (q MoviesQuery) Type() query.Type {
	return MoviesQueryType
}

type MoviesQueryHandler struct {
	movieService MovieService
}

func NewMoviesQueryHandler(movieService MovieService) MoviesQueryHandler {
	return MoviesQueryHandler{
		movieService: movieService,
	}
}

func (h MoviesQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	_, ok := query.(MoviesQuery)
	if !ok {
		return nil, nil
	}

	return h.movieService.ListMovies(ctx)
}

type GroupsQuery struct{}

func NewGroupsQuery() GroupsQuery {
	return GroupsQuery{}
}

func (q GroupsQuery) Type() query.Type {
	return GroupsQueryType
}

type GroupsQueryHandler struct {
	groupService GroupService
}

func NewGroupsQueryHandler(groupService GroupService) GroupsQueryHandler {
	return GroupsQueryHandler{
		groupService: groupService,
	}
}

func (h GroupsQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	_, ok := query.(GroupsQuery)
	if !ok {
		return nil, nil
	}

	return h.groupService.ListGroups(ctx)
}

type CategoriesQuery struct{}

func NewCategoriesQuery() CategoriesQuery {
	return CategoriesQuery{}
}

func (q CategoriesQuery) Type() query.Type {
	return CategoriesQueryType
}

type CategoriesQueryHandler struct {
	categoryService CategoryService
}

func NewCategoriesQueryHandler(categoryService CategoryService) CategoriesQueryHandler {
	return CategoriesQueryHandler{
		categoryService: categoryService,
	}
}

func (h CategoriesQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	_, ok := query.(CategoriesQuery)
	if !ok {
		return nil, nil
	}

	return h.categoryService.ListCategories(ctx)
}

type TracksQuery struct{}

func NewTracksQuery() TracksQuery {
	return TracksQuery{}
}

func (q TracksQuery) Type() query.Type {
	return TracksQueryType
}

type TracksQueryHandler struct {
	trackService TrackService
}

func NewTracksQueryHandler(trackService TrackService) TracksQueryHandler {
	return TracksQueryHandler{
		trackService: trackService,
	}
}

func (h TracksQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	_, ok := query.(TracksQuery)
	if !ok {
		return nil, nil
	}

	return h.trackService.ListTracks(ctx)
}

type TracksByMovieQuery struct {
	MovieID string
}

func NewTracksByMovieQuery(movieID string) TracksByMovieQuery {
	return TracksByMovieQuery{
		MovieID: movieID,
	}
}

func (q TracksByMovieQuery) Type() query.Type {
	return TracksByMovieQueryType
}

type TracksByMovieQueryHandler struct {
	trackService TrackService
}

func NewTracksByMovieQueryHandler(trackService TrackService) TracksByMovieQueryHandler {
	return TracksByMovieQueryHandler{
		trackService: trackService,
	}
}

func (h TracksByMovieQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	q, ok := query.(TracksByMovieQuery)
	if !ok {
		return nil, nil
	}

	return h.trackService.ListTracksByMovie(ctx, q.MovieID)
}

type ThemesQuery struct{}

func NewThemesQuery() ThemesQuery {
	return ThemesQuery{}
}

func (q ThemesQuery) Type() query.Type {
	return ThemesQueryType
}

type ThemesQueryHandler struct {
	themeService ThemeService
}

func NewThemesQueryHandler(themeService ThemeService) ThemesQueryHandler {
	return ThemesQueryHandler{
		themeService: themeService,
	}
}

func (h ThemesQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	_, ok := query.(ThemesQuery)
	if !ok {
		return nil, nil
	}

	return h.themeService.ListThemes(ctx)
}

type ThemesByGroupQuery struct {
	GroupID string
}

func NewThemesByGroupQuery(groupID string) ThemesByGroupQuery {
	return ThemesByGroupQuery{
		GroupID: groupID,
	}
}

func (q ThemesByGroupQuery) Type() query.Type {
	return ThemesByGroupQueryType
}

type ThemesByGroupQueryHandler struct {
	themeService ThemeService
}

func NewThemesByGroupQueryHandler(themeService ThemeService) ThemesByGroupQueryHandler {
	return ThemesByGroupQueryHandler{
		themeService: themeService,
	}
}

func (h ThemesByGroupQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	q, ok := query.(ThemesByGroupQuery)
	if !ok {
		return nil, nil
	}

	return h.themeService.ListThemesByGroup(ctx, q.GroupID)
}

type TracksThemesByTrackQuery struct {
	TrackID string
}

func NewTracksThemesByTrackQuery(trackID string) TracksThemesByTrackQuery {
	return TracksThemesByTrackQuery{
		TrackID: trackID,
	}
}

func (q TracksThemesByTrackQuery) Type() query.Type {
	return TracksThemesByTrackQueryType
}

type TracksThemesByTrackQueryHandler struct {
	trackThemeService TrackThemeService
}

func NewTracksThemesByTrackQueryHandler(trackThemeService TrackThemeService) TracksThemesByTrackQueryHandler {
	return TracksThemesByTrackQueryHandler{
		trackThemeService: trackThemeService,
	}
}

func (h TracksThemesByTrackQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	q, ok := query.(TracksThemesByTrackQuery)
	if !ok {
		return nil, nil
	}

	return h.trackThemeService.ListTracksThemesByTrack(ctx, q.TrackID)
}

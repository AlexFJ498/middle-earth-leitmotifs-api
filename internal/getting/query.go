package getting

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
)

const (
	MoviesQueryType     = "query.getting.movies"
	GroupsQueryType     = "query.getting.groups"
	CategoriesQueryType = "query.getting.categories"
	TracksQueryType     = "query.getting.tracks"
	ThemesQueryType     = "query.getting.themes"
)

// MoviesQuery is the query for getting a movie.
type MoviesQuery struct {
	ID string
}

// NewMoviesQuery creates a new MoviesQuery instance.
func NewMoviesQuery(id string) MoviesQuery {
	return MoviesQuery{
		ID: id,
	}
}

// Type returns the query type.
func (q MoviesQuery) Type() query.Type {
	return MoviesQueryType
}

// MoviesQueryHandler handles the movies query.
type MoviesQueryHandler struct {
	movieService MovieService
}

// NewMoviesQueryHandler creates a new MoviesQueryHandler instance.
func NewMoviesQueryHandler(movieService MovieService) MoviesQueryHandler {
	return MoviesQueryHandler{
		movieService: movieService,
	}
}

// Handle handles the movies query.
func (h MoviesQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	movieQuery, ok := query.(MoviesQuery)
	if !ok {
		return nil, nil
	}

	return h.movieService.GetMovie(ctx, movieQuery.ID)
}

// GroupsQuery is the query for getting a group.
type GroupsQuery struct {
	ID string
}

// NewGroupsQuery creates a new GroupsQuery instance.
func NewGroupsQuery(id string) GroupsQuery {
	return GroupsQuery{
		ID: id,
	}
}

// Type returns the query type.
func (q GroupsQuery) Type() query.Type {
	return GroupsQueryType
}

// GroupsQueryHandler handles the groups query.
type GroupsQueryHandler struct {
	groupService GroupService
}

// NewGroupsQueryHandler creates a new GroupsQueryHandler instance.
func NewGroupsQueryHandler(groupService GroupService) GroupsQueryHandler {
	return GroupsQueryHandler{
		groupService: groupService,
	}
}

// Handle handles the groups query.
func (h GroupsQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	groupQuery, ok := query.(GroupsQuery)
	if !ok {
		return nil, nil
	}

	return h.groupService.GetGroup(ctx, groupQuery.ID)
}

// CategoriesQuery is the query for getting a category.
type CategoriesQuery struct {
	ID string
}

// NewCategoriesQuery creates a new CategoriesQuery instance.
func NewCategoriesQuery(id string) CategoriesQuery {
	return CategoriesQuery{
		ID: id,
	}
}

// Type returns the query type.
func (q CategoriesQuery) Type() query.Type {
	return CategoriesQueryType
}

// CategoriesQueryHandler handles the categories query.
type CategoriesQueryHandler struct {
	categoryService CategoryService
}

// NewCategoriesQueryHandler creates a new CategoriesQueryHandler instance.
func NewCategoriesQueryHandler(categoryService CategoryService) CategoriesQueryHandler {
	return CategoriesQueryHandler{
		categoryService: categoryService,
	}
}

// Handle handles the categories query.
func (h CategoriesQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	categoryQuery, ok := query.(CategoriesQuery)
	if !ok {
		return nil, nil
	}

	return h.categoryService.GetCategory(ctx, categoryQuery.ID)
}

// TracksQuery is the query for getting a track.
type TracksQuery struct {
	ID string
}

// NewTracksQuery creates a new TracksQuery instance.
func NewTracksQuery(id string) TracksQuery {
	return TracksQuery{
		ID: id,
	}
}

// Type returns the query type.
func (q TracksQuery) Type() query.Type {
	return TracksQueryType
}

// TracksQueryHandler handles the tracks query.
type TracksQueryHandler struct {
	trackService TrackService
}

// NewTracksQueryHandler creates a new TracksQueryHandler instance.
func NewTracksQueryHandler(trackService TrackService) TracksQueryHandler {
	return TracksQueryHandler{
		trackService: trackService,
	}
}

// Handle handles the tracks query.
func (h TracksQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	trackQuery, ok := query.(TracksQuery)
	if !ok {
		return nil, nil
	}

	return h.trackService.GetTrack(ctx, trackQuery.ID)
}

// ThemesQuery is the query for getting a theme.
type ThemesQuery struct {
	ID string
}

// NewThemesQuery creates a new ThemesQuery instance.
func NewThemesQuery(id string) ThemesQuery {
	return ThemesQuery{
		ID: id,
	}
}

// Type returns the query type.
func (q ThemesQuery) Type() query.Type {
	return ThemesQueryType
}

// ThemesQueryHandler handles the themes query.
type ThemesQueryHandler struct {
	themeService ThemeService
}

// NewThemesQueryHandler creates a new ThemesQueryHandler instance.
func NewThemesQueryHandler(themeService ThemeService) ThemesQueryHandler {
	return ThemesQueryHandler{
		themeService: themeService,
	}
}

// Handle handles the themes query.
func (h ThemesQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	themeQuery, ok := query.(ThemesQuery)
	if !ok {
		return nil, nil
	}

	return h.themeService.GetTheme(ctx, themeQuery.ID)
}

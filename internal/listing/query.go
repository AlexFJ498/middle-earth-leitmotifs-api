package listing

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
)

const (
	UsersQueryType         = "query.listing.users"
	MoviesQueryType        = "query.listing.movies"
	GroupsQueryType        = "query.listing.groups"
	CategoriesQueryType    = "query.listing.categories"
	TracksQueryType        = "query.listing.tracks"
	ThemesQueryType        = "query.listing.themes"
	ThemesByGroupQueryType = "query.listing.themes.by_group"
)

// UsersQuery represents a query for listing all users.
type UsersQuery struct{}

// NewUsersQuery creates a new UsersQuery instance.
func NewUsersQuery() UsersQuery {
	return UsersQuery{}
}

// Type returns the query type.
func (q UsersQuery) Type() query.Type {
	return UsersQueryType
}

// UsersQueryHandler handles the users query.
type UsersQueryHandler struct {
	userService UserService
}

// NewUsersQueryHandler creates a new UsersQueryHandler instance.
func NewUsersQueryHandler(userService UserService) UsersQueryHandler {
	return UsersQueryHandler{
		userService: userService,
	}
}

// Handle handles the users query.
func (h UsersQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	_, ok := query.(UsersQuery)
	if !ok {
		return nil, nil
	}

	return h.userService.ListUsers(ctx)
}

// MoviesQuery represents a query for listing all movies.
type MoviesQuery struct{}

// NewMoviesQuery creates a new MoviesQuery instance.
func NewMoviesQuery() MoviesQuery {
	return MoviesQuery{}
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
	_, ok := query.(MoviesQuery)
	if !ok {
		return nil, nil
	}

	return h.movieService.ListMovies(ctx)
}

// GroupsQuery represents a query for listing all groups.
type GroupsQuery struct{}

// NewGroupsQuery creates a new GroupsQuery instance.
func NewGroupsQuery() GroupsQuery {
	return GroupsQuery{}
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
	_, ok := query.(GroupsQuery)
	if !ok {
		return nil, nil
	}

	return h.groupService.ListGroups(ctx)
}

// CategoriesQuery represents a query for listing all categories.
type CategoriesQuery struct{}

// NewCategoriesQuery creates a new CategoriesQuery instance.
func NewCategoriesQuery() CategoriesQuery {
	return CategoriesQuery{}
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
	_, ok := query.(CategoriesQuery)
	if !ok {
		return nil, nil
	}

	return h.categoryService.ListCategories(ctx)
}

// TracksQuery represents a query for listing all tracks.
type TracksQuery struct{}

// NewTracksQuery creates a new TracksQuery instance.
func NewTracksQuery() TracksQuery {
	return TracksQuery{}
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
	_, ok := query.(TracksQuery)
	if !ok {
		return nil, nil
	}

	return h.trackService.ListTracks(ctx)
}

// ThemesQuery represents a query for listing all themes.
type ThemesQuery struct{}

// NewThemesQuery creates a new ThemesQuery instance.
func NewThemesQuery() ThemesQuery {
	return ThemesQuery{}
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
	_, ok := query.(ThemesQuery)
	if !ok {
		return nil, nil
	}

	return h.themeService.ListThemes(ctx)
}

// ThemesByGroupQuery represents a query for listing all themes by group.
type ThemesByGroupQuery struct {
	GroupID string
}

// NewThemesByGroupQuery creates a new ThemesByGroupQuery instance.
func NewThemesByGroupQuery(groupID string) ThemesByGroupQuery {
	return ThemesByGroupQuery{
		GroupID: groupID,
	}
}

// Type returns the query type.
func (q ThemesByGroupQuery) Type() query.Type {
	return ThemesByGroupQueryType
}

// ThemesByGroupQueryHandler handles the themes by group query.
type ThemesByGroupQueryHandler struct {
	themeService ThemeService
}

// NewThemesByGroupQueryHandler creates a new ThemesByGroupQueryHandler instance.
func NewThemesByGroupQueryHandler(themeService ThemeService) ThemesByGroupQueryHandler {
	return ThemesByGroupQueryHandler{
		themeService: themeService,
	}
}

// Handle handles the themes by group query.
func (h ThemesByGroupQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	q, ok := query.(ThemesByGroupQuery)
	if !ok {
		return nil, nil
	}

	return h.themeService.ListThemesByGroup(ctx, q.GroupID)
}

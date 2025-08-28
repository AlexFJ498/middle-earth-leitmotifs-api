package listing

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
)

const (
	UsersQueryType      = "query.listing.users"
	MoviesQueryType     = "query.listing.movies"
	GroupsQueryType     = "query.listing.groups"
	CategoriesQueryType = "query.listing.categories"
	TracksQueryType     = "query.listing.tracks"
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
	return h.trackService.ListTracks(ctx)
}

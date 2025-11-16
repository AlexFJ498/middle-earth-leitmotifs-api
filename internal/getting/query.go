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

type MoviesQuery struct {
	ID string
}

func NewMoviesQuery(id string) MoviesQuery {
	return MoviesQuery{
		ID: id,
	}
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
	movieQuery, ok := query.(MoviesQuery)
	if !ok {
		return nil, nil
	}

	return h.movieService.GetMovie(ctx, movieQuery.ID)
}

type GroupsQuery struct {
	ID string
}

func NewGroupsQuery(id string) GroupsQuery {
	return GroupsQuery{
		ID: id,
	}
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
	groupQuery, ok := query.(GroupsQuery)
	if !ok {
		return nil, nil
	}

	return h.groupService.GetGroup(ctx, groupQuery.ID)
}

type CategoriesQuery struct {
	ID string
}

func NewCategoriesQuery(id string) CategoriesQuery {
	return CategoriesQuery{
		ID: id,
	}
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
	categoryQuery, ok := query.(CategoriesQuery)
	if !ok {
		return nil, nil
	}

	return h.categoryService.GetCategory(ctx, categoryQuery.ID)
}

type TracksQuery struct {
	ID string
}

func NewTracksQuery(id string) TracksQuery {
	return TracksQuery{
		ID: id,
	}
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
	trackQuery, ok := query.(TracksQuery)
	if !ok {
		return nil, nil
	}

	return h.trackService.GetTrack(ctx, trackQuery.ID)
}

type ThemesQuery struct {
	ID string
}

func NewThemesQuery(id string) ThemesQuery {
	return ThemesQuery{
		ID: id,
	}
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
	themeQuery, ok := query.(ThemesQuery)
	if !ok {
		return nil, nil
	}

	return h.themeService.GetTheme(ctx, themeQuery.ID)
}

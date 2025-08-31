package updating

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
)

const (
	MovieCommandType    command.Type = "command.update.movie"
	GroupCommandType    command.Type = "command.update.group"
	CategoryCommandType command.Type = "command.update.category"
	TrackCommandType    command.Type = "command.update.track"
	ThemeCommandType    command.Type = "command.update.theme"
)

// MovieCommand is the command dispatched to update a new movie.
type MovieCommand struct {
	id  string
	dto dto.MovieUpdateRequest
}

// NewMovieCommand updates a new MovieCommand instance.
func NewMovieCommand(id string, dto dto.MovieUpdateRequest) MovieCommand {
	return MovieCommand{
		id:  id,
		dto: dto,
	}
}

// Type returns the type of the command.
func (c MovieCommand) Type() command.Type {
	return MovieCommandType
}

// MovieCommandHandler is the handler responsible for updating movies.
type MovieCommandHandler struct {
	service MovieService
}

// NewMovieCommandHandler updates a new MovieCommandHandler instance.
func NewMovieCommandHandler(service MovieService) MovieCommandHandler {
	return MovieCommandHandler{
		service: service,
	}
}

// Handle processes the MovieCommand to update a movie.
func (h MovieCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	movieCmd, ok := cmd.(MovieCommand)
	if !ok {
		return nil
	}

	return h.service.UpdateMovie(ctx, movieCmd.id, movieCmd.dto)
}

// GroupCommand is the command dispatched to update a new group.
type GroupCommand struct {
	id  string
	dto dto.GroupUpdateRequest
}

// NewGroupCommand updates a new GroupCommand instance.
func NewGroupCommand(id string, dto dto.GroupUpdateRequest) GroupCommand {
	return GroupCommand{
		id:  id,
		dto: dto,
	}
}

// Type returns the type of the command.
func (c GroupCommand) Type() command.Type {
	return GroupCommandType
}

// GroupCommandHandler is the handler responsible for updating groups.
type GroupCommandHandler struct {
	service GroupService
}

// NewGroupCommandHandler updates a new GroupCommandHandler instance.
func NewGroupCommandHandler(service GroupService) GroupCommandHandler {
	return GroupCommandHandler{
		service: service,
	}
}

// Handle processes the GroupCommand to update a group.
func (h GroupCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	groupCmd, ok := cmd.(GroupCommand)
	if !ok {
		return nil
	}

	return h.service.UpdateGroup(ctx, groupCmd.id, groupCmd.dto)
}

// CategoryCommand is the command dispatched to update a new category.
type CategoryCommand struct {
	id  string
	dto dto.CategoryUpdateRequest
}

// NewCategoryCommand updates a new CategoryCommand instance.
func NewCategoryCommand(id string, dto dto.CategoryUpdateRequest) CategoryCommand {
	return CategoryCommand{
		id:  id,
		dto: dto,
	}
}

// Type returns the type of the command.
func (c CategoryCommand) Type() command.Type {
	return CategoryCommandType
}

// CategoryCommandHandler is the handler responsible for updating categories.
type CategoryCommandHandler struct {
	service CategoryService
}

// NewCategoryCommandHandler updates a new CategoryCommandHandler instance.
func NewCategoryCommandHandler(service CategoryService) CategoryCommandHandler {
	return CategoryCommandHandler{
		service: service,
	}
}

// Handle processes the CategoryCommand to update a category.
func (h CategoryCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	categoryCmd, ok := cmd.(CategoryCommand)
	if !ok {
		return nil
	}

	return h.service.UpdateCategory(ctx, categoryCmd.id, categoryCmd.dto)
}

// TrackCommand is the command dispatched to update a track.
type TrackCommand struct {
	id  string
	dto dto.TrackUpdateRequest
}

// NewTrackCommand creates a new TrackCommand instance.
func NewTrackCommand(id string, dto dto.TrackUpdateRequest) TrackCommand {
	return TrackCommand{
		id:  id,
		dto: dto,
	}
}

// Type returns the type of the command.
func (c TrackCommand) Type() command.Type {
	return TrackCommandType
}

// TrackCommandHandler is the handler responsible for updating tracks.
type TrackCommandHandler struct {
	service TrackService
}

// NewTrackCommandHandler creates a new TrackCommandHandler instance.
func NewTrackCommandHandler(service TrackService) TrackCommandHandler {
	return TrackCommandHandler{
		service: service,
	}
}

// Handle processes the TrackCommand to update a track.
func (h TrackCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	trackCmd, ok := cmd.(TrackCommand)
	if !ok {
		return nil
	}

	return h.service.UpdateTrack(ctx, trackCmd.id, trackCmd.dto)
}

// ThemeCommand is the command dispatched to update a new theme.
type ThemeCommand struct {
	id  string
	dto dto.ThemeUpdateRequest
}

// NewThemeCommand creates a new ThemeCommand instance.
func NewThemeCommand(id string, dto dto.ThemeUpdateRequest) ThemeCommand {
	return ThemeCommand{
		id:  id,
		dto: dto,
	}
}

// Type returns the type of the command.
func (c ThemeCommand) Type() command.Type {
	return ThemeCommandType
}

// ThemeCommandHandler is the handler responsible for updating themes.
type ThemeCommandHandler struct {
	service ThemeService
}

// NewThemeCommandHandler creates a new ThemeCommandHandler instance.
func NewThemeCommandHandler(service ThemeService) ThemeCommandHandler {
	return ThemeCommandHandler{
		service: service,
	}
}

// Handle processes the ThemeCommand to update a theme.
func (h ThemeCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	themeCmd, ok := cmd.(ThemeCommand)
	if !ok {
		return nil
	}

	return h.service.UpdateTheme(ctx, themeCmd.id, themeCmd.dto)
}

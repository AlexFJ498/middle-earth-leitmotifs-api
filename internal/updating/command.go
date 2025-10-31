package updating

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
)

const (
	MovieCommandType      command.Type = "command.update.movie"
	GroupCommandType      command.Type = "command.update.group"
	CategoryCommandType   command.Type = "command.update.category"
	TrackCommandType      command.Type = "command.update.track"
	ThemeCommandType      command.Type = "command.update.theme"
	TrackThemeCommandType command.Type = "command.update.track_theme"
)

type MovieCommand struct {
	id  string
	dto dto.MovieUpdateRequest
}

func NewMovieCommand(id string, dto dto.MovieUpdateRequest) MovieCommand {
	return MovieCommand{
		id:  id,
		dto: dto,
	}
}

func (c MovieCommand) Type() command.Type {
	return MovieCommandType
}

type MovieCommandHandler struct {
	service MovieService
}

func NewMovieCommandHandler(service MovieService) MovieCommandHandler {
	return MovieCommandHandler{
		service: service,
	}
}

func (h MovieCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	movieCmd, ok := cmd.(MovieCommand)
	if !ok {
		return nil
	}

	return h.service.UpdateMovie(ctx, movieCmd.id, movieCmd.dto)
}

type GroupCommand struct {
	id  string
	dto dto.GroupUpdateRequest
}

func NewGroupCommand(id string, dto dto.GroupUpdateRequest) GroupCommand {
	return GroupCommand{
		id:  id,
		dto: dto,
	}
}

func (c GroupCommand) Type() command.Type {
	return GroupCommandType
}

type GroupCommandHandler struct {
	service GroupService
}

func NewGroupCommandHandler(service GroupService) GroupCommandHandler {
	return GroupCommandHandler{
		service: service,
	}
}

func (h GroupCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	groupCmd, ok := cmd.(GroupCommand)
	if !ok {
		return nil
	}

	return h.service.UpdateGroup(ctx, groupCmd.id, groupCmd.dto)
}

type CategoryCommand struct {
	id  string
	dto dto.CategoryUpdateRequest
}

func NewCategoryCommand(id string, dto dto.CategoryUpdateRequest) CategoryCommand {
	return CategoryCommand{
		id:  id,
		dto: dto,
	}
}

func (c CategoryCommand) Type() command.Type {
	return CategoryCommandType
}

type CategoryCommandHandler struct {
	service CategoryService
}

func NewCategoryCommandHandler(service CategoryService) CategoryCommandHandler {
	return CategoryCommandHandler{
		service: service,
	}
}

func (h CategoryCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	categoryCmd, ok := cmd.(CategoryCommand)
	if !ok {
		return nil
	}

	return h.service.UpdateCategory(ctx, categoryCmd.id, categoryCmd.dto)
}

type TrackCommand struct {
	id  string
	dto dto.TrackUpdateRequest
}

func NewTrackCommand(id string, dto dto.TrackUpdateRequest) TrackCommand {
	return TrackCommand{
		id:  id,
		dto: dto,
	}
}

func (c TrackCommand) Type() command.Type {
	return TrackCommandType
}

type TrackCommandHandler struct {
	service TrackService
}

func NewTrackCommandHandler(service TrackService) TrackCommandHandler {
	return TrackCommandHandler{
		service: service,
	}
}

func (h TrackCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	trackCmd, ok := cmd.(TrackCommand)
	if !ok {
		return nil
	}

	return h.service.UpdateTrack(ctx, trackCmd.id, trackCmd.dto)
}

type ThemeCommand struct {
	id  string
	dto dto.ThemeUpdateRequest
}

func NewThemeCommand(id string, dto dto.ThemeUpdateRequest) ThemeCommand {
	return ThemeCommand{
		id:  id,
		dto: dto,
	}
}

func (c ThemeCommand) Type() command.Type {
	return ThemeCommandType
}

type ThemeCommandHandler struct {
	service ThemeService
}

func NewThemeCommandHandler(service ThemeService) ThemeCommandHandler {
	return ThemeCommandHandler{
		service: service,
	}
}

func (h ThemeCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	themeCmd, ok := cmd.(ThemeCommand)
	if !ok {
		return nil
	}

	return h.service.UpdateTheme(ctx, themeCmd.id, themeCmd.dto)
}

type TrackThemeCommand struct {
	dto dto.TrackThemeUpdateRequest
}

func NewTrackThemeCommand(dto dto.TrackThemeUpdateRequest) TrackThemeCommand {
	return TrackThemeCommand{
		dto: dto,
	}
}

func (c TrackThemeCommand) Type() command.Type {
	return TrackThemeCommandType
}

type TrackThemeCommandHandler struct {
	service TrackThemeService
}

func NewTrackThemeCommandHandler(service TrackThemeService) TrackThemeCommandHandler {
	return TrackThemeCommandHandler{
		service: service,
	}
}

func (h TrackThemeCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	trackThemeCmd, ok := cmd.(TrackThemeCommand)
	if !ok {
		return nil
	}

	return h.service.UpdateTrackTheme(ctx, trackThemeCmd.dto)
}

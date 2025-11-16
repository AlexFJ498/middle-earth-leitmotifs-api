package creating

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
)

const (
	UserCommandType       command.Type = "command.create.user"
	MovieCommandType      command.Type = "command.create.movie"
	GroupCommandType      command.Type = "command.create.group"
	CategoryCommandType   command.Type = "command.create.category"
	TrackCommandType      command.Type = "command.create.track"
	ThemeCommandType      command.Type = "command.create.theme"
	TrackThemeCommandType command.Type = "command.create.track_theme"
)

type UserCommand struct {
	dto dto.UserCreateRequest
}

func NewUserCommand(dto dto.UserCreateRequest) UserCommand {
	return UserCommand{
		dto: dto,
	}
}

func (c UserCommand) Type() command.Type {
	return UserCommandType
}

type UserCommandHandler struct {
	service UserService
}

func NewUserCommandHandler(service UserService) UserCommandHandler {
	return UserCommandHandler{
		service: service,
	}
}

func (h UserCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	userCmd, ok := cmd.(UserCommand)
	if !ok {
		return nil
	}

	return h.service.CreateUser(ctx, userCmd.dto)
}

type MovieCommand struct {
	dto dto.MovieCreateRequest
}

func NewMovieCommand(dto dto.MovieCreateRequest) MovieCommand {
	return MovieCommand{
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

	return h.service.CreateMovie(ctx, movieCmd.dto)
}

type GroupCommand struct {
	dto dto.GroupCreateRequest
}

func NewGroupCommand(dto dto.GroupCreateRequest) GroupCommand {
	return GroupCommand{
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

	return h.service.CreateGroup(ctx, groupCmd.dto)
}

type CategoryCommand struct {
	dto dto.CategoryCreateRequest
}

func NewCategoryCommand(dto dto.CategoryCreateRequest) CategoryCommand {
	return CategoryCommand{
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

	return h.service.CreateCategory(ctx, categoryCmd.dto)
}

type TrackCommand struct {
	dto dto.TrackCreateRequest
}

func NewTrackCommand(dto dto.TrackCreateRequest) TrackCommand {
	return TrackCommand{
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

	return h.service.CreateTrack(ctx, trackCmd.dto)
}

type ThemeCommand struct {
	dto dto.ThemeCreateRequest
}

func NewThemeCommand(dto dto.ThemeCreateRequest) ThemeCommand {
	return ThemeCommand{
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

	return h.service.CreateTheme(ctx, themeCmd.dto)
}

type TrackThemeCommand struct {
	dto dto.TrackThemeCreateRequest
}

func NewTrackThemeCommand(dto dto.TrackThemeCreateRequest) TrackThemeCommand {
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

	return h.service.CreateTrackTheme(ctx, trackThemeCmd.dto)
}

package deleting

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
)

const (
	MovieCommandType      = "command.delete.movie"
	GroupCommandType      = "command.delete.group"
	CategoryCommandType   = "command.delete.category"
	TrackCommandType      = "command.delete.track"
	ThemeCommandType      = "command.delete.theme"
	TrackThemeCommandType = "command.delete.track_theme"
)

type MovieCommand struct {
	ID string
}

func NewMovieCommand(id string) MovieCommand {
	return MovieCommand{
		ID: id,
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

	movieID, err := domain.NewMovieIDFromString(movieCmd.ID)
	if err != nil {
		return err
	}
	return h.service.DeleteMovie(ctx, movieID)
}

type GroupCommand struct {
	ID string
}

func NewGroupCommand(id string) GroupCommand {
	return GroupCommand{
		ID: id,
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

	groupID, err := domain.NewGroupIDFromString(groupCmd.ID)
	if err != nil {
		return err
	}
	return h.service.DeleteGroup(ctx, groupID)
}

type CategoryCommand struct {
	ID string
}

func NewCategoryCommand(id string) CategoryCommand {
	return CategoryCommand{
		ID: id,
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

	categoryID, err := domain.NewCategoryIDFromString(categoryCmd.ID)
	if err != nil {
		return err
	}
	return h.service.DeleteCategory(ctx, categoryID)
}

type TrackCommand struct {
	ID string
}

func NewTrackCommand(id string) TrackCommand {
	return TrackCommand{
		ID: id,
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

	trackID, err := domain.NewTrackIDFromString(trackCmd.ID)
	if err != nil {
		return err
	}
	return h.service.DeleteTrack(ctx, trackID)
}

type ThemeCommand struct {
	ID string
}

func NewThemeCommand(id string) ThemeCommand {
	return ThemeCommand{
		ID: id,
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

	themeID, err := domain.NewThemeIDFromString(themeCmd.ID)
	if err != nil {
		return err
	}
	return h.service.DeleteTheme(ctx, themeID)
}

type TrackThemeCommand struct {
	dto dto.TrackThemeDeleteRequest
}

func NewTrackThemeCommand(dto dto.TrackThemeDeleteRequest) TrackThemeCommand {
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

	trackID, err := domain.NewTrackIDFromString(trackThemeCmd.dto.TrackID)
	if err != nil {
		return err
	}

	themeID, err := domain.NewThemeIDFromString(trackThemeCmd.dto.ThemeID)
	if err != nil {
		return err
	}

	startSecond, err := domain.NewStartSecond(trackThemeCmd.dto.StartSecond)
	if err != nil {
		return err
	}

	return h.service.DeleteTrackTheme(ctx, trackID, themeID, startSecond)
}

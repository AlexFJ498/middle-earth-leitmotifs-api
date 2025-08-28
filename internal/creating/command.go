package creating

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
)

const (
	UserCommandType     command.Type = "command.create.user"
	MovieCommandType    command.Type = "command.create.movie"
	GroupCommandType    command.Type = "command.create.group"
	CategoryCommandType command.Type = "command.create.category"
	TrackCommandType    command.Type = "command.create.track"
)

// UserCommand is the command dispatched to create a new user.
type UserCommand struct {
	dto dto.UserCreateRequest
}

// NewUserCommand creates a new UserCommand instance.
func NewUserCommand(dto dto.UserCreateRequest) UserCommand {
	return UserCommand{
		dto: dto,
	}
}

// Type returns the type of the command.
func (c UserCommand) Type() command.Type {
	return UserCommandType
}

// UserCommandHandler is the handler responsible for creating users.
type UserCommandHandler struct {
	service UserService
}

// NewUserCommandHandler creates a new UserCommandHandler instance.
func NewUserCommandHandler(service UserService) UserCommandHandler {
	return UserCommandHandler{
		service: service,
	}
}

// Handle processes the UserCommand to create a new user.
func (h UserCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	userCmd, ok := cmd.(UserCommand)
	if !ok {
		return nil
	}

	return h.service.CreateUser(ctx, userCmd.dto)
}

// MovieCommand is the command dispatched to create a new movie.
type MovieCommand struct {
	dto dto.MovieCreateRequest
}

// NewMovieCommand creates a new MovieCommand instance.
func NewMovieCommand(dto dto.MovieCreateRequest) MovieCommand {
	return MovieCommand{
		dto: dto,
	}
}

// Type returns the type of the command.
func (c MovieCommand) Type() command.Type {
	return MovieCommandType
}

// MovieCommandHandler is the handler responsible for creating movies.
type MovieCommandHandler struct {
	service MovieService
}

// NewMovieCommandHandler creates a new MovieCommandHandler instance.
func NewMovieCommandHandler(service MovieService) MovieCommandHandler {
	return MovieCommandHandler{
		service: service,
	}
}

// Handle processes the MovieCommand to create a new movie.
func (h MovieCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	movieCmd, ok := cmd.(MovieCommand)
	if !ok {
		return nil
	}

	return h.service.CreateMovie(ctx, movieCmd.dto)
}

// GroupCommand is the command dispatched to create a new group.
type GroupCommand struct {
	dto dto.GroupCreateRequest
}

// NewGroupCommand creates a new GroupCommand instance.
func NewGroupCommand(dto dto.GroupCreateRequest) GroupCommand {
	return GroupCommand{
		dto: dto,
	}
}

// Type returns the type of the command.
func (c GroupCommand) Type() command.Type {
	return GroupCommandType
}

// GroupCommandHandler is the handler responsible for creating groups.
type GroupCommandHandler struct {
	service GroupService
}

// NewGroupCommandHandler creates a new GroupCommandHandler instance.
func NewGroupCommandHandler(service GroupService) GroupCommandHandler {
	return GroupCommandHandler{
		service: service,
	}
}

// Handle processes the GroupCommand to create a new group.
func (h GroupCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	groupCmd, ok := cmd.(GroupCommand)
	if !ok {
		return nil
	}

	return h.service.CreateGroup(ctx, groupCmd.dto)
}

// CategoryCommand is the command dispatched to create a new category.
type CategoryCommand struct {
	dto dto.CategoryCreateRequest
}

// NewCategoryCommand creates a new CategoryCommand instance.
func NewCategoryCommand(dto dto.CategoryCreateRequest) CategoryCommand {
	return CategoryCommand{
		dto: dto,
	}
}

// Type returns the type of the command.
func (c CategoryCommand) Type() command.Type {
	return CategoryCommandType
}

// CategoryCommandHandler is the handler responsible for creating categories.
type CategoryCommandHandler struct {
	service CategoryService
}

// NewCategoryCommandHandler creates a new CategoryCommandHandler instance.
func NewCategoryCommandHandler(service CategoryService) CategoryCommandHandler {
	return CategoryCommandHandler{
		service: service,
	}
}

// Handle processes the CategoryCommand to create a new category.
func (h CategoryCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	categoryCmd, ok := cmd.(CategoryCommand)
	if !ok {
		return nil
	}

	return h.service.CreateCategory(ctx, categoryCmd.dto)
}

// TrackCommand is the command dispatched to create a new track.
type TrackCommand struct {
	dto dto.TrackCreateRequest
}

// NewTrackCommand creates a new TrackCommand instance.
func NewTrackCommand(dto dto.TrackCreateRequest) TrackCommand {
	return TrackCommand{
		dto: dto,
	}
}

// Type returns the type of the command.
func (c TrackCommand) Type() command.Type {
	return TrackCommandType
}

// TrackCommandHandler is the handler responsible for creating tracks.
type TrackCommandHandler struct {
	service TrackService
}

// NewTrackCommandHandler creates a new TrackCommandHandler instance.
func NewTrackCommandHandler(service TrackService) TrackCommandHandler {
	return TrackCommandHandler{
		service: service,
	}
}

// Handle processes the TrackCommand to create a new track.
func (h TrackCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	trackCmd, ok := cmd.(TrackCommand)
	if !ok {
		return nil
	}

	return h.service.CreateTrack(ctx, trackCmd.dto)
}

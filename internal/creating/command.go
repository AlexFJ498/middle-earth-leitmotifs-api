package creating

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
)

const (
	UserCommandType  command.Type = "command.create.user"
	MovieCommandType command.Type = "command.create.movie"
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

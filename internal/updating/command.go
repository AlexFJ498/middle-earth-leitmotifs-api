package updating

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
)

const (
	MovieCommandType command.Type = "command.update.movie"
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

package deleting

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
)

const (
	MovieCommandType = "command.delete.movie"
)

// MovieCommand is the command dispatched for deleting a movie.
type MovieCommand struct {
	ID string
}

// NewMovieCommand creates a new MovieCommand instance.
func NewMovieCommand(id string) MovieCommand {
	return MovieCommand{
		ID: id,
	}
}

// Type returns the type of the command.
func (c MovieCommand) Type() command.Type {
	return MovieCommandType
}

// MovieCommandHandler handles the deletion of a movie.
type MovieCommandHandler struct {
	service MovieService
}

// NewMovieCommandHandler creates a new MovieCommandHandler.
func NewMovieCommandHandler(service MovieService) MovieCommandHandler {
	return MovieCommandHandler{
		service: service,
	}
}

// Handle processes the MovieCommand.
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

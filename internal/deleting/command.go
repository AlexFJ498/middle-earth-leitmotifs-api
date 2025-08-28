package deleting

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
)

const (
	MovieCommandType    = "command.delete.movie"
	GroupCommandType    = "command.delete.group"
	CategoryCommandType = "command.delete.category"
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

// GroupCommand is the command dispatched for deleting a group.
type GroupCommand struct {
	ID string
}

// NewGroupCommand creates a new GroupCommand instance.
func NewGroupCommand(id string) GroupCommand {
	return GroupCommand{
		ID: id,
	}
}

// Type returns the type of the command.
func (c GroupCommand) Type() command.Type {
	return GroupCommandType
}

// GroupCommandHandler handles the deletion of a group.
type GroupCommandHandler struct {
	service GroupService
}

// NewGroupCommandHandler creates a new GroupCommandHandler.
func NewGroupCommandHandler(service GroupService) GroupCommandHandler {
	return GroupCommandHandler{
		service: service,
	}
}

// Handle processes the GroupCommand.
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

// CategoryCommand is the command dispatched for deleting a category.
type CategoryCommand struct {
	ID string
}

// NewCategoryCommand creates a new CategoryCommand instance.
func NewCategoryCommand(id string) CategoryCommand {
	return CategoryCommand{
		ID: id,
	}
}

// Type returns the type of the command.
func (c CategoryCommand) Type() command.Type {
	return CategoryCommandType
}

// CategoryCommandHandler handles the deletion of a category.
type CategoryCommandHandler struct {
	service CategoryService
}

// NewCategoryCommandHandler creates a new CategoryCommandHandler.
func NewCategoryCommandHandler(service CategoryService) CategoryCommandHandler {
	return CategoryCommandHandler{
		service: service,
	}
}

// Handle processes the CategoryCommand.
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

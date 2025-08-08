package creating

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
)

const UserCommandType command.Type = "command.create.user"

// UserCommand is the command dispatched to create a new user.
type UserCommand struct {
	ID    string
	Name  string
	Email string
}

// NewUserCommand creates a new UserCommand instance.
func NewUserCommand(id, name, email string) UserCommand {
	return UserCommand{
		ID:    id,
		Name:  name,
		Email: email,
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

	return h.service.CreateUser(ctx, userCmd.ID, userCmd.Name, userCmd.Email)
}

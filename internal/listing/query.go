package listing

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
)

const (
	UsersQueryType = "query.listing.users"
)

// UsersQuery represents a query for listing all users.
type UsersQuery struct{}

// NewUsersQuery creates a new UsersQuery instance.
func NewUsersQuery() UsersQuery {
	return UsersQuery{}
}

// Type returns the query type.
func (q UsersQuery) Type() query.Type {
	return UsersQueryType
}

// UsersQueryHandler handles the users query.
type UsersQueryHandler struct {
	userService UserService
}

// NewUsersQueryHandler creates a new UsersQueryHandler instance.
func NewUsersQueryHandler(userService UserService) UsersQueryHandler {
	return UsersQueryHandler{
		userService: userService,
	}
}

// Handle handles the users query.
func (h UsersQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	return h.userService.ListUsers(ctx)
}

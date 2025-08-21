package authenticating

import (
	"context"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
)

const LoginQueryType = "query.authenticating.login"

// LoginQuery represents a query for user login.
type LoginQuery struct {
	Email    string
	Password string
}

// NewLoginQuery creates a new LoginQuery instance.
func NewLoginQuery(email, password string) LoginQuery {
	return LoginQuery{
		Email:    email,
		Password: password,
	}
}

// Type returns the query type.
func (q LoginQuery) Type() query.Type {
	return LoginQueryType
}

// LoginQueryHandler handles the login query.
type LoginQueryHandler struct {
	service LoginService
}

// NewLoginQueryHandler creates a new LoginQueryHandler instance.
func NewLoginQueryHandler(service LoginService) LoginQueryHandler {
	return LoginQueryHandler{
		service: service,
	}
}

// Handle processes the login query.
func (h LoginQueryHandler) Handle(ctx context.Context, query query.Query) (any, error) {
	loginQuery, ok := query.(LoginQuery)
	if !ok {
		return nil, nil
	}

	return h.service.LoginUser(ctx, loginQuery.Email, loginQuery.Password)
}

package listing

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
)

// UserService is the default implementation of the UserService interface
type UserService struct {
	userRepository domain.UserRepository
}

// NewUserService returns a new UserService instance.
func NewUserService(userRepository domain.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

// ListUsers implements the UserService interface for listing all users.
func (s UserService) ListUsers(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := s.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	userResponses := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		userResponses = append(userResponses, dto.NewUserResponse(
			user.ID().String(),
			user.Name().String(),
			user.Email().String(),
		))
	}

	return userResponses, nil
}

package authenticating

import (
	"context"
	"fmt"
	"time"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/auth"
)

// LoginService is the default implementation of the LoginService interface
type LoginService struct {
	userRepository domain.UserRepository
	jwtKey         auth.JWTKey
	exp            time.Duration
}

// NewLoginService creates a new instance of LoginService.
func NewLoginService(userRepository domain.UserRepository, jwtKey auth.JWTKey, exp time.Duration) LoginService {
	fmt.Println(jwtKey, exp)
	return LoginService{
		userRepository: userRepository,
		jwtKey:         jwtKey,
		exp:            exp,
	}
}

// LoginUser handles user login.
func (s LoginService) LoginUser(ctx context.Context, email, password string) (string, error) {
	// Obtain the user by email
	emailVO, err := domain.NewUserEmail(email)
	if err != nil {
		return "", err
	}

	// Find the user in the repository
	user, err := s.userRepository.FindByEmail(ctx, emailVO)
	if err != nil {
		return "", err
	}

	// Check if the provided password matches the stored hashed password
	if err := auth.CheckPassword(user.Password().String(), password); err != nil {
		return "", domain.ErrInvalidUserPassword
	}

	// Generate a new JWT token for the user
	token, err := auth.GenerateJWTKey(user, s.jwtKey, s.exp)
	if err != nil {
		return "", err
	}

	// Return the generated token
	return token, nil
}

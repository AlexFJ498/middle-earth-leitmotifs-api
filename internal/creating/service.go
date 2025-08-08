package creating

// UserService is the default implementation of the UserService interface
// returned by creating.NewUserService.
import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/event"
)

type UserService struct {
	userRepository domain.UserRepository
	eventBus       event.Bus
}

// NewUserService returns a new UserService instance.
func NewUserService(userRepository domain.UserRepository, eventBus event.Bus) UserService {
	return UserService{
		userRepository: userRepository,
		eventBus:       eventBus,
	}
}

// CreateUser implements the UserService interface for creating a new user.
func (s UserService) CreateUser(ctx context.Context, id, name, email string) error {
	user, err := domain.NewUser(id, name, email)
	if err != nil {
		return err
	}

	if err := s.userRepository.Save(ctx, user); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, user.PullEvents())
}

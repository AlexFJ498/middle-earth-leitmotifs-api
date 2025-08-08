package increasing

type UserCounterIncreaserService struct{}

func NewUserCounterIncreaserService() UserCounterIncreaserService {
	return UserCounterIncreaserService{}
}

// Increase increments the user counter.
// At the moment, this is not implemented. It shows how an inmemory event bus can be used to handle events.
func (s UserCounterIncreaserService) Increase(id string) error {
	return nil
}

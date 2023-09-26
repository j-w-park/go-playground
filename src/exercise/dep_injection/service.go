package dep_injection

import "fmt"

type Service interface {
	SayHelloTo(id string) (string, error)
}

type SimpleService struct {
	logger Logger
	store  DataStore
}

func (service SimpleService) SayHelloTo(id string) (string, error) {
	if user, ok := service.store.GetUserNameFromId(id); ok {
		service.logger.Log("SayHello: " + id)
		return "Hello " + user, nil
	} else {
		return "", fmt.Errorf("user not found: %s", id)
	}
}

func (service SimpleService) SayByeTo(id string) (string, error) {
	if user, ok := service.store.GetUserNameFromId(id); ok {
		service.logger.Log("SayBye: " + id)
		return "Bye " + user, nil
	} else {
		return "", fmt.Errorf("user not found: %s", id)
	}
}

func NewSimpleService(logger Logger, store DataStore) SimpleService {
	return SimpleService{logger: logger, store: store}
}

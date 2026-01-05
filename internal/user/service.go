package user

import "github.com/shegai01/msg/internal/shared"

type Service struct {
	Events chan<- shared.UserCreatedEvent
}

func (s *Service) UserCreated(userName, email string) {
	s.Events <- shared.UserCreatedEvent{
		UserName: userName,
		Email:    email,
	}
}

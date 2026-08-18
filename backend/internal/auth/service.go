package auth

import (
	"kanban-app-backend/internal/user"
)

type Service struct {
	userRepository *user.Repository
}

func (s *Service) CreateNewUser(username, email, plainPassword string) {

}

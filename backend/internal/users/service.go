package users

import (
	"context"
	"errors"
	"kanban-app-backend/pkg"
)

var (
	ErrInvalidCredentials = errors.New("Invalid credentials. Please check your credentials and try again.")
)

type Service struct {
	userRepository *Repository
}

func NewService(r *Repository) *Service {
	return &Service{userRepository: r}
}

func (s *Service) CreateNewUser(ctx context.Context, username, email, plainPassword string) (*User, error) {
	existingUserEmail, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if existingUserEmail != nil {
		return nil, ErrInvalidCredentials
	}
	existingUserUsername, err := s.userRepository.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if existingUserUsername != nil {
		return nil, ErrInvalidCredentials
	}

	hashedPassword, err := pkg.HashPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	createdUser, err := s.userRepository.Create(ctx, username, email, hashedPassword)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

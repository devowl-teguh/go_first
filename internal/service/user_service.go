package service

import (
	"errors"
	"go_first/internal/models"
	"go_first/internal/repository"
	"time"
)

// UserService defines the interface for user business logic
type UserService interface {
	CreateUser(name, email string) (*models.User, error)
	GetUser(id int) (*models.User, error)
}

// userService implements UserService
type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a new instance of userService
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// CreateUser handles the business logic for creating a user
func (s *userService) CreateUser(name, email string) (*models.User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	user := &models.User{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser handles the business logic for retrieving a user
func (s *userService) GetUser(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}

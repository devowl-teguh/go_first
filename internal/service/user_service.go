package service

import (
	"errors"
	"go_first/internal/models"
	"go_first/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// UserService defines the interface for user business logic
type UserService interface {
	CreateUser(name, email, phoneNumber, pin string) (*models.User, error)
	GetUser(id int) (*models.User, error)
	Login(phoneNumber, pin string) (*models.User, error)
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
func (s *userService) CreateUser(name, email, phoneNumber, pin string) (*models.User, error) {
	if name == "" || email == "" || phoneNumber == "" || pin == "" {
		return nil, errors.New("all fields are required")
	}

	// Hash PIN
	hashedPIN, err := bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		PIN:         string(hashedPIN),
		CreatedAt:   time.Now(),
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

// Login verifies phone number and PIN
func (s *userService) Login(phoneNumber, pin string) (*models.User, error) {
	user, err := s.repo.GetByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, errors.New("invalid phone number or pin")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PIN), []byte(pin)); err != nil {
		return nil, errors.New("invalid phone number or pin")
	}

	return user, nil
}

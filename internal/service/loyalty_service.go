package service

import (
	"errors"
	"fmt"
	"go_first/internal/models"
	"go_first/internal/repository"
	"time"
)

// LoyaltyService defines the interface for loyalty business logic
type LoyaltyService interface {
	EarnPoints(userID int, amount int, description string) error
	RedeemPoints(userID int, amount int) error
	GetBalance(userID int) (int, error)
	GetHistory(userID int) ([]models.LoyaltyTransaction, error)
}

// loyaltyService implements LoyaltyService
type loyaltyService struct {
	repo repository.LoyaltyRepository
}

// NewLoyaltyService creates a new instance of loyaltyService
func NewLoyaltyService(repo repository.LoyaltyRepository) LoyaltyService {
	return &loyaltyService{
		repo: repo,
	}
}

// EarnPoints adds points to a user's account
func (s *loyaltyService) EarnPoints(userID int, amount int, description string) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	tx := &models.LoyaltyTransaction{
		UserID:    userID,
		Amount:    amount,
		Type:      "EARN",
		Reference: description,
		CreatedAt: time.Now(),
	}

	return s.repo.AddTransaction(tx)
}

// RedeemPoints deducts points from a user's account if sufficient balance exists
func (s *loyaltyService) RedeemPoints(userID int, amount int) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	// Check balance first
	balance, err := s.repo.GetBalance(userID)
	if err != nil {
		return fmt.Errorf("failed to get balance: %w", err)
	}

	if balance < amount {
		return errors.New("insufficient points balance")
	}

	tx := &models.LoyaltyTransaction{
		UserID:    userID,
		Amount:    -amount, // Negative amount for redemption
		Type:      "REDEEM",
		Reference: "Redemption",
		CreatedAt: time.Now(),
	}

	return s.repo.AddTransaction(tx)
}

// GetBalance retrieves the current points balance for a user
func (s *loyaltyService) GetBalance(userID int) (int, error) {
	return s.repo.GetBalance(userID)
}

// GetHistory retrieves the transaction history for a user
func (s *loyaltyService) GetHistory(userID int) ([]models.LoyaltyTransaction, error) {
	return s.repo.GetHistory(userID)
}

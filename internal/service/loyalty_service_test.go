package service

import (
	"go_first/internal/models"
	"testing"
)

// MockLoyaltyRepository is a mock implementation of LoyaltyRepository
type MockLoyaltyRepository struct {
	balance int
	txs     []models.LoyaltyTransaction
}

func (m *MockLoyaltyRepository) GetBalance(userID int) (int, error) {
	return m.balance, nil
}

func (m *MockLoyaltyRepository) AddTransaction(tx *models.LoyaltyTransaction) error {
	m.txs = append(m.txs, *tx)
	m.balance += tx.Amount
	return nil
}

func (m *MockLoyaltyRepository) GetHistory(userID int) ([]models.LoyaltyTransaction, error) {
	return m.txs, nil
}

func TestEarnPoints(t *testing.T) {
	repo := &MockLoyaltyRepository{}
	svc := NewLoyaltyService(repo)

	err := svc.EarnPoints(1, 100, "Test Earn")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	balance, _ := svc.GetBalance(1)
	if balance != 100 {
		t.Errorf("Expected balance 100, got %d", balance)
	}
}

func TestRedeemPoints_Success(t *testing.T) {
	repo := &MockLoyaltyRepository{balance: 100}
	svc := NewLoyaltyService(repo)

	err := svc.RedeemPoints(1, 50)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	balance, _ := svc.GetBalance(1)
	if balance != 50 {
		t.Errorf("Expected balance 50, got %d", balance)
	}
}

func TestRedeemPoints_InsufficientFunds(t *testing.T) {
	repo := &MockLoyaltyRepository{balance: 30}
	svc := NewLoyaltyService(repo)

	err := svc.RedeemPoints(1, 50)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "insufficient points balance" {
		t.Errorf("Expected 'insufficient points balance', got '%s'", err.Error())
	}
}

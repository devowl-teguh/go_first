package service

import (
	"go_first/internal/models"
	"go_first/internal/repository"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestUserLogin(t *testing.T) {
	hashedPIN, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	repo := repository.NewInMemoryUserRepository()
	repo.Create(&models.User{
		ID:          1,
		PhoneNumber: "08123456789",
		PIN:         string(hashedPIN),
	})

	svc := NewUserService(repo)

	// Test Success
	user, err := svc.Login("08123456789", "123456")
	if err != nil {
		t.Fatalf("Expected success, got %v", err)
	}
	if user.ID != 1 {
		t.Errorf("Expected user ID 1, got %d", user.ID)
	}

	// Test Failure (Wrong PIN)
	_, err = svc.Login("08123456789", "wrongpin")
	if err == nil {
		t.Fatal("Expected error for wrong PIN, got nil")
	}

	// Test Failure (Wrong Phone)
	_, err = svc.Login("99999999999", "123456")
	if err == nil {
		t.Fatal("Expected error for wrong phone, got nil")
	}
}

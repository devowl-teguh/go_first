package repository

import (
	"errors"
	"go_first/internal/models"
	"gorm.io/gorm"
)

// LoyaltyRepository defines the interface for loyalty data access
type LoyaltyRepository interface {
	GetBalance(userID int) (int, error)
	AddTransaction(tx *models.LoyaltyTransaction) error
	GetHistory(userID int) ([]models.LoyaltyTransaction, error)
}

// GormLoyaltyRepository implements LoyaltyRepository using GORM
type GormLoyaltyRepository struct {
	db *gorm.DB
}

// NewGormLoyaltyRepository creates a new GormLoyaltyRepository
func NewGormLoyaltyRepository(db *gorm.DB) *GormLoyaltyRepository {
	return &GormLoyaltyRepository{
		db: db,
	}
}

// GetBalance retrieves the current points balance for a user
func (r *GormLoyaltyRepository) GetBalance(userID int) (int, error) {
	var account models.LoyaltyAccount
	if err := r.db.Where("user_id = ?", userID).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil // User has no specific loyalty account yet, so 0 points
		}
		return 0, err
	}
	return account.Points, nil
}

// AddTransaction adds a transaction and updates the user's balance atomically
func (r *GormLoyaltyRepository) AddTransaction(tx *models.LoyaltyTransaction) error {
	return r.db.Transaction(func(dbTx *gorm.DB) error {
		// 1. Create the transaction record
		if err := dbTx.Create(tx).Error; err != nil {
			return err
		}

		// 2. Update or Create the LoyaltyAccount
		var account models.LoyaltyAccount
		if err := dbTx.Where("user_id = ?", tx.UserID).First(&account).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Create new account
				account = models.LoyaltyAccount{
					UserID: tx.UserID,
					Points: tx.Amount,
				}
				if err := dbTx.Create(&account).Error; err != nil {
					return err
				}
				return nil
			}
			return err
		}

		// Update existing account
		account.Points += tx.Amount
		if err := dbTx.Save(&account).Error; err != nil {
			return err
		}

		return nil
	})
}

// GetHistory retrieves the transaction history for a user
func (r *GormLoyaltyRepository) GetHistory(userID int) ([]models.LoyaltyTransaction, error) {
	var history []models.LoyaltyTransaction
	if err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&history).Error; err != nil {
		return nil, err
	}
	return history, nil
}

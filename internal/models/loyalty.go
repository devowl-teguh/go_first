package models

import "time"

// LoyaltyAccount represents a user's loyalty account
type LoyaltyAccount struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id" gorm:"uniqueIndex"`
	Points    int       `json:"points"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoyaltyTransaction represents a point transaction
type LoyaltyTransaction struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id" gorm:"index"`
	Amount    int       `json:"amount"` // Can be positive (earn) or negative (redeem)
	Type      string    `json:"type"`   // "EARN" or "REDEEM"
	Reference string    `json:"reference"`
	CreatedAt time.Time `json:"created_at"`
}

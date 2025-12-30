package models

import "time"

// User represents a user in the system
type User struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number" gorm:"uniqueIndex"`
	PIN         string    `json:"-"` // Store hashed PIN, don't expose in JSON
	CreatedAt   time.Time `json:"created_at"`
}

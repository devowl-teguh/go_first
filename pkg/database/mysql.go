package database

import (
	"fmt"
	"log"

	"go_first/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectMySQL initializes a connection to the MySQL database
func ConnectMySQL(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Successfully connected to MySQL database")
	return db, nil
}

// Migrate performs database migrations
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}

package main

import (
	"go_first/internal/config"
	"go_first/internal/models"
	pkg_database "go_first/pkg/database"
	"log"
	"time"
)

func main() {
	// 1. Config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Database
	db, err := pkg_database.ConnectMySQL(cfg.DSN)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// 3. Dummy Data
	users := []models.User{
		{Name: "Alice Smith", Email: "alice@example.com", PhoneNumber: "08123456789", PIN: "$2a$10$8K1p/vX/8.8K1p/vX/8.8K1p/vX/8.8K1p/vX/8.8K1p/vX/8.", CreatedAt: time.Now()}, // PIN: 123456
		{Name: "Bob Johnson", Email: "bob@example.com", PhoneNumber: "08123456790", PIN: "$2a$10$8K1p/vX/8.8K1p/vX/8.8K1p/vX/8.8K1p/vX/8.8K1p/vX/8.", CreatedAt: time.Now()},
	}

	// 4. Seeding
	log.Println("Seeding users...")
	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Printf("Failed to seed user %s: %v", user.Name, err)
		} else {
			log.Printf("Seeded user: %s (ID: %d)", user.Name, user.ID)
		}
	}

	log.Println("Seeding completed successfully")
}

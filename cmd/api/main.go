package main

import (
	"fmt"
	"go_first/internal/api"
	"go_first/internal/config"
	"go_first/internal/handler"
	"go_first/internal/repository"
	"go_first/internal/service"
	pkg_database "go_first/pkg/database"
	"log"
	"net/http"
)

func main() {
	// 1. Config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Dependency Injection / Wiring
	// 2. Database
	db, err := pkg_database.ConnectMySQL(cfg.DSN)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Auto Migrate
	if err := pkg_database.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 2. Repository
	// userRepo := repository.NewInMemoryUserRepository()
	userRepo := repository.NewGormUserRepository(db)

	// 2. Service
	userService := service.NewUserService(userRepo)

	// 3. Handler
	userHandler := handler.NewUserHandler(userService)

	// Router Setup
	mux := http.NewServeMux()
	api.RegisterRoutes(mux, userHandler)

	// Server Start
	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

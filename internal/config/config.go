package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	DSN       string
	JWTSecret string
}

// Load loads the configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default env vars")
	}

	dsn := os.Getenv("DSN")

	// If DSN is not set, construct it from individual variables
	if dsn == "" {
		user := os.Getenv("MYSQL_USER")
		password := os.Getenv("MYSQL_PASSWORD")
		host := os.Getenv("MYSQL_HOST")
		port := os.Getenv("MYSQL_PORT")
		dbname := os.Getenv("MYSQL_DBNAME")

		if user != "" && host != "" && port != "" && dbname != "" {
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				user, password, host, port, dbname)
		}
	}

	if dsn == "" {
		return nil, fmt.Errorf("DSN environment variable is not set and could not be constructed from MYSQL_* variables")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_secret_for_dev_only" // Fallback but should be set
	}

	return &Config{
		DSN:       dsn,
		JWTSecret: jwtSecret,
	}, nil
}

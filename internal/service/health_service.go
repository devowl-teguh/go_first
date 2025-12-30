package service

import (
	"context"
	"go_first/internal/models"
	"go_first/internal/repository"
)

// HealthService defines the business logic for health checks
type HealthService interface {
	CheckHealth(ctx context.Context) (*models.HealthStatus, error)
}

type healthService struct {
	repo repository.HealthRepository
}

func NewHealthService(repo repository.HealthRepository) HealthService {
	return &healthService{repo: repo}
}

func (s *healthService) CheckHealth(ctx context.Context) (*models.HealthStatus, error) {
	// Business logic could go here (e.g. checking db connection state)
	return s.repo.GetStatus()
}

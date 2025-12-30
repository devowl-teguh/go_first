package repository

import (
	"go_first/internal/models"
	"time"
)

// HealthRepository defines the interface for health data access (even if just dummy)
type HealthRepository interface {
	GetStatus() (*models.HealthStatus, error)
}

// DummyHealthRepository implements HealthRepository
type DummyHealthRepository struct{}

func NewDummyHealthRepository() *DummyHealthRepository {
	return &DummyHealthRepository{}
}

func (r *DummyHealthRepository) GetStatus() (*models.HealthStatus, error) {
	return &models.HealthStatus{
		Status:    "OK",
		Timestamp: time.Now(),
	}, nil
}

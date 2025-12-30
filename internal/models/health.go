package models

import "time"

// HealthStatus represents the status of the application
type HealthStatus struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

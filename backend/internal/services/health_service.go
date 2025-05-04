package services

import (
	"time"

	"github.com/DB-Vincent/want-to-read/internal/models"
)

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

// CheckHealth performs a health check of the system
func (s *HealthService) CheckHealth() models.Health {
	return models.Health{
		Status:    "ok",
		Timestamp: time.Now().Format(time.RFC3339),
	}
} 
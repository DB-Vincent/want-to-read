package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DB-Vincent/want-to-read/internal/services"
)

type HealthHandler struct {
	healthService *services.HealthService
}

func NewHealthHandler(healthService *services.HealthService) *HealthHandler {
	return &HealthHandler{
		healthService: healthService,
	}
}

// Check handles GET /health requests
// @Summary     Health check endpoint
// @Description Get the health status of the API
// @Tags        health
// @Produce     json
// @Success     200 {object} models.Health
// @Router      /health [get]
func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	health := h.healthService.CheckHealth()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(health); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

package handler

import (
	"encoding/json"
	"go_first/internal/service"
	"net/http"
)

type HealthHandler struct {
	service service.HealthService
}

func NewHealthHandler(service service.HealthService) *HealthHandler {
	return &HealthHandler{service: service}
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	status, err := h.service.CheckHealth(r.Context())
	if err != nil {
		http.Error(w, "Failed to check health", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

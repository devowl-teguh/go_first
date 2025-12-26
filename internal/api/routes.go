package api

import (
	"go_first/internal/handler"
	"net/http"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(mux *http.ServeMux, userHandler *handler.UserHandler) {
	mux.HandleFunc("/users/create", userHandler.CreateUser)
	mux.HandleFunc("/users/get", userHandler.GetUser)
}

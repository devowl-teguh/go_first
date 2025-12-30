import (
	"go_first/internal/handler"
	"go_first/internal/middleware"
	"net/http"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(mux *http.ServeMux, userHandler *handler.UserHandler, loyaltyHandler *handler.LoyaltyHandler, healthHandler *handler.HealthHandler, jwtSecret string) {
	// Public routes
	mux.HandleFunc("/login", userHandler.Login)
	mux.HandleFunc("/users/create", userHandler.CreateUser)
	mux.HandleFunc("/health", healthHandler.HealthCheck)

	// Protected routes
	authMW := middleware.AuthMiddleware(jwtSecret)

	mux.Handle("/users/get", authMW(http.HandlerFunc(userHandler.GetUser)))
	mux.Handle("/loyalty/earn", authMW(http.HandlerFunc(loyaltyHandler.EarnPoints)))
	mux.Handle("/loyalty/redeem", authMW(http.HandlerFunc(loyaltyHandler.RedeemPoints)))
	mux.Handle("/loyalty/balance", authMW(http.HandlerFunc(loyaltyHandler.GetBalance)))
	mux.Handle("/loyalty/history", authMW(http.HandlerFunc(loyaltyHandler.GetHistory)))
}

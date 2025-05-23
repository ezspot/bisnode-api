package routes

import (
	"bisnode/internal/handlers"
	"net/http"
)

// RegisterMotorVehicleRoutes registers the motor vehicle routes
func RegisterMotorVehicleRoutes(mux *http.ServeMux, handler *handlers.MotorVehicleHandler) {
	// Search motor vehicle by license number or VIN
	mux.HandleFunc("GET /api/v1/motor-vehicles/search", handler.Search)
	mux.HandleFunc("POST /api/v1/motor-vehicles/search", handler.Search)
}

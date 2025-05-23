package handlers

import (
	"bisnode/internal/services/bisnode"
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResponse represents a standard error response
// swagger:response errorResponse
type ErrorResponse struct {
	// The error message
	Error string `json:"error"`
}

// SearchRequest represents the request body for searching a motor vehicle
type SearchRequest struct {
	LicenseNumber string `json:"licenseNumber,omitempty"`
	VIN           string `json:"vin,omitempty"`
}

// MotorVehicleHandler handles HTTP requests for motor vehicle information
type MotorVehicleHandler struct {
	service *bisnode.MotorVehicleClient
}

// NewMotorVehicleHandler creates a new MotorVehicleHandler
func NewMotorVehicleHandler(service *bisnode.MotorVehicleClient) *MotorVehicleHandler {
	return &MotorVehicleHandler{
		service: service,
	}
}

// Search handles the search request for motor vehicles (GET)
// @Summary Search for a motor vehicle by license number or VIN
// @Description Search for motor vehicle information using either license number or VIN
// @Tags Motor Vehicles
// @Accept json
// @Produce json
// @Param licenseNumber query string false "License number of the vehicle"
// @Param vin query string false "Vehicle Identification Number"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/motor-vehicles/search [get]
// @Security BasicAuth

// Search handles the search request for motor vehicles (POST)
// @Summary Search for a motor vehicle by license number or VIN (POST)
// @Description Search for motor vehicle information using either license number or VIN with JSON body
// @Tags Motor Vehicles
// @Accept json
// @Produce json
// @Param request body SearchRequest true "Search parameters"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/motor-vehicles/search [post]
// @Security BasicAuth
func (h *MotorVehicleHandler) Search(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Log the incoming request for debugging
	log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
	log.Printf("Headers: %v", r.Header)

	// Get context from request
	ctx := r.Context()
	var request SearchRequest

	if r.Method == http.MethodGet {
		query := r.URL.Query()
		licenseNumber := query.Get("licenseNumber")
		vin := query.Get("vin")

		// Check if either license number or VIN is provided
		if licenseNumber == "" && vin == "" {
			http.Error(w, "Either licenseNumber or VIN must be provided", http.StatusBadRequest)
			return
		}

		var result interface{}
		var err error

		if licenseNumber != "" {
			result, err = h.service.SearchByLicenseNumber(ctx, licenseNumber)
		} else {
			result, err = h.service.SearchByVIN(ctx, vin)
		}

		if err != nil {
			log.Printf("Error searching motor vehicle: %v", err)
			http.Error(w, "Failed to search motor vehicle", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
		return
	}

	// Handle POST request
	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if either license number or VIN is provided
	if request.LicenseNumber == "" && request.VIN == "" {
		http.Error(w, "Either licenseNumber or VIN must be provided in the request body", http.StatusBadRequest)
		return
	}

	var result interface{}
	var err error

	if request.LicenseNumber != "" {
		result, err = h.service.SearchByLicenseNumber(ctx, request.LicenseNumber)
	} else {
		result, err = h.service.SearchByVIN(ctx, request.VIN)
	}
	if err != nil {
		log.Printf("Error searching motor vehicle: %v", err)
		http.Error(w, "Failed to search motor vehicle", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

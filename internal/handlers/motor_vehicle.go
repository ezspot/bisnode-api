package handlers

import (
	"bisnode/internal/services/bisnode"
	"encoding/json"
	"log"
	"net/http"
)

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

// SearchRequest represents the request body for searching a motor vehicle
type SearchRequest struct {
	LicenseNumber string `json:"licenseNumber,omitempty"`
	VIN           string `json:"vin,omitempty"`
}

// Search handles the search request for a motor vehicle
func (h *MotorVehicleHandler) Search(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to %s", r.Method, r.URL.Path)
	log.Printf("Headers: %v", r.Header)
	
	ctx := r.Context()

	// Check if it's a GET request with query parameters
	if r.Method == http.MethodGet {
		query := r.URL.Query()
		licenseNumber := query.Get("licenseNumber")
		vin := query.Get("vin")

		if licenseNumber == "" && vin == "" {
			http.Error(w, "Either licenseNumber or vin query parameter is required", http.StatusBadRequest)
			return
		}

		searchTerm := licenseNumber
		if searchTerm == "" {
			searchTerm = vin
		}

		result, err := h.service.SearchByLicenseNumber(ctx, searchTerm)
		if err != nil {
			log.Printf("Error searching motor vehicle: %v", err)
			http.Error(w, "Failed to search motor vehicle", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
		return
	}

	// Handle POST request with JSON body
	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.LicenseNumber == "" && req.VIN == "" {
		http.Error(w, "Either licenseNumber or vin is required in the request body", http.StatusBadRequest)
		return
	}

	searchTerm := req.LicenseNumber
	if searchTerm == "" {
		searchTerm = req.VIN
	}

	result, err := h.service.SearchByLicenseNumber(ctx, searchTerm)
	if err != nil {
		log.Printf("Error searching motor vehicle: %v", err)
		http.Error(w, "Failed to search motor vehicle", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

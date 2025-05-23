package handlers

import (
	"bisnode/internal/models"
	"bisnode/internal/services/bisnode"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// SearchPersonRequest represents the request body for searching a person
type SearchPersonRequest struct {
	MobileNumber string `json:"mobileNumber"`
}

// SearchOrganizationRequest represents the request body for searching an organization
type SearchOrganizationRequest struct {
	OrganizationNumber string `json:"organizationNumber"`
}

// DirectoryHandler handles HTTP requests for directory search
type DirectoryHandler struct {
	service *bisnode.DirectoryService
}

// NewDirectoryHandler creates a new DirectoryHandler
func NewDirectoryHandler(service *bisnode.DirectoryService) *DirectoryHandler {
	return &DirectoryHandler{
		service: service,
	}
}

// SearchPerson handles the search request for a person by mobile number
// @Summary Search for a person by mobile number
// @Description Search for a person using their mobile number
// @Tags Directory
// @Accept json
// @Produce json
// @Param mobileNumber query string false "Mobile number of the person"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/directory/persons/search [get]
// @Security BasicAuth

// SearchPerson handles the search request for a person by mobile number (POST)
// @Summary Search for a person by mobile number (POST)
// @Description Search for a person using their mobile number with JSON body
// @Tags Directory
// @Accept json
// @Produce json
// @Param request body SearchPersonRequest true "Search parameters"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/directory/persons/search [post]
// @Security BasicAuth
func (h *DirectoryHandler) SearchPerson(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Log the incoming request for debugging
	log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
	log.Printf("Headers: %v", r.Header)

	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req SearchPersonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.MobileNumber == "" {
		respondWithError(w, http.StatusBadRequest, "Mobile number is required")
		return
	}

	result, err := h.service.SearchByMobileNumber(r.Context(), req.MobileNumber)
	if err != nil {
		// Check if the error is due to no results found
		if strings.Contains(err.Error(), "no results") {
			respondWithJSON(w, http.StatusOK, &models.DirectorySearchResponse{
				Result: []models.DirectoryResult{},
			})
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

// SearchOrganization handles the search request for an organization by organization number
// @Summary Search for an organization by organization number
// @Description Search for an organization using its organization number
// @Tags Directory
// @Accept json
// @Produce json
// @Param organizationNumber query string false "Organization number"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/directory/organizations/search [get]
// @Security BasicAuth

// SearchOrganization handles the search request for an organization by organization number (POST)
// @Summary Search for an organization by organization number (POST)
// @Description Search for an organization using its organization number with JSON body
// @Tags Directory
// @Accept json
// @Produce json
// @Param request body SearchOrganizationRequest true "Search parameters"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/directory/organizations/search [post]
// @Security BasicAuth
func (h *DirectoryHandler) SearchOrganization(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Log the incoming request for debugging
	log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
	log.Printf("Headers: %v", r.Header)

	var req SearchOrganizationRequest
	// Try to get the organization number from query parameters first (for backward compatibility)
	orgNo := r.URL.Query().Get("orgNo")
	if orgNo == "" {
		// If not in query params, try to parse from JSON body
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		orgNo = req.OrganizationNumber
	}

	if orgNo == "" {
		respondWithError(w, http.StatusBadRequest, "Organization number is required")
		return
	}

	result, err := h.service.SearchByOrganizationNumber(r.Context(), orgNo)
	if err != nil {
		// Check if the error is due to no results found
		if strings.Contains(err.Error(), "no results") {
			respondWithJSON(w, http.StatusOK, &models.DirectorySearchResponse{
				Result: []models.DirectoryResult{},
			})
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

// respondWithError sends an error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON sends a JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if payload != nil {
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ") // Pretty print JSON
		if err := encoder.Encode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

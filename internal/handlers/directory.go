package handlers

import (
	"bisnode/internal/models"
	"bisnode/internal/services/bisnode"
	"encoding/json"
	"net/http"
	"strings"
)

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

// SearchPerson handles person search by mobile number
func (h *DirectoryHandler) SearchPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SearchRequest
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

// SearchOrganization handles organization search by organization number
func (h *DirectoryHandler) SearchOrganization(w http.ResponseWriter, r *http.Request) {
	var req models.SearchRequest
	
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

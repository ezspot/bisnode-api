package handlers

import (
	"bisnode/internal/models"
	"bisnode/internal/services/bisnode"
	"encoding/json"
	"net/http"
)

// Handler handles HTTP requests
type Handler struct {
	personService *bisnode.PersonService
}

// NewHandler creates a new handler
func NewHandler(personService *bisnode.PersonService) *Handler {
	return &Handler{
		personService: personService,
	}
}

// SearchPerson handles person search requests
func (h *Handler) SearchPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	person, err := h.personService.SearchByMobileNumber(r.Context(), req.MobileNumber)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, person)
}

// Helper functions for JSON responses
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, models.ErrorResponse{
		Error:   http.StatusText(code),
		Message: message,
	})
}

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

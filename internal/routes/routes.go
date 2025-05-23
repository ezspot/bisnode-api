package routes

import (
	"bisnode/internal/handlers"
	"net/http"
)

// NewRouter sets up the HTTP routes
func NewRouter(h *handlers.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	// API v1 routes
	mux.HandleFunc("POST /api/v1/persons/search", h.SearchPerson)

	// Health check endpoint
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return mux
}

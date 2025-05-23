package routes

import (
	"bisnode/internal/handlers"
	"net/http"
)

// RegisterDirectoryRoutes registers all directory search routes
func RegisterDirectoryRoutes(mux *http.ServeMux, h *handlers.DirectoryHandler) {
	// Person search by mobile number
	mux.HandleFunc("POST /api/v1/directory/persons/search", h.SearchPerson)
	mux.HandleFunc("GET /api/v1/directory/persons/search", h.SearchPerson)

	// Organization search by organization number
	mux.HandleFunc("GET /api/v1/directory/organizations/search", h.SearchOrganization)
	mux.HandleFunc("POST /api/v1/directory/organizations/search", h.SearchOrganization)

	// Health check endpoint
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}

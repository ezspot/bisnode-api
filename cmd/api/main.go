package main

import (
	_ "bisnode/docs"
	"bisnode/internal/config"
	"bisnode/internal/handlers"
	"bisnode/internal/routes"
	bisnodeservice "bisnode/internal/services/bisnode"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Bisnode API
// @version         1.0
// @description     A Go service that provides an HTTP API for searching Bisnode data.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	// Initialize configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Bisnode clients
	directoryClient := bisnodeservice.NewDirectoryClient(&cfg.Bisnode)
	motorVehicleClient := bisnodeservice.NewMotorVehicleClient(&cfg.Bisnode)

	// Initialize services
	directoryService := bisnodeservice.NewDirectoryService(directoryClient)

	// Initialize handlers
	directoryHandler := handlers.NewDirectoryHandler(directoryService)
	motorVehicleHandler := handlers.NewMotorVehicleHandler(motorVehicleClient)

	// Setup router
	mux := http.NewServeMux()

	// Register routes
	routes.RegisterDirectoryRoutes(mux, directoryHandler)
	routes.RegisterMotorVehicleRoutes(mux, motorVehicleHandler)

	// Swagger documentation
	docURL := "/swagger/doc.json"
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(docURL), // The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
		httpSwagger.PersistAuthorization(true),
	))

	// Serve the swagger.json file
	mux.HandleFunc(docURL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	router := mux

	// Create HTTP server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}

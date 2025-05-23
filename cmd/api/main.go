package main

import (
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
)

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
	routes.RegisterDirectoryRoutes(mux, directoryHandler)
	routes.RegisterMotorVehicleRoutes(mux, motorVehicleHandler)

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

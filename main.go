package main

import (
	"log"

	"github.com/balu6914/KYC-Match-API/config"
	"github.com/balu6914/KYC-Match-API/handlers"
	"github.com/balu6914/KYC-Match-API/repositories"
	"github.com/balu6914/KYC-Match-API/server"
	"github.com/balu6914/KYC-Match-API/usecases"
	"github.com/joho/godotenv"
)

// main initializes and starts the KYC Match API application
func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, using default environment variables:", err)
	}

	cfg := config.LoadConfig()

	// Initialize repository
	repo, err := repositories.NewHarperDBRepository(cfg)
	if err != nil {
		log.Fatal("failed to initialize repository: ", err)
	}

	// Initialize use case
	useCase := usecases.NewKYCUseCaseImpl(repo)

	// Initialize handler
	handler := handlers.NewKYCHandler(useCase)

	// Initialize and start server
	srv := server.NewEchoServer(handler)
	if err := srv.Start(); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}

package main

import (
	"log"
	"your_project/config"
	"your_project/handlers"
	"your_project/repositories"
	"your_project/server"
	"your_project/usecases"
)

func main() {
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
//
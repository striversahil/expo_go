package main

import (
	"log"
	"mybackend/core/config"
	"mybackend/internal/config"
	"net/http"
	"user-management/internal/api/handlers"
	"user-management/internal/repository"
	"user-management/internal/service"
)

func main() {
    // Load configuration (e.g., database URL, server port)
    cfg := config.LoadConfig()

    // Set up the database repository
    userRepo := repository.NewUserRepository(cfg.DatabaseURL)

    // Set up the service (business logic)
    userService := service.NewUserService(userRepo)

    // Set up the HTTP handlers
    userHandler := handlers.NewUserHandler(userService)

    // Define routes
    http.HandleFunc("/signup", userHandler.Signup)
    http.HandleFunc("/login", userHandler.Login)

    // Start the server
    log.Printf("Server is running on %s...\n", cfg.ServerPort)
    log.Fatal(http.ListenAndServe(cfg.ServerPort, nil))
}
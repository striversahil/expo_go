// This serves as the entry point for the application.

package main

import (
	"log"
	"myapp/core/config"
	"net/http"
	"myapp/core/api/handlers"
	"myapp/core/repository"
	"myapp/core/service"
    "github.com/joho/godotenv"
)

func main() {
    // Load configuration (e.g., database URL, server port)

    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    cfg , err := config.LoadConfig()

    if err != nil {
        log.Fatal(err)
    }

    // Set up the database repository
    userRepo := repository.NewUserRepository(cfg)

    // Set up the service (business logic)
    userService := service.NewUserService(userRepo)

    // Set up the HTTP handlers
    userHandler := handlers.NewUserHandler(userService)

    // Define routes
    http.HandleFunc("/signup", userHandler.RegisterHandler)
    http.HandleFunc("/login", userHandler.Login)

    // Start the server
    log.Printf("Server is running on %s...\n", cfg.DbHost)
    log.Fatal(http.ListenAndServe(cfg.DbHost, nil))
}
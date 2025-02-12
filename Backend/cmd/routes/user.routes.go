package routes

import (
	"myapp/core/api/handlers"
	"myapp/core/config"
	"myapp/core/repository"
	"myapp/core/service"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router , c *config.Config) {
	userRepo := repository.NewUserRepository(c)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/signup", userHandler.RegisterHandler).Methods("POST")
	v1.HandleFunc("/login", userHandler.LoginHandler).Methods("POST")
}
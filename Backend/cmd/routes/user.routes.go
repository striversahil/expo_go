package routes

import (
	"database/sql"
	"myapp/core/api/handlers"
	"myapp/core/config"
	"myapp/core/repository"
	"myapp/core/service"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router , c *config.Config , DB *sql.DB) {
	userRepo := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/signup", userHandler.RegisterHandler).Methods("POST")
	v1.HandleFunc("/login", userHandler.LoginHandler).Methods("POST")
}
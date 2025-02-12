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
	userRepo := repository.NewUserRepository(DB)        //This is like Hydrating the struct of UserRepository with db context value
	userService := service.NewUserService(userRepo)       // Also Hydration of UserService
	userHandler := handlers.NewUserHandler(userService)  // Also Hydration of UserHandler
	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/signup", userHandler.RegisterHandler).Methods("POST")
	v1.HandleFunc("/login", userHandler.LoginHandler).Methods("POST")
}
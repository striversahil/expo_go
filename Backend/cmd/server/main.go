// This serves as the entry point for the application.

package main

import (
	"log"
	"myapp/cmd/database"
	"myapp/cmd/routes"
	"myapp/core/config"
	"net/http"

	"github.com/gorilla/mux"
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


    database.InitDb(cfg)


    // Define routes
    r := mux.NewRouter()

    
    // r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     w.Header().Set("Content-Type", "application/json")
    //     w.Write([]byte(`{"fucking message": "Welcome to myapp from Mux Router", "version": "1.0", "status": "OK"}`))
    //     }).Methods("GET")
        
    routes.UserRoutes(r, cfg , database.DB)


    // Start the server
    log.Printf("Server is running on %s...\n", cfg.ServerHost)
    log.Fatal(http.ListenAndServe(cfg.ServerHost, r))
}
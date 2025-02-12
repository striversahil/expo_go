// This is Connection with Database and Db executable defination goes here
package repository

import (
	"database/sql"
	"fmt"
	// "log"
	"myapp/core/config"

	_ "github.com/lib/pq"
	// "golang.org/x/crypto/bcrypt"
)

type Repository struct {
    db *sql.DB
}

func NewRepository(config *config.Config ) *Repository {
    databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
     config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
	    panic(err)
	}
	return &Repository{db: db}   //It's like referencing back to struct like save this to UserRepository struct for next function Operations
}
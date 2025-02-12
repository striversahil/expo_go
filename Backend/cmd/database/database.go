// This is Connection with Database and Db executable defination goes here
package database

import (
	"database/sql"
	"fmt"
	"log"

	// "log"
	"myapp/core/config"

	_ "github.com/lib/pq"
	// "golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func InitDb(config *config.Config ) {

	
    databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
     config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)
    DB, err := sql.Open("postgres", databaseURL)
    if err != nil {
	    log.Fatal("Error connecting to database:", err)
	}
	schema := `
	CREATE TABLE IF NOT EXISTS Users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT UNIQUE
		password TEXT
	);
	CREATE TABLE IF NOT EXISTS Goals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		Goal Title TEXT,
		Chapter JSONB,
		FOREIGN KEY(user_id) REFERENCES Users(id)
	);`
	DB.Exec(schema)
}
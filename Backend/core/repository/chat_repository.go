// This is Connection with Database and Db executable defination goes here
package repository

import (
	"database/sql"
	"fmt"
	// "log"
	"myapp/core/config"
	"myapp/core/model"

	_ "github.com/lib/pq"
	// "golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(config *config.Config ) *UserRepository {
    databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
     config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        panic(err)
    }
    _ , err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(255), email VARCHAR(255), password VARCHAR(255) , token VARCHAR(255))")
    if err != nil {
        panic(err)
    }
    return &UserRepository{db: db}   //It's like referencing back to struct like save this to UserRepository struct for next function Operations
}


func (r *UserRepository) Save(user *model.User) error {
    // Save user to the database
    _, err := r.db.Exec("INSERT INTO users (name, email, password , token) VALUES ($1, $2, $3 , $4)", user.Name, user.Email, user.Password , user.Token)
    return err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
    // Fetch user by email
    var user model.User
    // log.Default().Println(email , user)
    err := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password , &user.Token)  
     //Save user to the Query to user Model Just reffered
    // log.Default().Println(err)
    return &user, err
}


// func (u *User) HashPassword(password string) error {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	u.PasswordHash = string(hash)
// 	return nil
// }

// func (u *User) CheckPassword(password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
// }

// func CreateUser(username, password string) (*User, error) {
// 	user := &User{Username: username}
// 	err := user.HashPassword(password)
// 	if err != nil {
// 		return nil, err
// 	}

// 	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id`
// 	err = db.QueryRow(query, user.Username, user.PasswordHash).Scan(&user.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// func GetUserByUsername(username string) (*User, error) {
// 	user := &User{}
// 	query := `SELECT id, username, password_hash FROM users WHERE username = $1`
// 	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.PasswordHash)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }
// This is Connection with Database and Db executable defination goes here
package repository

import (
	"database/sql"
	"fmt"
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
    _ , err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(255), email VARCHAR(255), password VARCHAR(255))")
    if err != nil {
        panic(err)
    }
    return &UserRepository{db: db}
}


func (r *UserRepository) UserExist(user *model.User) error {
    _,err := r.db.Exec("SELECT * FROM users WHERE email = $1", user.Email)
    return err
}

func (r *UserRepository) Save(user *model.User) error {
    // Save user to the database
    _, err := r.db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
    return err
}

// func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
//     // Fetch user by email
//     var user domain.User
//     err := r.db.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
//     return &user, err
// }


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
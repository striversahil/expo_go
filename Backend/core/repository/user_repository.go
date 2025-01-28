package repository

import (
	"database/sql"
	"mybackend/core/model"
	// "golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(databaseURL string) *UserRepository {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        panic(err)
    }
    return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *domain.User) error {
    // Save user to the database
    _, err := r.db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
    return err
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
    // Fetch user by email
    var user domain.User
    err := r.db.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
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
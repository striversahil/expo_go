// This is Connection with Database and Db executable defination goes here
package repository

import (
	"database/sql"
	_"log"
	"myapp/core/model"

	_ "github.com/lib/pq"
	// "golang.org/x/crypto/bcrypt"
)

type GoalRepository struct {
    DB *sql.DB
}

func NewGoalRepository( db *sql.DB ) *GoalRepository {
    return &GoalRepository{DB: db}   //It's like referencing back to struct like save this to GoalRepository struct for next function Operations
}


func (r *GoalRepository) CreateGoal(user_id int , goal string , chapters []model.Chapter) error {
    // Save user to the database
    _, err := r.DB.Exec("INSERT INTO goals (user_id , goal, chapters) VALUES ($1, $2, $3 , $4)", user_id, goal, chapters)
    return err
}

func (r *GoalRepository) FindById(id int) (*model.Goal, error) {
    // Fetch user by email
    var user model.User
    // log.Default().Println(email , user)
    err := r.DB.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password , &user.Token)  
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
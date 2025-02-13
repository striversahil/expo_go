// This is Connection with Database and Db executable defination goes here
package repository

import (
	"database/sql"
	"encoding/json"
	_ "log"
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


func (r *GoalRepository) CreateGoal(user_id int , goal string , chapters []string) error {
    // Save user to the database
    chapters_Json , _ := json.Marshal(chapters)

    _, err := r.DB.Exec("INSERT INTO goals (user_id , goal, chapters) VALUES ($1, $2, $3 )", user_id, goal, chapters_Json)
    return err
}

func (r *GoalRepository) FetchGoalById(user_id int) ([]model.Goal, error) {
    // Fetch user by email
    var goals []model.Goal
    // log.Default().Println(email , user)
    rows , err := r.DB.Query("SELECT * FROM goals WHERE user_id = $1", user_id)  
    if err != nil {
        return nil, err
    }
    // Iterating on the rows of Goals
    for rows.Next() {
        var goal model.Goal
        var ChaptersJson string
        err := rows.Scan(&goal.ID, &goal.Goal, &goal.UserID, &ChaptersJson)
        if err != nil {
            return nil, err
        }
        json.Unmarshal([]byte(ChaptersJson), &goal.Chapters)
        goals = append(goals, goal)
    }
    return goals, nil
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
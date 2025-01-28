// This is Where i will write Controller's for my routes which will talk to service

package handlers

import (
	"encoding/json"
	"log"
	"myapp/core/model"
	"myapp/core/service"
	"net/http"
	_ "strings"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
  }

func (uh *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	user_info := &model.User{Name: req.Username, Email: req.Email, Password: req.Password}

	user , err := uh.userService.CreateUser(user_info)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		log.Default().Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	var req struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	user, err := GetUserByUsername(req.Username)
// 	if err != nil {
// 		http.Error(w, "User not found", http.StatusUnauthorized)
// 		return
// 	}

// 	if err := user.CheckPassword(req.Password); err != nil {
// 		http.Error(w, "Invalid password", http.StatusUnauthorized)
// 		return
// 	}

// 	token, err := GenerateJWT(user.ID)
// 	if err != nil {
// 		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"token": token})
// }

// func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
// 	authHeader := r.Header.Get("Authorization")
// 	if authHeader == "" {
// 		http.Error(w, "Missing authorization header", http.StatusUnauthorized)
// 		return
// 	}

// 	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 	token, err := ValidateJWT(tokenString)
// 	if err != nil || !token.Valid {
// 		http.Error(w, "Invalid token", http.StatusUnauthorized)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "You are authenticated"})
// }
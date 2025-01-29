// This is Where i will write Controller's for my routes which will talk to service

package handlers

import (
	// "database/sql"
	"encoding/json"
	// "errors"
	_ "log"
	"myapp/core/model"
	"myapp/core/service"
	"myapp/core/utils"
	"net/http"
	_ "strings"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
  }


var jwtSecret = []byte("secret")



func (uh *UserHandler) RegisterHandler( w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// New Decoder takes what to decode and put's it in the struct by Decode
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if req.Username == "" || req.Email == "" || req.Password == "" {
		utils.ErrorResponse(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	
	user_info := &model.User{Name: req.Username, Email: req.Email, Password: req.Password}
	
	user , err := uh.userService.CreateUser(user_info)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.NewRespose(w, "User created successfully", http.StatusCreated, user)
}

func (uh *UserHandler) LoginHandler( w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	
	user, err := uh.userService.GetUser(&model.User{Email: req.Email})
	if err != nil {
		// Handle other errors (e.g., database issues)
		utils.ErrorResponse(w, "User not found", http.StatusInternalServerError)
		return
	}
	
	// At this point, user is valid and not nil
	// Proceed with login logic (e.g., password check)

	if user.Password != req.Password {
		// http.Error(w, "Invalid password", http.StatusUnauthorized)
		utils.ErrorResponse(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	utils.NewRespose(w, "Login successful", http.StatusOK, user)
}

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
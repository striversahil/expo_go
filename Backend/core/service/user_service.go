// Abstraction of business logic for handlers talking model and repository/db

package service

import (
	"errors"
	_ "log"
	"myapp/core/model"
	"myapp/core/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

var jwtSecret = []byte("secret")

func GenerateJWT(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"name": user.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid signing method")
        }
        return jwtSecret, nil
    })
    if err != nil {
        return nil, err
    }
    return token, nil
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
    // Business logic (e.g., validation)
    // Checking if user already exists
    userexist , err := s.repo.FindByEmail(user.Email) 
    if userexist.Email != "" && err == nil {
        return nil, errors.New("user already exists")
    }

    jwtToken, err := GenerateJWT(user)
    if err != nil {
        return nil, err
    }
    user.Token = jwtToken

    err = s.repo.Save(user)
    if err != nil {
        return nil, err
    }
    return user, nil
}
func (s *UserService) GetUser(user *model.User) (*model.User, error) {
    // Business logic (e.g., validation)
    // Checking if user already exists
    userexist , err := s.repo.FindByEmail(user.Email) 
    if err != nil {
        return nil, errors.New(("user not found"))
    }

    return userexist, nil
}


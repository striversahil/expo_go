// Abstraction of business logic for handlers talking model and repository/db

package service

import (
	"errors"
	_"log"
	"myapp/core/model"
	"myapp/core/repository"
)

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
    // Business logic (e.g., validation)
    // Checking if user already exists
    userexist , err := s.repo.FindByEmail(user.Email) 
    if userexist.Email != "" && err == nil {
        return nil, errors.New("user already exists")
    }

    err = s.repo.Save(user)
    if err != nil {
        return nil, err
    }
    return user, nil
}


func (s *UserService) CheckPassword(req struct{ Username, Password string }) error {
    // Business logic (e.g., validation)

    return nil
}

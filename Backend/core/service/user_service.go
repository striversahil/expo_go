// Abstraction of business logic for handlers talking model and repository/db

package service

import (
    "myapp/core/model"
    "myapp/core/repository"
)

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
    // Business logic (e.g., validation)
    err := s.repo.Save(user)
    if err != nil {
        return nil, err
    }
    return user, nil
}

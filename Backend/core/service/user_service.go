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

func (s *UserService) CreateUser(user *model.User) error {
    // Business logic (e.g., validation)
    return s.repo.Save(user)
}
// Abstraction of business logic for handlers talking model and repository/db

package service

import (
    "myapp/internal/models"
    "myapp/core/repository"
)

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
    // Business logic (e.g., validation)
    return s.repo.Create(user)
}
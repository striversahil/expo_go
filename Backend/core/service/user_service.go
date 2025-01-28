package service

import (
	"mybackend/core/model"
	"mybackend/core/repository"
)

type UserService struct {
    userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
    return s.userRepo.GetAllUsers()
}
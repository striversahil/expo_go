package service

import (
	"myapp/core/repository"
	"myapp/core/utils"
)

type GoalService struct {
	repo *repository.GoalRepository
}

func NewGoalService(repo *repository.GoalRepository) *GoalService {
	return &GoalService{repo: repo}
}


func (s *GoalService) CreateGoal(user_id int , goal string) error {

	utils.
    // Save user to the database
    return s.repo.CreateGoal(user_id , goal , chapters)
}
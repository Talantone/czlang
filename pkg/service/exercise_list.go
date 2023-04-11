package service

import (
	czlang "awesomeProject"
	"awesomeProject/pkg/repository"
)

type ExerciseListService struct {
	repo repository.ExerciseList
}

func NewExerciseListService(repo repository.ExerciseList) *ExerciseListService {
	return &ExerciseListService{repo: repo}
}

func (s *ExerciseListService) Create(userId int, exerciseList czlang.Exercise) (int, error) {
	return s.repo.Create(userId, exerciseList)
}

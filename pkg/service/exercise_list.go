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

func (s *ExerciseListService) GetAll(userId int) ([]czlang.Exercise, error) {
	return s.repo.GetAll(userId)
}

func (s *ExerciseListService) GetById(userId, listId int) (czlang.Exercise, error) {
	return s.repo.GetById(userId, listId)
}

func (s *ExerciseListService) Update(userId int, listId int, input czlang.UpdateExerciseInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}

func (s *ExerciseListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

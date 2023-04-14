package service

import (
	czlang "awesomeProject"
	"awesomeProject/pkg/repository"
)

type ExerciseItemService struct {
	repo     repository.ExerciseItem
	listRepo repository.ExerciseList
}

func NewExerciseItemService(repo repository.ExerciseItem, listRepo repository.ExerciseList) *ExerciseItemService {
	return &ExerciseItemService{repo: repo, listRepo: listRepo}
}

func (s *ExerciseItemService) Create(userId int, listId int, exerciseItem czlang.ExerciseItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, exerciseItem)
}

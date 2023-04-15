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

func (s *ExerciseItemService) GetAll(userId, listId int) ([]czlang.ExerciseItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *ExerciseItemService) GetById(userId, itemId int) (czlang.ExerciseItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *ExerciseItemService) Update(userId int, itemId int, input czlang.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, itemId, input)
}

func (s *ExerciseItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

package service

import (
	czlang "awesomeProject"
	"awesomeProject/pkg/repository"
)

type Authorization interface {
	CreateUser(user czlang.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ExerciseList interface {
	Create(userId int, exerciseList czlang.Exercise) (int, error)
	GetAll(userId int) ([]czlang.Exercise, error)
	GetById(userId, listId int) (czlang.Exercise, error)
	Update(userId int, listId int, input czlang.UpdateExerciseInput) error
	Delete(userId, listId int) error
}

type ExerciseItem interface {
	Create(userId int, listId int, exerciseItem czlang.ExerciseItem) (int, error)
	GetAll(userId, listId int) ([]czlang.ExerciseItem, error)
	GetById(userId, itemId int) (czlang.ExerciseItem, error)
	Update(userId int, itemId int, input czlang.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Service struct {
	Authorization
	ExerciseList
	ExerciseItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		ExerciseList:  NewExerciseListService(repo.ExerciseList),
		ExerciseItem:  NewExerciseItemService(repo.ExerciseItem, repo.ExerciseList),
	}
}

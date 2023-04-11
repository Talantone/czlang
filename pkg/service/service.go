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
}

type ExerciseItem interface {
}

type Service struct {
	Authorization
	ExerciseList
	ExerciseItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}

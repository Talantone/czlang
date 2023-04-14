package repository

import (
	czlang "awesomeProject"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user czlang.User) (int, error)
	GetUser(username, password string) (czlang.User, error)
}

type ExerciseList interface {
	Create(userId int, exerciseList czlang.Exercise) (int, error)
	GetAll(userId int) ([]czlang.Exercise, error)
	GetById(userId, listId int) (czlang.Exercise, error)
	Update(userId int, listId int, input czlang.UpdateExerciseInput) error
	Delete(userId, listId int) error
}

type ExerciseItem interface {
	Create(listId int, exerciseItem czlang.ExerciseItem) (int, error)
}

type Repository struct {
	Authorization
	ExerciseList
	ExerciseItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ExerciseList:  NewExerciseListPostgres(db),
		ExerciseItem:  NewExerciseItemPostgres(db),
	}
}

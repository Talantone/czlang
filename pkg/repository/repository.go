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
}

type ExerciseItem interface {
}

type Repository struct {
	Authorization
	ExerciseList
	ExerciseItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

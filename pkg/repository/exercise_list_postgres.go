package repository

import (
	czlang "awesomeProject"
	"github.com/jmoiron/sqlx"
)

type ExerciseListPostgres struct {
	db *sqlx.DB
}

func NewExerciseListPostgres(db *sqlx.DB) *ExerciseListPostgres {
	return &ExerciseListPostgres{db: db}
}

func (r *ExerciseListPostgres) Create(userId int, exerciseList czlang.Exercise) (int, error) {
	return 0, nil
}

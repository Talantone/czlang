package repository

import (
	czlang "awesomeProject"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ExerciseListPostgres struct {
	db *sqlx.DB
}

func NewExerciseListPostgres(db *sqlx.DB) *ExerciseListPostgres {
	return &ExerciseListPostgres{db: db}
}

func (r *ExerciseListPostgres) Create(userId int, exerciseList czlang.Exercise) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", exerciseListTable)
	row := tx.QueryRow(createListQuery, exerciseList.Title, exerciseList.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) values ($1, 2) RETURNING id", usersListTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

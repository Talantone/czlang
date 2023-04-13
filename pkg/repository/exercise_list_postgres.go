package repository

import (
	czlang "awesomeProject"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) values ($1, $2) RETURNING id", usersListTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *ExerciseListPostgres) GetAll(userId int) ([]czlang.Exercise, error) {
	var lists []czlang.Exercise
	query := fmt.Sprintf("SELECT el.id, el.title, el.description FROM %s el INNER JOIN %s ul on el.id = ul.list_id WHERE ul.user_id = $1", exerciseListTable, usersListTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *ExerciseListPostgres) GetById(userId, listId int) (czlang.Exercise, error) {
	var exercise czlang.Exercise
	query := fmt.Sprintf(`SELECT el.id, el.title, el.description FROM %s el
								INNER JOIN %s ul on el.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		exerciseListTable, usersListTable)
	err := r.db.Get(&exercise, query, userId, listId)
	return exercise, err
}

func (r *ExerciseListPostgres) Update(userId int, listId int, input czlang.UpdateExerciseInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$1%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$1%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s el SET %s FROM %s ul WHERE el.id = ul.list_id AND ul.user_id=$%d AND ul.list_id=$%d",
		exerciseListTable, setQuery, usersListTable, argId, argId+1)
	args = append(args, listId, userId)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ExerciseListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s el USING %s ul WHERE el.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2", exerciseListTable, usersListTable)
	_, err := r.db.Exec(query, userId, listId)
	return err
}

package repository

import (
	czlang "awesomeProject"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ExerciseItemPostgres struct {
	db *sqlx.DB
}

func NewExerciseItemPostgres(db *sqlx.DB) *ExerciseItemPostgres {
	return &ExerciseItemPostgres{db: db}
}

func (r *ExerciseItemPostgres) Create(listId int, exerciseItem czlang.ExerciseItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", exerciseItemTable)

	row := tx.QueryRow(createItemQuery, exerciseItem.Title, exerciseItem.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

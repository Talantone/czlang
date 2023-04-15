package repository

import (
	czlang "awesomeProject"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (r *ExerciseItemPostgres) GetAll(userId int, listId int) ([]czlang.ExerciseItem, error) {
	var items []czlang.ExerciseItem
	query := fmt.Sprintf("SELECT * FROM %s ei INNER JOIN %s li on li.item_id = ei.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2",
		exerciseItemTable, listsItemsTable, usersListTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ExerciseItemPostgres) GetById(userId, itemId int) (czlang.ExerciseItem, error) {
	var item czlang.ExerciseItem
	query := fmt.Sprintf("SELECT ei.id, ei.title, ei.description, ei.done FROM %s ei INNER JOIN %s li on li.item_id = ei.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE ei.id = $2 AND ul.user_id = $1",
		exerciseItemTable, listsItemsTable, usersListTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *ExerciseItemPostgres) Update(userId int, itemId int, input czlang.UpdateItemInput) error {
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
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$1%d", argId))
		args = append(args, *input.Done)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s ei SET %s FROM %s li, %s ul WHERE ei.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ei.id = $%d",
		exerciseItemTable, setQuery, listsItemsTable, usersListTable, argId, argId+1)
	args = append(args, userId, itemId)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ExerciseItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ei USING %s li, %s ul 
									WHERE ei.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ei.id = $2`,
		exerciseItemTable, listsItemsTable, usersListTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

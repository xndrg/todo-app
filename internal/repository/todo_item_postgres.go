package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	todo "github.com/xndrg/crud-app"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{
		db: db,
	}
}

func (r *TodoItemPostgres) CreateItem(listID int64, item todo.TodoItem) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemID int64
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err = row.Scan(&itemID); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listID, itemID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemID, tx.Commit()
}

func (r *TodoItemPostgres) GetAllItems(userID int64, listID int64) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.is_done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
				INNER JOIN %s ul on ul.list_id = li.list_id
				WHERE li.list_id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable,
	)

	if err := r.db.Select(&items, query, listID, userID); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) GetItemByID(userID int64, itemID int64) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.is_done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
				INNER JOIN %s ul on ul.list_id = li.list_id
				WHERE ti.id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable,
	)

	if err := r.db.Get(&item, query, itemID, userID); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemPostgres) UpdateByID(userID int64, itemID int64, input todo.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argID))
		args = append(args, *input.Description)
		argID++
	}

	if input.IsDone != nil {
		setValues = append(setValues, fmt.Sprintf("is_done = $%d", argID))
		args = append(args, *input.IsDone)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
				WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemsTable,
		setQuery,
		listsItemsTable,
		usersListsTable,
		argID,
		argID+1,
	)
	args = append(args, userID, itemID)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (r *TodoItemPostgres) DeleteItemByID(userID int64, itemID int64) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul
				WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable,
	)
	res, err := r.db.Exec(query, userID, itemID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

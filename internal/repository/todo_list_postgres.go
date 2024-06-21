package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	todo "github.com/xndrg/crud-app"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) CreateTodoList(userID int64, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listID int
	createListQuery := fmt.Sprintf(
		"INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id",
		todoListsTable,
	)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&listID); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf(
		"INSERT INTO %s (user_id, list_id) VALUES ($1, $2)",
		usersListsTable,
	)
	_, err = tx.Exec(createUsersListQuery, userID, listID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listID, tx.Commit()
}

func (r *TodoListPostgres) GetAllLists(userID int64) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable,
		usersListsTable,
	)
	err := r.db.Select(&lists, query, userID)

	return lists, err
}

func (r *TodoListPostgres) GetListByID(userID int64, listID int64) (todo.TodoList, error) {
	var list todo.TodoList
	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable,
		usersListsTable,
	)
	err := r.db.Get(&list, query, userID, listID)

	return list, err
}

func (r *TodoListPostgres) UpdateByID(userID int64, listID int64, input todo.UpdateListInput) error {
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

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		"UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d",
		todoListsTable,
		setQuery,
		usersListsTable,
		argID,
		argID+1,
	)
	args = append(args, listID, userID)

	logrus.Debugf("updateQuery: %s\n", query)
	logrus.Debugf("args: %v\n", args)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *TodoListPostgres) DeleteListByID(userID int64, listID int64) error {
	query := fmt.Sprintf(
		"DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable,
		usersListsTable,
	)
	_, err := r.db.Exec(query, userID, listID)

	return err
}

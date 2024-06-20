package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
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

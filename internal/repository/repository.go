package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/xndrg/crud-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int64, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	CreateTodoList(userID int64, list todo.TodoList) (int, error)
	GetAllLists(userID int64) ([]todo.TodoList, error)
	GetListByID(userID int64, listID int64) (todo.TodoList, error)
	UpdateByID(userID int64, listID int64, input todo.UpdateListInput) error
	DeleteListByID(userID int64, listID int64) error
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}

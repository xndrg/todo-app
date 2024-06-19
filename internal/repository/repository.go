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
	}
}

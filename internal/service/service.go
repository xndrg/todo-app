package service

import (
	todo "github.com/xndrg/crud-app"
	"github.com/xndrg/crud-app/internal/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int64, error)
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

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}

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
	}
}

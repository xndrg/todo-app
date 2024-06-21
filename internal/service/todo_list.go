package service

import (
	todo "github.com/xndrg/crud-app"
	"github.com/xndrg/crud-app/internal/repository"
)

type TodoListService struct {
	repos repository.TodoList
}

func NewTodoListService(repos repository.TodoList) *TodoListService {
	return &TodoListService{repos: repos}
}

func (s *TodoListService) CreateTodoList(userID int64, list todo.TodoList) (int, error) {
	return s.repos.CreateTodoList(userID, list)
}

func (s *TodoListService) GetAllLists(userID int64) ([]todo.TodoList, error) {
	return s.repos.GetAllLists(userID)
}

func (s *TodoListService) GetListByID(userID int64, listID int64) (todo.TodoList, error) {
	return s.repos.GetListByID(userID, listID)
}

func (s *TodoListService) UpdateByID(userID int64, listID int64, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repos.UpdateByID(userID, listID, input)
}

func (s *TodoListService) DeleteListByID(userID int64, listID int64) error {
	return s.repos.DeleteListByID(userID, listID)
}

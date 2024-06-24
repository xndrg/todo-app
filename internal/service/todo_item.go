package service

import (
	todo "github.com/xndrg/crud-app"
	"github.com/xndrg/crud-app/internal/repository"
)

type TodoItemService struct {
	itemRepos repository.TodoItem
	listRepos repository.TodoList
}

func NewTodoItemService(itemRepos repository.TodoItem, listRepos repository.TodoList) *TodoItemService {
	return &TodoItemService{
		itemRepos: itemRepos,
		listRepos: listRepos,
	}
}

func (s *TodoItemService) CreateItem(userID int64, listID int64, item todo.TodoItem) (int64, error) {
	_, err := s.listRepos.GetListByID(userID, listID)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	return s.itemRepos.CreateItem(listID, item)
}

func (s *TodoItemService) GetAllItems(userID int64, listID int64) ([]todo.TodoItem, error) {
	return s.itemRepos.GetAllItems(userID, listID)
}

func (s *TodoItemService) GetItemByID(userID int64, itemID int64) (todo.TodoItem, error) {
	return s.itemRepos.GetItemByID(userID, itemID)
}

func (s *TodoItemService) UpdateByID(userID int64, itemID int64, input todo.UpdateItemInput) error {
	return s.itemRepos.UpdateByID(userID, itemID, input)
}

func (s *TodoItemService) DeleteItemByID(userID int64, itemID int64) error {
	return s.itemRepos.DeleteItemByID(userID, itemID)
}

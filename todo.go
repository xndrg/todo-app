package todo

import "errors"

type TodoList struct {
	ID          int64  `json:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	ID     int64
	UserID int64
	ListID int64
}

type TodoItem struct {
	ID          int64  `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	IsDone      bool   `json:"is_done" db:"is_done"`
}

type ListsItem struct {
	ID     int64
	ListID int64
	ItemID int64
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	IsDone      *bool   `json:"isDone"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.IsDone == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

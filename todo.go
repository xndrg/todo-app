package todo

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
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

type ListsItem struct {
	ID     int64
	ListID int64
	ItemID int64
}

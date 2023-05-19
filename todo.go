// Package todo has general models and methods.
package todo

// TodoList is list of tasks.
type TodoList struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UsersList is list of users.
type UsersList struct {
	ID     int
	UserID int
	ListID int
}

// TodoItem is item for todo list.
type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// ListsItem is list of todo items.
type ListsItem struct {
	ID     int
	ListID int
	ItemID int
}

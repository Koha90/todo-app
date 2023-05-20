// Package repository has repository for application.
package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/koha90/todo-app"
)

// Authorization has methods for authorization in application todo-app.
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

// TodoList has methods for todo list in application todo-app.
type TodoList interface{}

// TodoItem has methods for todo item.
type TodoItem interface{}

// Repository has interfaces for application.
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewService constructor of service application.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

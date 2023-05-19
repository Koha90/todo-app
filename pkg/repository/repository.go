// Package repository has repository for application.
package repository

import "github.com/jmoiron/sqlx"

// Authorization has methods for authorization in application todo-app.
type Authorization interface{}

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
	return &Repository{}
}

// Package service has services for application.
package service

import (
	"github.com/koha90/todo-app"
	"github.com/koha90/todo-app/pkg/repository"
)

// Authorization has methods for authorization in application todo-app.
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

// TodoList has methods for todo list in application todo-app.
type TodoList interface{}

// TodoItem has methods for todo item.
type TodoItem interface{}

// Service has interfaces for application.
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService constructor of service application.
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

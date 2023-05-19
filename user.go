// Package todo has general models and methods.
package todo

// User is model for database.
type User struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

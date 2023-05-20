// Package todo has general models and methods.
package todo

// User is model for database.
type User struct {
	ID       int    `json:"-"        db:"id"`
	Name     string `json:"name"             binding:"required"`
	Username string `json:"username"         binding:"required"`
	Password string `json:"password"         binding:"required"`
}

package repositories

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

// NewUserRepository creates a new users repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create inset user on db
func (u Users) Create(user models.User) (uint64, error) {
	return 0, nil
}
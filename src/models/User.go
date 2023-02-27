package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user using the app
type User struct {
	ID uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Nick string  `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("Name is required!")
	}

	if user.Nick == "" {
		return errors.New("Nick is required!")
	}

	if user.Email == "" {
		return errors.New("Email is required!")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Email format is invalid!")
	}

	if stage == "create" && user.Password == "" {
		return errors.New("Password is required!")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
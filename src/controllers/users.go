package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateUser creates user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body) 
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUsersRepository(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("User created sucessfully with ID %d!", userID)))
}

// GetAllUsers get all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users!"))
}

// GetUserById get user by id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user by id!"))
}

// UpdateUser update user info
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user!"))
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User!"))
}

package controllers

import "net/http"

// CreateUser creates user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user!"))
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

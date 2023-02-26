package controllers

import "net/http"

// CreateUser creates user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user!"))
}


func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users!"))
}


func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user by id!"))
}


func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user!"))
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User!"))
}

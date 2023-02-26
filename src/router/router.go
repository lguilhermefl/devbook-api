package router

import "github.com/gorilla/mux"

// Generate returns a router with api routes
func Generate() *mux.Router {
	return mux.NewRouter()
}
package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a router with api routes
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Config(r)
}
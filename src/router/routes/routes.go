package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all api routes
type Route struct {
	URI string
	Method string
	Function func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// Config puts all routes in router
func Config(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r;
}
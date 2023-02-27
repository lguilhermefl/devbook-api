package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route {
	{
		URI: "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Function: controllers.GetUsersByNameOrNick,
		RequireAuth: false,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodGet,
		Function: controllers.GetUserById,
		RequireAuth: false,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodPut,
		Function: controllers.UpdateUserById,
		RequireAuth: false,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		RequireAuth: false,
	},
}
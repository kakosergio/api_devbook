package routes

import (
	users "api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI: "/users",
		Method: http.MethodPost,
		Handler: users.Create,
		HasAuth: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Handler: users.FindAll,
		HasAuth: true,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodGet,
		Handler: users.FindById,
		HasAuth: true,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodPut,
		Handler: users.Update,
		HasAuth: true,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodDelete,
		Handler: users.Delete,
		HasAuth: true,
	},
	{
		URI: "/users/{id}/follow",
		Method: http.MethodPost,
		Handler: users.Follow,
		HasAuth: true,
	},
}
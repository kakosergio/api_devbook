package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI: "/users",
		Method: http.MethodPost,
		Handler: controllers.CreateUser,
		HasAuth: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Handler: controllers.FindAllUsers,
		HasAuth: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodGet,
		Handler: controllers.FindUserById,
		HasAuth: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdateUser,
		HasAuth: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteUser,
		HasAuth: false,
	},
}
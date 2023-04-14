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
	{
		URI: "/users/{id}/unfollow",
		Method: http.MethodDelete,
		Handler: users.Unfollow,
		HasAuth: true,
	},
	{
		URI: "/users/{id}/followers",
		Method: http.MethodGet,
		Handler: users.FindFollowers,
		HasAuth: true,
	},
	{
		URI: "/users/{id}/following",
		Method: http.MethodGet,
		Handler: users.FindFollowing,
		HasAuth: true,
	},
}
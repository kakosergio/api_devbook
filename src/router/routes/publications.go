package routes

import (
	"api/src/controllers"
	"net/http"
)

var publicationsRoutes = []Route{
	{
		URI: "/pub",
		Method: http.MethodPost,
		Handler: controllers.CreatePub,
		HasAuth: true,
	},
	{
		URI: "/pub",
		Method: http.MethodGet,
		Handler: controllers.FindPubs,
		HasAuth: true,
	},
	{
		URI: "/pub/{id}",
		Method: http.MethodGet,
		Handler: controllers.FindPub,
		HasAuth: true,
	},
	{
		URI: "/pub/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdatePub,
		HasAuth: true,
	},
	{
		URI: "/pub/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeletePub,
		HasAuth: true,
	},
	{
		URI: "/users/{id}/pub",
		Method: http.MethodGet,
		Handler: controllers.FindPubByUser,
		HasAuth: true,
	},
	{
		URI: "/pub/{id}/like",
		Method: http.MethodPost,
		Handler: controllers.LikePub,
		HasAuth: true,
	},
}
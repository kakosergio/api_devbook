package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Esse struct é um blueprint de todas as rotas da API
type Route struct {
	URI string
	Method string
	Handler func(http.ResponseWriter, *http.Request)
	HasAuth bool
}

// Configura as rotas de um jeito muito legal, como se fosse uma fabriquinha de rotas
func SetRoutes(r *mux.Router) *mux.Router{
	// pega uma cópia das userRoutes
	routes := userRoutes

	// itera sobre a userRoutes, configurando cada uma das rotas
	for _, route := range routes{
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	// retorna todas as rotas configuradas
	return r
}
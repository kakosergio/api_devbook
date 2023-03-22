package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um router com as rotas configuradas
func Gerar() *mux.Router{
	// cria um novo router
	r := mux.NewRouter()
	// retorna a rota configurada pela função SetRoutes
	return routes.SetRoutes(r)
}
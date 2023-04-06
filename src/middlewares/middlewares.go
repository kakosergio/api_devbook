package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"log"
	"net/http"
	"time"
)

// Logger escreve informações da requisição no terminal
func Logger (next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		log.Printf("\n%s %s [%s] %s", time.Now(), r.Host, r.Method, r.URL.Path)
		next(w, r)
	}
}

// Authenticate autentica o usuário que está sendo logado
func Authenticate (next http.HandlerFunc) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		if err := auth.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
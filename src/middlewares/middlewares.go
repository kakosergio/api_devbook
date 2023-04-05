package middlewares

import (
	"fmt"
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

func Authenticate (next http.HandlerFunc) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		fmt.Println("Authenticating...")
		next(w, r)
	}
}
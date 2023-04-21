package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"log"
	"net/http"

)

// statusRecorder cria um struct que vai guardar as informações do responsewriter para passar para o próximo handler e gravar o statuscode
type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader é uma função auxiliar que grava o statuscode no header do responsewriter. Ela implementa a interface http.ResponseWriter
// Isso significa que será chamada toda vez que o responseWriter do struct statusRecorder for solicitado, no lugar do WriteHeader do handler original.
func (recorder *statusRecorder) WriteHeader (code int){
	recorder.statusCode = code
	recorder.ResponseWriter.WriteHeader(code)
}

// Logger escreve informações da requisição no terminal
func Logger (next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		// Cria um statusCode padrão para caso o WriteHeader não for chamado.
		rec := statusRecorder{w, 200}

		// Passa o controle para o próximo handler
		next.ServeHTTP(&rec, r)
		// Printa na tela o log com as informaçõe obtidas da request e do writer
		log.Printf("%s%s [%d] %s", r.Host, r.URL.Path, rec.statusCode, r.Method)
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
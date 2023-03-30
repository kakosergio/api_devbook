package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta em JSON para requisição
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// testa se retorna algum dado na resposta do servidor, para o caso de métodos como o PUT que não retorna corpo
	if data != nil {

		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	},
	)
}

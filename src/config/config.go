package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	StringConnDB = ""
	Port         = ""
)

// Carregar vai inicializar as variáveis de ambiente
func Load() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	Port = os.Getenv("API_PORT")

	StringConnDB = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("API_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_NAME"),
	)
}

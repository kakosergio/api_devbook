package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/lib/pq" // Driver
)

// Connect abre a conexão com o banco de dados e a retorna
func Connect() (*sql.DB, error) {
	// Abre a conexão
	db, err := sql.Open("postgres", config.StringConnDB)
	// se der erro, retorna
	if err != nil {
		return nil, err
	}
	// se der ping e apresentar erro, retorna o erro
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	// se tudo estiver certo, retorna a conexão com o db
	return db, nil
}

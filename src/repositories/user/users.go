package repositories

import (
	"api/src/models"
	"database/sql"
)

// Struct que recebe apenas a conexão com o banco de dados para manipulação
type user struct {
	db *sql.DB
}

// UsersRepository inicia um novo repositório para manipulação do banco de dados
func UsersRepository(db *sql.DB) *user {
	return &user{db}
}

// Create insere um usuário no banco de dados
func (u user) Create(user models.User) (uint64, error){
	return 0, nil
}
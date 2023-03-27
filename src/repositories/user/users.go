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
func (repositorio user) Create(user models.User) (uint64, error) {
	statement, err := repositorio.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var id uint64
	err = statement.QueryRow(user.Name, user.Nick, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err

}

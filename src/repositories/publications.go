package repositories

import (
	"api/src/models"
	"database/sql"
)

// Publications representa o repositório de publicações
type Publications struct {
	db *sql.DB
}

// PublicationsRepository cria um repositório de publications
func PublicationsRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

// Cria uma nova publicação no banco de dados, retornando seu ID
func (repository Publications) Create(pub models.Publication) (uint64, error){
	statement, err := repository.db.Prepare("INSERT INTO publications (title, body, author_id) VALUES ($1, $2, $3) RETURNING id")

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var id uint64
	err = statement.QueryRow(pub.Title, pub.Body, pub.AuthorId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
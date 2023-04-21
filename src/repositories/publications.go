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

// FindById devolve uma publicação através de seu ID
func (repository Publications) FindById(id uint64) (models.Publication, error){
	rows, err := repository.db.Query("SELECT p.*, u.nick FROM publications p INNER JOIN users u ON u.id = p.author_id WHERE p.id = $1", id)

	if err != nil {
		return models.Publication{}, err
	}
	defer rows.Close()

	var publication models.Publication
	if rows.Next(){
		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Body,
			&publication.AuthorId,
			&publication.Likes,
			&publication.CreatedOn,
			&publication.AuthorNick,
		); err != nil {
			return models.Publication{}, nil
		}
	}
	return publication, nil
}
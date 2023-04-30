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
func (repository Publications) Create(pub models.Publication) (uint64, error) {
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
func (repository Publications) FindById(id uint64) (models.Publication, error) {
	rows, err := repository.db.Query("SELECT p.*, u.nick FROM publications p INNER JOIN users u ON u.id = p.author_id WHERE p.id = $1", id)

	if err != nil {
		return models.Publication{}, err
	}
	defer rows.Close()

	var publication models.Publication
	if rows.Next() {
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

// FindPubs busca publicações dos usuários seguidos por quem a requisitou
func (repository Publications) FindPubs(userId uint64) ([]models.Publication, error) {
	rows, err := repository.db.Query(
		`SELECT DISTINCT p.*, u.nick FROM publications p 
		INNER JOIN users u ON u.id = p.author_id 
		INNER JOIN followers f ON p.author_id = f.user_id 
		WHERE u.id = $1 OR f.follower_id = $2 ORDER BY 1 DESC`, userId, userId)

	if err != nil {
		return []models.Publication{}, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication
		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Body,
			&publication.AuthorId,
			&publication.Likes,
			&publication.CreatedOn,
			&publication.AuthorNick,
		); err != nil {
			return []models.Publication{}, err
		}
		publications = append(publications, publication)
	}
	return publications, nil
}

// Update modifica uma publicação do usuário
func (repository Publications) Update(pubID uint64, publication models.Publication) error {
	statement, err := repository.db.Prepare("UPDATE publications SET title = $1, body = $2 WHERE id = $3")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Body, pubID); err != nil {
		return err
	}
	return nil
}

// Delete apaga uma publicação de um determinado usuário
func (repository Publications) Delete(pubID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM publications WHERE id = $1")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(pubID); err != nil {
		return err
	}
	return nil
}

// FindByUser traz todas as publicações de um usuário específico
func (repository Publications) FindByUser(userId uint64) ([]models.Publication, error) {
	rows, err := repository.db.Query(`SELECT p.*, u.nick FROM publications p JOIN users u ON u.id = p.author_id WHERE p.author_id = $1`, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Body,
			&publication.AuthorId,
			&publication.Likes,
			&publication.CreatedOn,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}
		publications = append(publications, publication)
	}
	return publications, nil
}

// Like curte uma publicação de um determinado usuário
func (repository Publications) Like (pubId uint64) error {
	statement, err := repository.db.Prepare("UPDATE publications SET likes = likes + 1 WHERE id = $1")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(pubId); err != nil {
		return err
	}
	return nil
}

// Unlike descurte uma publicação de um determinado usuário
func (repository Publications) Unlike (pubId uint64) error {
	statement, err := repository.db.Prepare(`
	UPDATE publications SET likes = 
		CASE 
			WHEN likes > 0 THEN likes - 1 
			ELSE likes 
		END 
	WHERE id = $1`)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(pubId); err != nil {
		return err
	}
	return nil
}
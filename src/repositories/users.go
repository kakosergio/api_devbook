package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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
func (repository user) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES ($1, $2, $3, $4) RETURNING id")
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

// Find retorna os usuários com o nick ou nome requisitado
func (repository user) Find(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	rows, err := repository.db.Query(
		"SELECT id, name, nick, email, createdOn FROM users WHERE name LIKE $1 OR nick LIKE $2", nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next(){
		var user models.User
		if err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedOn,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// FindById traz um usuário do banco de dados através de seu número de Id
func (repository user) FindById(userId uint64) (models.User, error){
	rows, err := repository.db.Query(
		"SELECT id, name, nick, email, createdOn FROM users WHERE id = $1", userId,
	)
	if err != nil{
		return models.User{}, err
	}
	defer rows.Close()
	var user models.User
	if rows.Next(){
		if err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedOn,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// Update altera as informações de um usuário no banco de dados
func (repository user) Update(userId uint64, user models.User) error{
	statement, err := repository.db.Prepare("UPDATE users SET name = $1, nick = $2, email = $3 WHERE id = $4")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Nick, user.Email, userId); err != nil {
		return err
	}
	return nil

}

// Delete exclui um usuário do banco de dados
func (repository user) Delete(userId uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userId); err != nil {
		return err
	}
	return nil
}

// FindByEmail busca no banco de dados um usuário a partir de seu e-mail cadastrado
func (repository user) FindByEmail(email string) (models.User, error){
	row, err := repository.db.Query("SELECT id, password FROM users WHERE email = $1", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()
	
	var user models.User
	if row.Next(){
		if err = row.Scan(&user.Id, &user.Password); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}
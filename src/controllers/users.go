package controllers

import (
	"api/src/database"
	"api/src/models"
	repositories "api/src/repositories/user"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateUser cria um usuário e insere no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request){
	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}
	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.UsersRepository(db)
	id, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("Usuário inserido. ID: %d", id)))
}

// FindAllUsers busca no banco de dados todos os usuários cadastrados
func FindAllUsers(w http.ResponseWriter, r *http.Request){

}

// FindUserById busca um usuário pelo seu id
func FindUserById(w http.ResponseWriter, r *http.Request){

}

// UpdateUser atualiza as informações de um usuário pelo seu id
func UpdateUser(w http.ResponseWriter, r *http.Request){

}

// DeleteUser apaga um usuário do banco de dados
func DeleteUser(w http.ResponseWriter, r *http.Request){

}

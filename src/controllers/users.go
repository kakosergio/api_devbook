package controllers

import (
	"api/src/database"
	"api/src/models"
	repositories "api/src/repositories/user"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

// CreateUser cria um usuário e insere no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request){
	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UsersRepository(db)
	user.Id, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, user)
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

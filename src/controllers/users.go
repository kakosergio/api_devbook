package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// CreateUser cria um usuário e insere no banco de dados
func Create(w http.ResponseWriter, r *http.Request){
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
	
	err = user.Prepare()
	if err != nil {
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

// FindAll busca no banco de dados todos os usuários cadastrados
func FindAll(w http.ResponseWriter, r *http.Request){
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UsersRepository(db)
	users, err := repository.Find(nameOrNick)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

// FindUserById busca um usuário pelo seu id
func FindById(w http.ResponseWriter, r *http.Request){

}

// UpdateUser atualiza as informações de um usuário pelo seu id
func Update(w http.ResponseWriter, r *http.Request){

}

// DeleteUser apaga um usuário do banco de dados
func Delete(w http.ResponseWriter, r *http.Request){

}

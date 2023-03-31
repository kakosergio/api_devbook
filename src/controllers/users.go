package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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
	
	err = user.Prepare("signin")
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
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
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
	user, err := repository.FindById(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)

}

// UpdateUser atualiza as informações de um usuário pelo seu id
func Update(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10 , 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil{
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("edit"); err != nil {
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
	if err = repository.Update(userId, user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)

}

// DeleteUser apaga um usuário do banco de dados
func Delete(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)

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
	if err = repository.Delete(userId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent,nil)

}

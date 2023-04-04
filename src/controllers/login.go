package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	repository := repositories.UsersRepository(db)

	userStored, err := repository.FindByEmail(user.Email)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(userStored)

}

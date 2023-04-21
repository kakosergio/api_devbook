package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePub cria publicação nova
func CreatePub(w http.ResponseWriter, r *http.Request){
	userID, err := auth.ExtractIdFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var pub models.Publication
	if err = json.Unmarshal(requestBody, &pub); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	pub.AuthorId = userID

	if err = pub.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.PublicationsRepository(db)
	pub.ID, err = repository.Create(pub)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, pub)

}

// FindPubs encontra publicações dos seguidos pelo usuário, em seu feed
func FindPubs(w http.ResponseWriter, r *http.Request){

}

// FindPub encontra uma publicação em específico
func FindPub(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	pubID, err := strconv.ParseUint(params["id"], 10, 64)
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

	repository := repositories.PublicationsRepository(db)
	publication, err := repository.FindById(pubID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publication)

}

// UpdatePub atualiza uma publicação existente
func UpdatePub(w http.ResponseWriter, r *http.Request){

}

// DeletePub apaga uma publicação
func DeletePub(w http.ResponseWriter, r *http.Request){

}

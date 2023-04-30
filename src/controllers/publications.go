package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
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
	userID, err := auth.ExtractIdFromToken(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.PublicationsRepository(db)
	var publications []models.Publication
	publications, err = repository.FindPubs(userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publications)

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
	userID, err := auth.ExtractIdFromToken(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

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
	publicationSavedinDb, err := repository.FindById(pubID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if publicationSavedinDb.AuthorId != userID {
		responses.Error(w, http.StatusForbidden, errors.New("you can't update another user's publication"))
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

	if err = pub.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.Update(pubID, pub); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

// DeletePub apaga uma publicação
func DeletePub(w http.ResponseWriter, r *http.Request){
	userID, err := auth.ExtractIdFromToken(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

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
	publicationSavedinDb, err := repository.FindById(pubID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if publicationSavedinDb.AuthorId != userID {
		responses.Error(w, http.StatusForbidden, errors.New("you can't delete another user's publication"))
		return
	}

	if err = repository.Delete(pubID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FindPubByUser busca todas as publicações de um determinado usuário
func FindPubByUser (w http.ResponseWriter, r *http.Request){
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

	repository := repositories.PublicationsRepository(db)
	publications, err := repository.FindByUser(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publications)
}

// LikePub curte uma publicação de um determinado usuário
func LikePub (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	pubId, err := strconv.ParseUint(params["id"], 10, 64)

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
	if err = repository.Like(pubId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// UnlikePub descurte uma determinada publicação de um usuário
func UnlikePub (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	pubId, err := strconv.ParseUint(params["id"], 10, 64)

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
	if err = repository.Unlike(pubId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
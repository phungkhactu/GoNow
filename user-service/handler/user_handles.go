package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"GoNow/user-service/dao"
	"GoNow/user-service/models"
	"GoNow/user-service/utils"

	"github.com/gorilla/mux"
)

// UsersDao : Database Access User Object
var UsersDao = dao.UsersDAO{}

// AllUsersEndPoint : GET list of users
func AllUsersEndPoint(w http.ResponseWriter, r *http.Request) {
	users, err := UsersDao.FindAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, users)
}

// FindUserEndpointByID : GET a user by its ID
func FindUserEndpointByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := UsersDao.FindByID(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, user)
}

// FindUserEndpointByType : GET a user by its Type
func FindUserEndpointByType(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := UsersDao.FindByType(params["type"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid User Type")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, user)
}

// CreateUserEndPoint : POST a new user
func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	if err := UsersDao.Insert(user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, user)
}

// UpdateUserEndPoint : PUT update an existing user
func UpdateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := UsersDao.Update(user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteUserEndPoint : DELETE an existing user
func DeleteUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := UsersDao.Delete(user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

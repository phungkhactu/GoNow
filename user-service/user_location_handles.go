package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"GoNow/user-service/models"
	"GoNow/user-service/utils"

	"github.com/gorilla/mux"
)

// AllUserLocationsEndPoint : GET list of user locations
func AllUserLocationsEndPoint(w http.ResponseWriter, r *http.Request) {
	userLocations, err := UserLocationsDao.FindAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, userLocations)
}

// FindUserLocationEndpoint : GET a user location by its ID
func FindUserLocationEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userLocation, err := UserLocationsDao.FindByID(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Location ID")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, userLocation)
}

// CreateUserLocationEndPoint : POST a new user location
func CreateUserLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var userLocation models.UserLocation
	if err := json.NewDecoder(r.Body).Decode(&userLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	userLocation.ID = bson.NewObjectId()
	if err := UserLocationsDao.Insert(userLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, userLocation)
}

// UpdateUserLocationEndPoint : PUT update an existing location
func UpdateUserLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var userLocation models.UserLocation
	if err := json.NewDecoder(r.Body).Decode(&userLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := UserLocationsDao.Update(userLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteUserLocationEndPoint : DELETE an existing user location
func DeleteUserLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var userLocation models.UserLocation
	if err := json.NewDecoder(r.Body).Decode(&userLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := UserLocationsDao.Delete(userLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

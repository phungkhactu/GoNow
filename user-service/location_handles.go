package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	. "GoNow/user-service/models"
	"GoNow/user-service/utils"

	"github.com/gorilla/mux"
)

// GET list of locations
func AllLocationsEndPoint(w http.ResponseWriter, r *http.Request) {
	locations, err := LocationsDao.FindAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, locations)
}

// GET a location by its ID
func FindLocationEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	location, err := LocationsDao.FindById(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Location ID")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, location)
}

// POST a new location
func CreateLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var location Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	location.ID = bson.NewObjectId()
	if err := LocationsDao.Insert(location); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, location)
}

// PUT update an existing location
func UpdateLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var location Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := LocationsDao.Update(location); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing location
func DeleteLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var location Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := LocationsDao.Delete(location); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

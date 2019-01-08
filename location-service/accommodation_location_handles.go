package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"GoNow/location-service/models"
	"GoNow/location-service/utils"

	"github.com/gorilla/mux"
)

// AllAccommodationLocationsEndPoint : GET list of accommodation locations
func AllAccommodationLocationsEndPoint(w http.ResponseWriter, r *http.Request) {
	accommodationLocations, err := AccommodationLocationsDao.FindAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, accommodationLocations)
}

// FindAccommodationLocationEndpoint : GET a accommodation location by its ID
func FindAccommodationLocationEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accommodationLocation, err := AccommodationLocationsDao.FindByID(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Location ID")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, accommodationLocation)
}

// CreateAccommodationLocationEndPoint : POST a new accommodation location
func CreateAccommodationLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var accommodationLocation models.AccommodationLocation
	if err := json.NewDecoder(r.Body).Decode(&accommodationLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	accommodationLocation.ID = bson.NewObjectId()
	if err := AccommodationLocationsDao.Insert(accommodationLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, accommodationLocation)
}

// UpdateAccommodationLocationEndPoint : PUT update an existing accommodation location
func UpdateAccommodationLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var accommodationLocation models.AccommodationLocation
	if err := json.NewDecoder(r.Body).Decode(&accommodationLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := AccommodationLocationsDao.Update(accommodationLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteAccommodationLocationEndPoint : DELETE an existing accommodation location
func DeleteAccommodationLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var accommodationLocation models.AccommodationLocation
	if err := json.NewDecoder(r.Body).Decode(&accommodationLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := AccommodationLocationsDao.Delete(accommodationLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

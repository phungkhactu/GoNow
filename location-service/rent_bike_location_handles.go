package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"GoNow/location-service/models"
	"GoNow/location-service/utils"

	"github.com/gorilla/mux"
)

// AllRentBikeLocationsEndPoint : GET list of rentBike locations
func AllRentBikeLocationsEndPoint(w http.ResponseWriter, r *http.Request) {
	rentBikeLocations, err := RentBikeLocationsDao.FindAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, rentBikeLocations)
}

// FindRentBikeLocationEndpoint : GET a rentBike location by its ID
func FindRentBikeLocationEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rentBikeLocation, err := RentBikeLocationsDao.FindByID(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Location ID")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, rentBikeLocation)
}

// CreateRentBikeLocationEndPoint : POST a new rentBike location
func CreateRentBikeLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var rentBikeLocation models.RentBikeLocation
	if err := json.NewDecoder(r.Body).Decode(&rentBikeLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	rentBikeLocation.ID = bson.NewObjectId()
	if err := RentBikeLocationsDao.Insert(rentBikeLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, rentBikeLocation)
}

// UpdateRentBikeLocationEndPoint : PUT update an existing rentBike location
func UpdateRentBikeLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var rentBikeLocation models.RentBikeLocation
	if err := json.NewDecoder(r.Body).Decode(&rentBikeLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := RentBikeLocationsDao.Update(rentBikeLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteRentBikeLocationEndPoint : DELETE an existing rentBike location
func DeleteRentBikeLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var rentBikeLocation models.RentBikeLocation
	if err := json.NewDecoder(r.Body).Decode(&rentBikeLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := RentBikeLocationsDao.Delete(rentBikeLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

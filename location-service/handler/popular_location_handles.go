package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"GoNow/location-service/dao"
	"GoNow/location-service/models"
	"GoNow/location-service/utils"

	"github.com/gorilla/mux"
)

// PopularLocationsDao : Database Access Popular Location Object
var PopularLocationsDao = dao.PopularLocationsDAO{}

// AllPopularLocationsEndPoint : GET list of popular locations
func AllPopularLocationsEndPoint(w http.ResponseWriter, r *http.Request) {
	popularLocations, err := PopularLocationsDao.FindAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, popularLocations)
}

// FindPopularLocationEndpoint : GET a popular location by its ID
func FindPopularLocationEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	popularLocation, err := PopularLocationsDao.FindByID(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Location ID")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, popularLocation)
}

// CreatePopularLocationEndPoint : POST a new popular location
func CreatePopularLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var popularLocation models.PopularLocation
	if err := json.NewDecoder(r.Body).Decode(&popularLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	popularLocation.ID = bson.NewObjectId()
	if err := PopularLocationsDao.Insert(popularLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, popularLocation)
}

// UpdatePopularLocationEndPoint : PUT update an existing popular location
func UpdatePopularLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var popularLocation models.PopularLocation
	if err := json.NewDecoder(r.Body).Decode(&popularLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := PopularLocationsDao.Update(popularLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeletePopularLocationEndPoint : DELETE an existing popular location
func DeletePopularLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var popularLocation models.PopularLocation
	if err := json.NewDecoder(r.Body).Decode(&popularLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := PopularLocationsDao.Delete(popularLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

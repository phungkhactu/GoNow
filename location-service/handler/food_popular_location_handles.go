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

// FoodPopularLocationsDao : Database Access Food Popular Location Object
var FoodPopularLocationsDao = dao.FoodPopularLocationsDAO{}

// AllFoodPopularLocationsEndPoint : GET list of food popular locations
func AllFoodPopularLocationsEndPoint(w http.ResponseWriter, r *http.Request) {
	foodPopularLocations, err := FoodPopularLocationsDao.FindAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, foodPopularLocations)
}

// FindFoodPopularLocationEndpoint : GET a food popular location by its ID
func FindFoodPopularLocationEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	foodPopularLocation, err := FoodPopularLocationsDao.FindByID(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Location ID")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, foodPopularLocation)
}

// CreateFoodPopularLocationEndPoint : POST a new food popular location
func CreateFoodPopularLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var foodPopularLocation models.FoodPopularLocation
	if err := json.NewDecoder(r.Body).Decode(&foodPopularLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	foodPopularLocation.ID = bson.NewObjectId()
	if err := FoodPopularLocationsDao.Insert(foodPopularLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, foodPopularLocation)
}

// UpdateFoodPopularLocationEndPoint : PUT update an existing popular location
func UpdateFoodPopularLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var foodPopularLocation models.FoodPopularLocation
	if err := json.NewDecoder(r.Body).Decode(&foodPopularLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := FoodPopularLocationsDao.Update(foodPopularLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteFoodPopularLocationEndPoint : DELETE an existing popular location
func DeleteFoodPopularLocationEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var foodpopularLocation models.FoodPopularLocation
	if err := json.NewDecoder(r.Body).Decode(&foodpopularLocation); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := FoodPopularLocationsDao.Delete(foodpopularLocation); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

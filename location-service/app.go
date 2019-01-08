package main

import (
	"GoNow/location-service/config"
	"GoNow/location-service/dao"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Config : variable handle toml config file
var Config = config.Config{}

// PopularLocationsDao : Database Access Popular Location Object
var PopularLocationsDao = dao.PopularLocationsDAO{}

// FoodPopularLocationsDao : Database Access Food Popular Location Object
var FoodPopularLocationsDao = dao.FoodPopularLocationsDAO{}

// RentBikeLocationsDao : Database Access Rent Bike Location Object
var RentBikeLocationsDao = dao.RentBikeLocationsDAO{}

// AccommodationLocationsDao : Database Access Accommodation Location Object
var AccommodationLocationsDao = dao.AccommodationLocationsDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	Config.Read()

	PopularLocationsDao.Server = Config.Server
	PopularLocationsDao.Database = Config.Database
	PopularLocationsDao.Connect()

	FoodPopularLocationsDao.Server = Config.Server
	FoodPopularLocationsDao.Database = Config.Database
	FoodPopularLocationsDao.Connect()

	RentBikeLocationsDao.Server = Config.Server
	RentBikeLocationsDao.Database = Config.Database
	RentBikeLocationsDao.Connect()

	AccommodationLocationsDao.Server = Config.Server
	AccommodationLocationsDao.Database = Config.Database
	AccommodationLocationsDao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/popularlocaltions", AllPopularLocationsEndPoint).Methods("GET")
	r.HandleFunc("/popularlocaltions", CreatePopularLocationEndPoint).Methods("POST")
	r.HandleFunc("/popularlocaltions", UpdatePopularLocationEndPoint).Methods("PUT")
	r.HandleFunc("/popularlocaltions", DeletePopularLocationEndPoint).Methods("DELETE")
	r.HandleFunc("/popularlocaltions/{id}", FindPopularLocationEndpoint).Methods("GET")

	r.HandleFunc("/foodpopularlocaltions", AllPopularLocationsEndPoint).Methods("GET")
	r.HandleFunc("/foodpopularlocaltions", CreatePopularLocationEndPoint).Methods("POST")
	r.HandleFunc("/foodpopularlocaltions", UpdatePopularLocationEndPoint).Methods("PUT")
	r.HandleFunc("/foodpopularlocaltions", DeletePopularLocationEndPoint).Methods("DELETE")
	r.HandleFunc("/foodpopularlocaltions/{id}", FindPopularLocationEndpoint).Methods("GET")

	r.HandleFunc("/rentbikelocaltions", AllRentBikeLocationsEndPoint).Methods("GET")
	r.HandleFunc("/rentbikelocaltions", CreateRentBikeLocationEndPoint).Methods("POST")
	r.HandleFunc("/rentbikelocaltions", UpdateRentBikeLocationEndPoint).Methods("PUT")
	r.HandleFunc("/rentbikelocaltions", DeleteRentBikeLocationEndPoint).Methods("DELETE")
	r.HandleFunc("/rentbikelocaltions/{id}", FindRentBikeLocationEndpoint).Methods("GET")

	r.HandleFunc("/accommodationlocaltions", AllRentBikeLocationsEndPoint).Methods("GET")
	r.HandleFunc("/accommodationlocaltions", CreateRentBikeLocationEndPoint).Methods("POST")
	r.HandleFunc("/accommodationlocaltions", UpdateRentBikeLocationEndPoint).Methods("PUT")
	r.HandleFunc("/accommodationlocaltions", DeleteRentBikeLocationEndPoint).Methods("DELETE")
	r.HandleFunc("/accommodationlocaltions/{id}", FindRentBikeLocationEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

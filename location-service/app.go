package main

import (
	"GoNow/location-service/config"
	"GoNow/location-service/handler"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Config : variable handle toml config file
var Config = config.Config{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	Config.Read()

	handler.PopularLocationsDao.Server = Config.Server
	handler.PopularLocationsDao.Database = Config.Database
	handler.PopularLocationsDao.Connect()

	handler.FoodPopularLocationsDao.Server = Config.Server
	handler.FoodPopularLocationsDao.Database = Config.Database
	handler.FoodPopularLocationsDao.Connect()

	handler.RentBikeLocationsDao.Server = Config.Server
	handler.RentBikeLocationsDao.Database = Config.Database
	handler.RentBikeLocationsDao.Connect()

	handler.AccommodationLocationsDao.Server = Config.Server
	handler.AccommodationLocationsDao.Database = Config.Database
	handler.AccommodationLocationsDao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/popularlocaltions", handler.AllPopularLocationsEndPoint).Methods("GET")
	r.HandleFunc("/popularlocaltions", handler.CreatePopularLocationEndPoint).Methods("POST")
	r.HandleFunc("/popularlocaltions", handler.UpdatePopularLocationEndPoint).Methods("PUT")
	r.HandleFunc("/popularlocaltions", handler.DeletePopularLocationEndPoint).Methods("DELETE")
	r.HandleFunc("/popularlocaltions/{id}", handler.FindPopularLocationEndpoint).Methods("GET")

	r.HandleFunc("/foodpopularlocaltions", handler.AllPopularLocationsEndPoint).Methods("GET")
	r.HandleFunc("/foodpopularlocaltions", handler.CreatePopularLocationEndPoint).Methods("POST")
	r.HandleFunc("/foodpopularlocaltions", handler.UpdatePopularLocationEndPoint).Methods("PUT")
	r.HandleFunc("/foodpopularlocaltions", handler.DeletePopularLocationEndPoint).Methods("DELETE")
	r.HandleFunc("/foodpopularlocaltions/{id}", handler.FindPopularLocationEndpoint).Methods("GET")

	r.HandleFunc("/rentbikelocaltions", handler.AllRentBikeLocationsEndPoint).Methods("GET")
	r.HandleFunc("/rentbikelocaltions", handler.CreateRentBikeLocationEndPoint).Methods("POST")
	r.HandleFunc("/rentbikelocaltions", handler.UpdateRentBikeLocationEndPoint).Methods("PUT")
	r.HandleFunc("/rentbikelocaltions", handler.DeleteRentBikeLocationEndPoint).Methods("DELETE")
	r.HandleFunc("/rentbikelocaltions/{id}", handler.FindRentBikeLocationEndpoint).Methods("GET")

	r.HandleFunc("/accommodationlocaltions", handler.AllRentBikeLocationsEndPoint).Methods("GET")
	r.HandleFunc("/accommodationlocaltions", handler.CreateRentBikeLocationEndPoint).Methods("POST")
	r.HandleFunc("/accommodationlocaltions", handler.UpdateRentBikeLocationEndPoint).Methods("PUT")
	r.HandleFunc("/accommodationlocaltions", handler.DeleteRentBikeLocationEndPoint).Methods("DELETE")
	r.HandleFunc("/accommodationlocaltions/{id}", handler.FindRentBikeLocationEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

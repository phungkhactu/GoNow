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

// PopularLocationsDao : Database Access Location Object
var PopularLocationsDao = dao.PopularLocationsDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	Config.Read()

	PopularLocationsDao.Server = Config.Server
	PopularLocationsDao.Database = Config.Database
	PopularLocationsDao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/popularlocaltions", AllPopularLocationsEndPoint).Methods("GET")
	r.HandleFunc("/popularlocaltions", CreatePopularLocationEndPoint).Methods("POST")
	r.HandleFunc("/popularlocaltions", UpdatePopularLocationEndPoint).Methods("PUT")
	r.HandleFunc("/popularlocaltions", DeletePopularLocationEndPoint).Methods("DELETE")
	r.HandleFunc("/popularlocaltions/{id}", FindPopularLocationEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

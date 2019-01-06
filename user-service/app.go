package main

import (
	. "GoNow/user-service/config"
	. "GoNow/user-service/dao"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Config = Config{}
var UsersDao = UsersDAO{}
var LocationsDao = LocationsDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	Config.Read()

	UsersDao.Server = Config.Server
	UsersDao.Database = Config.Database
	UsersDao.Connect()

	LocationsDao.Server = Config.Server
	LocationsDao.Database = Config.Database
	LocationsDao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", AllUsersEndPoint).Methods("GET")
	r.HandleFunc("/users", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/users", UpdateUserEndPoint).Methods("PUT")
	r.HandleFunc("/users", DeleteUserEndPoint).Methods("DELETE")
	r.HandleFunc("/users/{id}", FindUserEndpointById).Methods("GET")
	r.HandleFunc("/users/{type}", FindUserEndpointByType).Methods("GET")
	r.HandleFunc("/localtions", AllLocationsEndPoint).Methods("GET")
	r.HandleFunc("/localtions", CreateLocationEndPoint).Methods("POST")
	r.HandleFunc("/localtions", UpdateLocationEndPoint).Methods("PUT")
	r.HandleFunc("/localtions", DeleteLocationEndPoint).Methods("DELETE")
	r.HandleFunc("/localtions/{id}", FindLocationEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

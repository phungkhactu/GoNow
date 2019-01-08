package main

import (
	"GoNow/user-service/config"
	"GoNow/user-service/dao"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Config : variable handle toml config file
var Config = config.Config{}

// UsersDao : Database Access User Object
var UsersDao = dao.UsersDAO{}

// UserLocationsDao : Database Access Location Object
var UserLocationsDao = dao.UserLocationsDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	Config.Read()

	UsersDao.Server = Config.Server
	UsersDao.Database = Config.Database
	UsersDao.Connect()

	UserLocationsDao.Server = Config.Server
	UserLocationsDao.Database = Config.Database
	UserLocationsDao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", AllUsersEndPoint).Methods("GET")
	r.HandleFunc("/users", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/users", UpdateUserEndPoint).Methods("PUT")
	r.HandleFunc("/users", DeleteUserEndPoint).Methods("DELETE")
	r.HandleFunc("/users/{id}", FindUserEndpointByID).Methods("GET")
	r.HandleFunc("/users/{type}", FindUserEndpointByType).Methods("GET")
	r.HandleFunc("/userlocaltions", AllUserLocationsEndPoint).Methods("GET")
	r.HandleFunc("/userlocaltions", CreateUserLocationEndPoint).Methods("POST")
	r.HandleFunc("/userlocaltions", UpdateUserLocationEndPoint).Methods("PUT")
	r.HandleFunc("/userlocaltions", DeleteUserLocationEndPoint).Methods("DELETE")
	r.HandleFunc("/userlocaltions/{id}", FindUserLocationEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

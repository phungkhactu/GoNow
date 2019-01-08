package main

import (
	"GoNow/user-service/config"
	"GoNow/user-service/handler"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Config : variable handle toml config file
var Config = config.Config{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	Config.Read()

	handler.UsersDao.Server = Config.Server
	handler.UsersDao.Database = Config.Database
	handler.UsersDao.Connect()

	handler.UserLocationsDao.Server = Config.Server
	handler.UserLocationsDao.Database = Config.Database
	handler.UserLocationsDao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", handler.AllUsersEndPoint).Methods("GET")
	r.HandleFunc("/users", handler.CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/users", handler.UpdateUserEndPoint).Methods("PUT")
	r.HandleFunc("/users", handler.DeleteUserEndPoint).Methods("DELETE")
	r.HandleFunc("/users/{id}", handler.FindUserEndpointByID).Methods("GET")
	r.HandleFunc("/users/{type}", handler.FindUserEndpointByType).Methods("GET")

	r.HandleFunc("/userlocaltions", handler.AllUserLocationsEndPoint).Methods("GET")
	r.HandleFunc("/userlocaltions", handler.CreateUserLocationEndPoint).Methods("POST")
	r.HandleFunc("/userlocaltions", handler.UpdateUserLocationEndPoint).Methods("PUT")
	r.HandleFunc("/userlocaltions", handler.DeleteUserLocationEndPoint).Methods("DELETE")
	r.HandleFunc("/userlocaltions/{id}", handler.FindUserLocationEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

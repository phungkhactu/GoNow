package dao

import (
	"log"

	"GoNow/user-service/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//UserLocationsDAO : Represents database server and credentials
type UserLocationsDAO struct {
	Server   string
	Database string
}

var userLocationDB *mgo.Database

const (
	LOCATIONS_COLLECTION = "user_locations"
)

// Connect : Establish a connection to database
func (m *UserLocationsDAO) Connect() {
	session, err := mgo.Dial(m.Server + ":32769")
	if err != nil {
		log.Fatal(err)
	}
	userLocationDB = session.DB(m.Database)
}

// FindAll : Find list of users location
func (m *UserLocationsDAO) FindAll() ([]models.UserLocation, error) {
	var userLocations []models.UserLocation
	err := userLocationDB.C(LOCATIONS_COLLECTION).Find(bson.M{}).All(&userLocations)
	return userLocations, err
}

// FindByID : Find a user location by its id
func (m *UserLocationsDAO) FindByID(id string) (models.UserLocation, error) {
	var userLocation models.UserLocation
	err := userLocationDB.C(LOCATIONS_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&userLocation)
	return userLocation, err
}

// Insert a location into database
func (m *UserLocationsDAO) Insert(userLocation models.UserLocation) error {
	err := userLocationDB.C(LOCATIONS_COLLECTION).Insert(&userLocation)
	return err
}

// Delete an existing location
func (m *UserLocationsDAO) Delete(userLocation models.UserLocation) error {
	err := userLocationDB.C(LOCATIONS_COLLECTION).Remove(&userLocation)
	return err
}

// Update an existing location
func (m *UserLocationsDAO) Update(userLocation models.UserLocation) error {
	err := userLocationDB.C(LOCATIONS_COLLECTION).UpdateId(userLocation.ID, &userLocation)
	return err
}

package dao

import (
	"log"

	"GoNow/location-service/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AccommodationLocationsDAO : Represents database server and credentials
type AccommodationLocationsDAO struct {
	Server   string
	Database string
}

var accommodationLocationDB *mgo.Database

const (
	ACCOMMODATION_LOCATIONS_COLLECTION = "accommodation_locations"
)

// Connect : Establish a connection to database
func (m *AccommodationLocationsDAO) Connect() {
	session, err := mgo.Dial(m.Server + ":32769")
	if err != nil {
		log.Fatal(err)
	}
	accommodationLocationDB = session.DB(m.Database)
}

// FindAll : Find list of locations
func (m *AccommodationLocationsDAO) FindAll() ([]models.AccommodationLocation, error) {
	var accommodationLocations []models.AccommodationLocation
	err := accommodationLocationDB.C(POPULAR_LOCATIONS_COLLECTION).Find(bson.M{}).All(&accommodationLocations)
	return accommodationLocations, err
}

// FindByID : Find a location by its id
func (m *AccommodationLocationsDAO) FindByID(id string) (models.AccommodationLocation, error) {
	var accommodationLocation models.AccommodationLocation
	err := accommodationLocationDB.C(POPULAR_LOCATIONS_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&accommodationLocation)
	return accommodationLocation, err
}

// Insert a location into database
func (m *AccommodationLocationsDAO) Insert(location models.AccommodationLocation) error {
	err := accommodationLocationDB.C(POPULAR_LOCATIONS_COLLECTION).Insert(&location)
	return err
}

// Delete an existing location
func (m *AccommodationLocationsDAO) Delete(location models.AccommodationLocation) error {
	err := accommodationLocationDB.C(POPULAR_LOCATIONS_COLLECTION).Remove(&location)
	return err
}

// Update an existing location
func (m *AccommodationLocationsDAO) Update(location models.AccommodationLocation) error {
	err := accommodationLocationDB.C(POPULAR_LOCATIONS_COLLECTION).UpdateId(location.ID, &location)
	return err
}

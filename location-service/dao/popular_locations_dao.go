package dao

import (
	"log"

	"GoNow/location-service/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PopularLocationsDAO : Represents database server and credentials
type PopularLocationsDAO struct {
	Server   string
	Database string
}

var popularLocationDB *mgo.Database

const (
	POPULAR_LOCATIONS_COLLECTION = "popular_locations"
)

// Connect : Establish a connection to database
func (m *PopularLocationsDAO) Connect() {
	session, err := mgo.Dial(m.Server + ":32769")
	if err != nil {
		log.Fatal(err)
	}
	popularLocationDB = session.DB(m.Database)
}

// FindAll : Find list of locations
func (m *PopularLocationsDAO) FindAll() ([]models.PopularLocation, error) {
	var popularLocations []models.PopularLocation
	err := popularLocationDB.C(POPULAR_LOCATIONS_COLLECTION).Find(bson.M{}).All(&popularLocations)
	return popularLocations, err
}

// FindByID : Find a location by its id
func (m *PopularLocationsDAO) FindByID(id string) (models.PopularLocation, error) {
	var popularLocation models.PopularLocation
	err := popularLocationDB.C(POPULAR_LOCATIONS_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&popularLocation)
	return popularLocation, err
}

// Insert a location into database
func (m *PopularLocationsDAO) Insert(location models.PopularLocation) error {
	err := popularLocationDB.C(POPULAR_LOCATIONS_COLLECTION).Insert(&location)
	return err
}

// Delete an existing location
func (m *PopularLocationsDAO) Delete(location models.PopularLocation) error {
	err := popularLocationDB.C(POPULAR_LOCATIONS_COLLECTION).Remove(&location)
	return err
}

// Update an existing location
func (m *PopularLocationsDAO) Update(location models.PopularLocation) error {
	err := popularLocationDB.C(POPULAR_LOCATIONS_COLLECTION).UpdateId(location.ID, &location)
	return err
}

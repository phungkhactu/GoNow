package dao

import (
	"log"

	. "GoNow/user-service/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LocationsDAO struct {
	Server   string
	Database string
}

var location_db *mgo.Database

const (
	LOCATIONS_COLLECTION = "locations"
)

// Connect : Establish a connection to database
func (m *LocationsDAO) Connect() {
	session, err := mgo.Dial(m.Server + ":32769")
	if err != nil {
		log.Fatal(err)
	}
	location_db = session.DB(m.Database)
}

// FindAll : Find list of locations
func (m *LocationsDAO) FindAll() ([]Location, error) {
	var locations []Location
	err := location_db.C(LOCATIONS_COLLECTION).Find(bson.M{}).All(&locations)
	return locations, err
}

// FindById : Find a location by its id
func (m *LocationsDAO) FindById(id string) (Location, error) {
	var location Location
	err := location_db.C(LOCATIONS_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&location)
	return location, err
}

// Insert a location into database
func (m *LocationsDAO) Insert(location Location) error {
	err := location_db.C(LOCATIONS_COLLECTION).Insert(&location)
	return err
}

// Delete an existing location
func (m *LocationsDAO) Delete(location Location) error {
	err := location_db.C(LOCATIONS_COLLECTION).Remove(&location)
	return err
}

// Update an existing location
func (m *LocationsDAO) Update(location Location) error {
	err := location_db.C(LOCATIONS_COLLECTION).UpdateId(location.ID, &location)
	return err
}

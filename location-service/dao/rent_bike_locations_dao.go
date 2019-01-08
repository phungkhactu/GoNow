package dao

import (
	"log"

	"GoNow/location-service/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RentBikeLocationsDAO : Represents database server and credentials
type RentBikeLocationsDAO struct {
	Server   string
	Database string
}

var rentBikeLocationDB *mgo.Database

const (
	RENT_BIKE_LOCATIONS_COLLECTION = "rent_bike_locations"
)

// Connect : Establish a connection to database
func (m *RentBikeLocationsDAO) Connect() {
	session, err := mgo.Dial(m.Server + ":32769")
	if err != nil {
		log.Fatal(err)
	}
	rentBikeLocationDB = session.DB(m.Database)
}

// FindAll : Find list of rentBikeLocations
func (m *RentBikeLocationsDAO) FindAll() ([]models.RentBikeLocation, error) {
	var rentBikeLocations []models.RentBikeLocation
	err := rentBikeLocationDB.C(POPULAR_LOCATIONS_COLLECTION).Find(bson.M{}).All(&rentBikeLocations)
	return rentBikeLocations, err
}

// FindByID : Find a rentBikeLocation by its id
func (m *RentBikeLocationsDAO) FindByID(id string) (models.RentBikeLocation, error) {
	var rentBikeLocation models.RentBikeLocation
	err := rentBikeLocationDB.C(POPULAR_LOCATIONS_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&rentBikeLocation)
	return rentBikeLocation, err
}

// Insert a rentBikeLocation into database
func (m *RentBikeLocationsDAO) Insert(rentBikeLocation models.RentBikeLocation) error {
	err := rentBikeLocationDB.C(POPULAR_LOCATIONS_COLLECTION).Insert(&rentBikeLocation)
	return err
}

// Delete an existing rentBikeLocation
func (m *RentBikeLocationsDAO) Delete(rentBikeLocation models.RentBikeLocation) error {
	err := rentBikeLocationDB.C(POPULAR_LOCATIONS_COLLECTION).Remove(&rentBikeLocation)
	return err
}

// Update an existing rentBikeLocation
func (m *RentBikeLocationsDAO) Update(rentBikeLocation models.RentBikeLocation) error {
	err := rentBikeLocationDB.C(POPULAR_LOCATIONS_COLLECTION).UpdateId(rentBikeLocation.ID, &rentBikeLocation)
	return err
}

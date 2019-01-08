package dao

import (
	"log"

	"GoNow/location-service/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// FoodPopularLocationsDAO : Represents database server and credentials
type FoodPopularLocationsDAO struct {
	Server   string
	Database string
}

var foodPopularLocationDB *mgo.Database

const (
	FOOD_POPULAR_LOCATIONS_COLLECTION = "food_popular_locations"
)

// Connect : Establish a connection to database
func (m *FoodPopularLocationsDAO) Connect() {
	session, err := mgo.Dial(m.Server + ":32769")
	if err != nil {
		log.Fatal(err)
	}
	foodPopularLocationDB = session.DB(m.Database)
}

// FindAll : Find list of locations
func (m *FoodPopularLocationsDAO) FindAll() ([]models.FoodPopularLocation, error) {
	var foodPopularLocation []models.FoodPopularLocation
	err := foodPopularLocationDB.C(POPULAR_LOCATIONS_COLLECTION).Find(bson.M{}).All(&foodPopularLocation)
	return foodPopularLocation, err
}

// FindByID : Find a location by its id
func (m *FoodPopularLocationsDAO) FindByID(id string) (models.FoodPopularLocation, error) {
	var foodPopularLocation models.FoodPopularLocation
	err := foodPopularLocationDB.C(POPULAR_LOCATIONS_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&foodPopularLocation)
	return foodPopularLocation, err
}

// Insert a location into database
func (m *FoodPopularLocationsDAO) Insert(foodPopularLocation models.FoodPopularLocation) error {
	err := foodPopularLocationDB.C(POPULAR_LOCATIONS_COLLECTION).Insert(&foodPopularLocation)
	return err
}

// Delete an existing location
func (m *FoodPopularLocationsDAO) Delete(foodPopularLocation models.FoodPopularLocation) error {
	err := foodPopularLocationDB.C(POPULAR_LOCATIONS_COLLECTION).Remove(&foodPopularLocation)
	return err
}

// Update an existing location
func (m *FoodPopularLocationsDAO) Update(foodPopularLocation models.FoodPopularLocation) error {
	err := foodPopularLocationDB.C(POPULAR_LOCATIONS_COLLECTION).UpdateId(foodPopularLocation.ID, &foodPopularLocation)
	return err
}

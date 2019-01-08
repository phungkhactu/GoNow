package models

import "gopkg.in/mgo.v2/bson"

// Represents a Rent Bike location, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type RentBikeLocation struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Longitude string        `bson:"lon" json:"lon"`
	Latitude  string        `bson:"lat" json:"lat"`
	Rate      int16         `bson:"rate" json:"rate"`
	Price     int16         `bson:"price" json:"price"`
	Comment   string        `bson:"comment" json:"comment"`
}

package models

import "gopkg.in/mgo.v2/bson"

// Represents a popular location, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type PopularLocation struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Longitude string        `bson:"lon" json:"lon"`
	Latitude  string        `bson:"lat" json:"lat"`
	Rate      int16         `bson:"rate" json:"rate"`
	Comment   string        `bson:"comment" json:"comment"`
}

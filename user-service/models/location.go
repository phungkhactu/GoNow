package models

import "gopkg.in/mgo.v2/bson"

// Represents a location, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Location struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Longitude string        `bson:"lon" json:"lon"`
	Latitude  string        `bson:"lat" json:"lat"`
}

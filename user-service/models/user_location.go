package models

import "gopkg.in/mgo.v2/bson"

// Represents a location, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type UserLocation struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"username" json:"username"`
	Longitude string        `bson:"lon" json:"lon"`
	Latitude  string        `bson:"lat" json:"lat"`
}

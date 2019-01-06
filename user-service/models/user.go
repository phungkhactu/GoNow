package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	UserName string        `bson:"username" json:"username"`
	Password string        `bson:"password" json:"password"`
	Email    string        `bson:"email" json:"email"`
	Address  string        `bson:"address" json:"address"`
	Avatar   string        `bson:"avatar" json:"avatar"`
	Type     string        `bson:"type" json:"type"`
}

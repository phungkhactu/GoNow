package dao

import (
	"log"

	. "GoNow/user-service/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UsersDAO struct {
	Server   string
	Database string
}

var user_db *mgo.Database

const (
	USERS_COLLECTION = "users"
)

// Connect : Establish a connection to database
func (m *UsersDAO) Connect() {
	session, err := mgo.Dial(m.Server + ":32769")
	if err != nil {
		log.Fatal(err)
	}
	user_db = session.DB(m.Database)
}

// FindAll : Find list of users
func (m *UsersDAO) FindAll() ([]User, error) {
	var users []User
	err := user_db.C(USERS_COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// FindById : Find a user by its id
func (m *UsersDAO) FindById(id string) (User, error) {
	var user User
	err := user_db.C(USERS_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// FindByType : Find a user by its type
func (m *UsersDAO) FindByType(user_type string) ([]User, error) {
	var users []User
	err := user_db.C(USERS_COLLECTION).Find(bson.M{"type": user_type}).All(&users)
	return users, err
}

// Insert a user into database
func (m *UsersDAO) Insert(user User) error {
	err := user_db.C(USERS_COLLECTION).Insert(&user)
	return err
}

// Delete an existing user
func (m *UsersDAO) Delete(user User) error {
	err := user_db.C(USERS_COLLECTION).Remove(&user)
	return err
}

// Update an existing user
func (m *UsersDAO) Update(user User) error {
	err := user_db.C(USERS_COLLECTION).UpdateId(user.ID, &user)
	return err
}

package dao

import (
	"log"

	"GoNow/user-service/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UsersDAO : Represents database server and credentials
type UsersDAO struct {
	Server   string
	Database string
}

var userDB *mgo.Database

const (
	USERS_COLLECTION = "users"
)

// Connect : Establish a connection to database
func (m *UsersDAO) Connect() {
	session, err := mgo.Dial(m.Server + ":32769")
	if err != nil {
		log.Fatal(err)
	}
	userDB = session.DB(m.Database)
}

// FindAll : Find list of users
func (m *UsersDAO) FindAll() ([]models.User, error) {
	var users []models.User
	err := userDB.C(USERS_COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// FindByID : Find a user by its id
func (m *UsersDAO) FindByID(id string) (models.User, error) {
	var user models.User
	err := userDB.C(USERS_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// FindByType : Find a user by its type
func (m *UsersDAO) FindByType(userType string) ([]models.User, error) {
	var users []models.User
	err := userDB.C(USERS_COLLECTION).Find(bson.M{"type": userType}).All(&users)
	return users, err
}

// Insert a user into database
func (m *UsersDAO) Insert(user models.User) error {
	err := userDB.C(USERS_COLLECTION).Insert(&user)
	return err
}

// Delete an existing user
func (m *UsersDAO) Delete(user models.User) error {
	err := userDB.C(USERS_COLLECTION).Remove(&user)
	return err
}

// Update an existing user
func (m *UsersDAO) Update(user models.User) error {
	err := userDB.C(USERS_COLLECTION).UpdateId(user.ID, &user)
	return err
}

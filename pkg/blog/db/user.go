package db

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User structure
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Email     string        `json:"email" bson:"email"`
	Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
}

// Create new user in mongodb user
func (u *User) Create() (string, error) {
	poolSession := MongoDBSession()
	session := poolSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	if u.ID == "" || len(u.ID) < 12 {
		u.ID = bson.NewObjectId()
	}
	var defaultTime time.Time
	if u.Timestamp == defaultTime {
		u.Timestamp = time.Now()
	}
	c := session.DB(database).C(UserCollection)
	return u.ID.String(), c.Insert(u)
}

// Update existing user
func (u *User) Update() error {
	poolSession := MongoDBSession()
	session := poolSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(database).C(UserCollection)

	return c.UpdateId(u.ID, u)
}

// Delete existing user
func (u *User) Delete() error {
	poolSession := MongoDBSession()
	session := poolSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(database).C(UserCollection)
	return c.RemoveId(u.ID)
}

// Read return existing user
func (u *User) Read() error {
	poolSession := MongoDBSession()
	session := poolSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(database).C(UserCollection)
	return c.FindId(u.ID).One(u)
}

// FindUserByID return single user by ID
func FindUserByID(id bson.ObjectId) (User, error) {
	poolSession := MongoDBSession()
	session := poolSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(database).C(UserCollection)
	user := User{}
	err := c.FindId(id).One(&user)
	return user, err
}

// FindUserByEmail return single user matching email id
func FindUserByEmail(email string) (User, error) {
	poolSession := MongoDBSession()
	session := poolSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(database).C(UserCollection)
	user := User{}
	err := c.Find(bson.M{"email": email}).One(&user)
	return user, err
}

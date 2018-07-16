package db

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// Comment structure
	Comment struct {
		ID        bson.ObjectId `json:"id" bson:"_id"`
		Text      string        `json:"text" bson:"text"`
		User      bson.ObjectId `json:"user" bson:"user"`
		Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
	}

	// Article structure
	Article struct {
		ID        bson.ObjectId `json:"id" bson:"_id"`
		Author    bson.ObjectId `json:"author" bson:"author"`
		Title     string        `json:"title" bson:"title"`
		Text      string        `json:"text" bson:"text"`
		Tags      []string      `json:"tags" bson:"tags"`
		Comments  []Comment     `json:"comments" bson:"comments"`
		Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
	}
)

// GetUser get user from database
func (c *Comment) GetUser(db *mgo.Database) (User, error) {
	user := User{}
	err := db.C("author").FindId(c.User).One(&user)
	return user, err
}

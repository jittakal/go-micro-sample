package db

import (
	"log"
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"
)

const (
	// ToDo read from configuration

	// Hosts MongoDB Database Server Hosts
	Hosts = "192.168.99.100:32332"
	// AuthDatabase used for initial connection authentication
	AuthDatabase = "admin"
	// AuthUserName used for database authentication
	AuthUserName = "root"
	// AuthPassword used for database authentication
	AuthPassword = "password"
	// Database for blogs storage
	Database = "myblogs"
	// UserCollection name of the collection
	UserCollection = "user"
	// ArticleCollection name of the article collection
	ArticleCollection = "articles"
)

var mongoSession *mgo.Session
var once sync.Once

//MongoDBSession will return singe instance of session
func MongoDBSession() *mgo.Session {
	once.Do(func() {
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{Hosts},
			Timeout:  60 * time.Second,
			Database: AuthDatabase,
			Username: AuthUserName,
			Password: AuthPassword,
		}

		// Create a session which maintains a pool of socket connections
		// to our MongoDB.
		var err error
		mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
		if err != nil {
			log.Fatalf("CreateSession: %s\n", err)
		}
	})
	return mongoSession
}

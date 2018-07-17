package db

import (
	"log"
	"sync"
	"time"

	"github.com/jittakal/go-micro-sample/pkg/blog/config"
	mgo "gopkg.in/mgo.v2"
)

const (
	// UserCollection name of the collection
	UserCollection = "user"
	// ArticleCollection name of the article collection
	ArticleCollection = "articles"
)

var (
	once         sync.Once
	mongoSession *mgo.Session

	c = config.GetConfig()
)

//MongoDBSession will return singe instance of session
func MongoDBSession() *mgo.Session {
	once.Do(func() {
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{c.MongoDB.Auth.Servers},
			Timeout:  60 * time.Second,
			Database: c.MongoDB.Auth.Database,
			Username: c.MongoDB.Auth.Username,
			Password: c.MongoDB.Auth.Password,
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

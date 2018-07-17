package db

import (
	"time"

	"github.com/jittakal/go-micro-sample/pkg/blog/util"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// Comment structure
	Comment struct {
		Text      string        `json:"text" bson:"text"`
		User      bson.ObjectId `json:"user" bson:"user"`
		Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
	}

	// Article structure
	Article struct {
		ID        bson.ObjectId `json:"id" bson:"_id"`
		Author    bson.ObjectId `json:"author" bson:"author"`
		Title     string        `json:"title" bson:"title"`
		Content   string        `json:"content" bson:"content"`
		Tags      []string      `json:"tags" bson:"tags"`
		Comments  []Comment     `json:"comments" bson:"comments"`
		Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
	}
)

// GetUser return user details of comment
func (c *Comment) GetUser() (User, error) {
	return FindUserByID(c.User)
}

// Create new article
func (a *Article) Create() (string, error) {
	poolSession := MongoDBSession()
	session := poolSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	if a.ID == "" || len(a.ID) < 12 {
		a.ID = bson.NewObjectId()
	}
	var defaultTime time.Time
	if a.Timestamp == defaultTime {
		a.Timestamp = time.Now()
	}
	c := session.DB(Database).C(ArticleCollection)
	return a.ID.String(), c.Insert(a)
}

// Update exsting article
func (a *Article) Update() error {
	poolSession := MongoDBSession()
	session := poolSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(Database).C(ArticleCollection)

	return c.UpdateId(a.ID, a)
}

// Delete existing article
func (a *Article) Delete() error {
	poolSession := MongoDBSession()
	session := poolSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(Database).C(ArticleCollection)
	return c.RemoveId(a.ID)
}

// Read return existing article
func (a *Article) Read() error {
	poolSession := MongoDBSession()
	session := poolSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(Database).C(ArticleCollection)
	return c.FindId(a.ID).One(a)
}

// AddComment to existing article
func (a *Article) AddComment(c Comment) error {
	var defaultTime time.Time
	if c.Timestamp == defaultTime {
		c.Timestamp = time.Now()
	}
	a.Comments = append(a.Comments[:], c)
	return a.Update()
}

// AddTags to existing article
func (a *Article) AddTags(tags []string) error {
	for _, tag := range tags {
		a.Tags = append(a.Tags[:], tag)
	}
	a.Tags = util.RemoveDuplicates(a.Tags)
	return a.Update()
}

// GetAuthor return user details of article
func (a *Article) GetAuthor() (User, error) {
	return FindUserByID(a.Author)
}

// FindArticleByID return single article by ID
func FindArticleByID(id bson.ObjectId) (Article, error) {
	poolSession := MongoDBSession()
	session := poolSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(Database).C(ArticleCollection)
	article := Article{}
	err := c.FindId(id).One(&article)
	return article, err
}

// FindArticlesByAuthor return list of all articles
func FindArticlesByAuthor(author bson.ObjectId) ([]Article, error) {
	poolSession := MongoDBSession()
	session := poolSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(Database).C(ArticleCollection)
	articles := []Article{}
	err := c.Find(bson.M{"author": author}).All(&articles)
	return articles, err
}

// FindArticlesByTags returns list of articles matching tags
func FindArticlesByTags(tags []string) ([]Article, error) {
	poolSession := MongoDBSession()
	session := poolSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(Database).C(ArticleCollection)
	articles := []Article{}
	err := c.Find(bson.M{"tags": tags}).All(&articles)
	return articles, err
}

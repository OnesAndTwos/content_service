package repositories

import (
	"content_service/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BlogRepository is a constructed repository
type blogRepository struct {
	C *mgo.Collection
}

// Find finds a Blog by reference
func (b *blogRepository) Find(reference string) models.Blog {
	blog := models.Blog{}

	b.C.Find(bson.M{"reference": reference}).One(&blog)

	return blog
}

// NewBlogRepository is the factory function that creates a blogRepository
func NewBlogRepository() *blogRepository {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic("Could not connect to Mongo DB at 'localhost'")
	}

	session.SetMode(mgo.Monotonic, true)

	return &blogRepository{
		session.DB("content_service").C("Blogs"),
	}
}

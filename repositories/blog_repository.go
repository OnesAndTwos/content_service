package repositories

import (
	"content_service/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BlogRepository is a constructed repository
type BlogRepository struct {
	s *mgo.Session
	c *mgo.Collection
}

// Find finds a Blog by reference
func (b *BlogRepository) Find(reference string) models.Blog {
	blog := models.Blog{}

	b.c.Find(bson.M{"reference": reference}).One(&blog)

	return blog
}

// Create allows the creation of a blog
func (b *BlogRepository) Create(m *models.Blog) error {
	return b.c.Insert(&m)
}

// Close releases all resources on the repository
func (b *BlogRepository) Close() {
	b.s.Close()
}

// NewBlogRepository is the factory function that creates a blogRepository
func NewBlogRepository() *BlogRepository {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic("Could not connect to Mongo DB at 'localhost'")
	}

	session.SetMode(mgo.Monotonic, true)

	return &BlogRepository{
		session, session.DB("content_service").C("Blogs"),
	}
}

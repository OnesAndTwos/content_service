package blogs

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BlogRepository is a repository for Blogs
type BlogRepository interface {
	Find(reference string) Blog
	Create(m *Blog) error
	Close()
}

type repository struct {
	s *mgo.Session
	c *mgo.Collection
}

// Find finds a Blog by reference
func (b *repository) Find(reference string) Blog {
	blog := Blog{}

	b.c.Find(bson.M{"reference": reference}).One(&blog)

	return blog
}

// Create allows the creation of a blog
func (b *repository) Create(m *Blog) error {
	return b.c.Insert(&m)
}

// Close releases all resources on the repository
func (b *repository) Close() {
	b.s.Close()
}

// NewRepository is the factory function that creates a blogRepository
func Repository() BlogRepository {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic("Could not connect to Mongo DB at 'localhost'")
	}

	session.SetMode(mgo.Monotonic, true)

	return &repository{
		session, session.DB("content_service").C("Blogs"),
	}
}

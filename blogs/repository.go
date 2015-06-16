package blogs

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repository is a constructed repository
type Repository struct {
	s *mgo.Session
	c *mgo.Collection
}

// Find finds a Blog by reference
func (b *Repository) Find(reference string) Blog {
	blog := Blog{}

	b.c.Find(bson.M{"reference": reference}).One(&blog)

	return blog
}

// Create allows the creation of a blog
func (b *Repository) Create(m *Blog) error {
	return b.c.Insert(&m)
}

// Close releases all resources on the repository
func (b *Repository) Close() {
	b.s.Close()
}

// NewRepository is the factory function that creates a blogRepository
func NewRepository() *Repository {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic("Could not connect to Mongo DB at 'localhost'")
	}

	session.SetMode(mgo.Monotonic, true)

	return &Repository{
		session, session.DB("content_service").C("Blogs"),
	}
}

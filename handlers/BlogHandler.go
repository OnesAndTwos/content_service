package handlers

import (
	"content_service/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//BlogHandler handles GET requests for Blogs
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("localhost")

	vars := mux.Vars(r)
	reference := vars["reference"]

	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("content_service").C("Blogs")

	blog := models.Blog{}

	err = c.Find(bson.M{"reference": reference}).One(&blog)

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(blog)

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

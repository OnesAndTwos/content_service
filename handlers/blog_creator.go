package handlers

import (
	"content_service/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

//BlogCreator handles POST requests to Blogs
func BlogCreator(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("content_service").C("Blogs")

	blog := models.Blog{}

	bodyText, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal([]byte(bodyText), &blog)

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.Insert(&blog)

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

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

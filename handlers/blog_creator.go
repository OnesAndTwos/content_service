package handlers

import (
	"content_service/models"
	"content_service/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//BlogCreator handles POST requests to Blogs
func BlogCreator(w http.ResponseWriter, r *http.Request) {
	blogRepository := repositories.NewBlogRepository()
	defer blogRepository.Close()

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

	err = blogRepository.Create(&blog)

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(blog)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

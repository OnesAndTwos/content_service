package handlers

import (
	"content_service/repositories"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//BlogHandler handles GET requests for Blogs
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	blogRepository := repositories.NewBlogRepository()

	reference := mux.Vars(r)["reference"]

	blog := blogRepository.Find(reference)

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

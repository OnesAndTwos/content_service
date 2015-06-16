package blogs

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Handler handles GET requests for Blogs
func Handler(w http.ResponseWriter, r *http.Request) {
	blogRepository := NewRepository()
	defer blogRepository.Close()

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

package blogs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Creator handles POST requests to Blogs
func Creator(w http.ResponseWriter, r *http.Request) {
	blogRepository := NewRepository()
	defer blogRepository.Close()

	blog := Blog{}

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

package blogs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/onesandtwos/content_service/helpers"
)

//Creator handles POST requests to Blogs
func Creator(repository BlogRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		err = repository.Create(&blog)

		helpers.WriteJSONResponse(w, blog, http.StatusOK)
	})
}

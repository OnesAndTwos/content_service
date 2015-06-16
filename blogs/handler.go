package blogs

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onesandtwos/content_service/helpers"
)

//Handler handles GET requests for Blogs
func Handler(repository BlogRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reference := mux.Vars(r)["reference"]

		blog := repository.Find(reference)

		helpers.WriteJSONResponse(w, blog, http.StatusCreated)
	})
}

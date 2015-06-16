package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteJSONResponse writes the given struct to the writer
func WriteJSONResponse(w http.ResponseWriter, s interface{}, status int) (e error) {
	js, e := json.Marshal(s)

	if e != nil {
		log.Fatal(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	return
}

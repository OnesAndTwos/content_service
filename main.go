package main

import (
	"content_service/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/blogs/{reference}", handlers.BlogHandler).Methods("GET").Name("GetBlog")
	r.HandleFunc("/blogs", handlers.BlogCreator).Methods("POST").Name("CreateBlog")

	http.ListenAndServe(":1234", r)
}

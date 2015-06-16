package main

import (
	"content_service/blogs"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/blogs/{reference}", blogs.Handler).Methods("GET").Name("GetBlog")
	r.HandleFunc("/blogs", blogs.Creator).Methods("POST").Name("CreateBlog")

	http.ListenAndServe(":1234", r)
}

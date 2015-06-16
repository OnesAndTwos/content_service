package main

import (
	"net/http"

	"github.com/onesandtwos/content_service/blogs"

	"github.com/gorilla/mux"
)

var (
	blogRepository = blogs.Repository()
)

func main() {
	defer blogRepository.Close()

	r := mux.NewRouter()

	r.HandleFunc("/blogs/{reference}", blogs.Handler(blogRepository)).
		Methods("GET").
		Name("GetBlog")

	r.HandleFunc("/blogs", blogs.Creator(blogRepository)).
		Methods("POST").
		Name("CreateBlog")

	http.ListenAndServe(":1234", r)
}

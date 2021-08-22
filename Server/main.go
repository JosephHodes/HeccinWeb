package main

import (
	"net/http"
	"github.com/JosephHodes/HeccinWeb/Server/Blog""
)

func main() {
	http.HandleFunc(
		"/blog",
		blog.Blog,
	)
	http.ListenAndServe(":8000", nil)
}

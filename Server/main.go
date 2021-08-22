package main

import (
	"github.com/JosephHodes/HeccinWeb/Server/Blog"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/blog",
		Blog,
	)
	http.ListenAndServe(":8000", nil)
}

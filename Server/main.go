package main

import (
	"github.com/JosephHodes/HeccinWeb/Server/Blog"
	"log"
	"net/http"
)

func main() {
	log.Println("working")
	http.HandleFunc(
		"/",
		blog.Blog,
	)
	http.ListenAndServe(":8000", nil)
}

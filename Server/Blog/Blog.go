package blog

import (
	"fmt"
	"log"
	"net/http"
)

func Blog(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			fmt.Fprintf(w, "hello worlds")
			log.Println("works")
		}
	}
}

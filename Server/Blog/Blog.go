package blog

import (
	"fmt"
	"log"
	"net/http"
)

func Blog(w http.ResponseWriter, r *http.Request) {
	switch r.method {
	case "GET":
		{
			fmt.Fprintf(w, "hello worlds")
			log.println("response sent")
		}
	}
}

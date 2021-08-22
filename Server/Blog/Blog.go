package blog

import (
	"net/http"
)

func Blog(w http.ResponseWriter, r *http.Request) {
 case r.method
   "GET" {
	fmt.Fprintf(w, "hello worlds")
	log.println("response sent")
	}
}

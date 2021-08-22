package blog

import (
	"net/http"
)

func Blog(w http.ResponseWriter, r *http.Request) {
	w.Write("helloworld")
}

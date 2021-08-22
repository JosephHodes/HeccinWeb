package Server

import "net/http"


func main(){
	http.HandleFunc(
		"/blog",
		blog
	)
	http.ListenAndServe(":8000", nil)
}
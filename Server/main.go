package main

import (
	"fmt"
	"github.com/JosephHodes/HeccinWeb/Server/Blog"
	"github.com/go-redis/redis"
	"log"
	"net/http"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	log.Println("working")
	http.HandleFunc(
		"/",
		blog.Blog,
	)

	http.ListenAndServe(":8000", nil)
}

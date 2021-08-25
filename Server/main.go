package main

import (
	"fmt"
	"github.com/JosephHodes/HeccinWeb/Server/Blog"
	"github.com/JosephHodes/HeccinWeb/Server/Utils"
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

	err := redisClient.Set(redisClient.Context(), "name", "122.223.444", 1)

	val := redisClient.Get(redisClient.Context(), "name")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	http.ListenAndServe(":8000", nil)
}

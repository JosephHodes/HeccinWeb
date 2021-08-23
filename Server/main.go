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
	err := redisClient.Set(redisClient.Context(), "name", "Elliot", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := redisClient.Get(redisClient.Context(), "name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	http.ListenAndServe(":8000", nil)
}

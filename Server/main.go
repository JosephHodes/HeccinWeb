package main

import (
	"fmt"
	"github.com/JosephHodes/HeccinWeb/Server/Blog"
	"github.com/JosephHodes/HeccinWeb/Server/Utils"
	"github.com/go-redis/redis"
	"github.com/go-redsync/redsync"
	"github.com/go-redsync/redsync/redis/redigo"
	redis3 "github.com/gomodule/redigo/redis"
	"log"
	"net/http"
	"reflect"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	p := &redis3.Pool{}
	pool := redigo.NewPool(p)
	log.Println(reflect.TypeOf(pool))
	utils.Mutex = redsync.New(pool)

	utils.RateLimiter("111.111.111.11", redisClient, 10, 5)
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

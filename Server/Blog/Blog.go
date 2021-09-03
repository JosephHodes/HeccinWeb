package blog

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"net/http"
)

var redisClient *redis.Client = nil

func Blog(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			fmt.Fprintf(w, "hello worlds")
			log.Println("works")
		}
	}
}

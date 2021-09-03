package blog

import (
	"fmt"
	utils "github.com/JosephHodes/HeccinWeb/Server/Utils"
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
			utils.IsRateLimited(r.RemoteAddr, redisClient, 10, 1)
		}
	}
}

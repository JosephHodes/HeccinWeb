package blog

import (
	"fmt"
	utils "github.com/JosephHodes/HeccinWeb/Server/Utils"
	"github.com/go-redis/redis"
	"log"
	"net/http"
)

var RedisClient *redis.Client = nil

func Blog(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{

			log.Println("works")
			err := utils.IsRateLimited(r.RemoteAddr, RedisClient, 10, 10)
			if err != nil {
				return
			}
			_, err = fmt.Fprintf(w, "hello worlds")
			if err != nil {
				return
			}
		}
	}
}

package utils

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

// RateLimiter this function you enter the ip and the amount you want to rate limit
// then the expiration for the time you want your ip to expire and it checks the and blocks ips that surpass
// the rate limit
func RateLimiter(ip string, redisClient *redis.Client, rateLimit int, expiration int) bool {
	val, error := redisClient.Get(redisClient.Context(), ip).Result()
	if error == redis.Nil {
		redisClient.Set(redisClient.Context(), ip, 0, time.Duration(expiration))
	} else {
		i, _ := strconv.ParseUint(val, 0, 32)
		if rateLimit == int(i) {
			return false
		} else {
			redisClient.Set(redisClient.Context(), ip, int(i), time.Duration(expiration))
		}
	}
	return true
}

package utils

import (
	"github.com/go-redis/redis"
)

func RateLimiter(ip string, redisClient redis.Client) boolean {
	return true
}

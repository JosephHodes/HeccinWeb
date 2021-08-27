package utils

import (
	main "./../main"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

// IsRateLimited this function you enter the ip and the amount you want to rate limit
// then the expiration for the time you want your ip to expire and it checks the and blocks ips that surpass
// the rate limit

func IsRateLimited(ip string, redisClient *redis.Client, rateLimit int, expiration int) error {
	ahh := main.Redsyncz
	fmt.Println(ahh)
	val, err := redisClient.Get(redisClient.Context(), ip).Result()
	if err != nil {
		if err == redis.Nil {
			err := redisClient.Set(redisClient.Context(), ip, 0, time.Duration(expiration))
			if err != nil {
				return fmt.Errorf("setting content in redis : ratelimit %w", err)
			}
		}
		return fmt.Errorf("getting content from redis : ratelimit: %w", err)
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		fmt.Errorf("strconv  : ratelimit %w", err)
	}
	if rateLimit == i {
		return fmt.Errorf("to many requests : ratelimit")
	}

	i++

	err = redisClient.Set(redisClient.Context(), ip, int(i), time.Duration(expiration))
	if err != nil {
		return fmt.Errorf("setting content in redis : ratelimit: %w", err)

	}
	return nil
}

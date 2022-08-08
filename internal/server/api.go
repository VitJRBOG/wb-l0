package server

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

func order(redisClient *redis.Client, orderUID string) ([]byte, error) {
	result, err := redisClient.Get(orderUID).Result()
	if err != nil {
		return nil, err
	}

	orderJSON, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return orderJSON, nil
}

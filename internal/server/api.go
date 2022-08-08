package server

import (
	"encoding/json"
	"wb-l0/internal/models"

	"github.com/go-redis/redis"
)

func order(redisClient *redis.Client, orderUID string) ([]byte, error) {
	result, err := redisClient.Get(orderUID).Result()
	if err != nil {
		return nil, err
	}

	order := models.Order{}
	err = json.Unmarshal([]byte(result), &order)
	if err != nil {
		return nil, err
	}

	orderJSON, err := json.Marshal(&order)
	if err != nil {
		return nil, err
	}

	return orderJSON, nil
}

package redis

import (
	"encoding/json"
	"fmt"
	"os"
	"wb-l0/config"
	"wb-l0/internal/models"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func NewRedisClient(redisCfg *config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisCfg.Host, redisCfg.HostPort),
		Password: "",
		DB:       0,
	})

	return client
}

func RestoreFromDB(client *redis.Client, dbConnection *gorm.DB) {
	orders, err := models.SelectAllOrders(dbConnection)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to restore data from DB to Redis: %v\n", err)
		os.Exit(1)
	}

	for _, order := range orders {
		jsonOrder, err := json.Marshal(order)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable marshal data from struct Order: %v\n", err)
			continue
		}

		client.Set(order.OrderUID, jsonOrder, 0)
	}
}

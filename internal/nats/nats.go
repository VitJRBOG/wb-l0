package nats

import (
	"encoding/json"
	"fmt"
	"os"

	"wb-l0/config"
	"wb-l0/internal/models"

	"github.com/go-redis/redis"
	"github.com/nats-io/stan.go"
	"gorm.io/gorm"
)

func Connect(cfg *config.NatsConfig) stan.Conn {
	sc, err := stan.Connect(
		cfg.NatsClusterID,
		cfg.NatsClientID,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to Nats Streaming: %v\n", err)
		os.Exit(1)
	}

	return sc
}

func Subcribe(sc stan.Conn, dbConnection *gorm.DB, redisClient *redis.Client) stan.Subscription {
	msgChannel := make(chan []byte)
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		msgChannel <- m.Data
	}, stan.StartWithLastReceived())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong on Nats Streaming: %v\n", err)
		os.Exit(1)
	}

	go parseReceivedData(dbConnection, redisClient, msgChannel)

	return sub
}

func parseReceivedData(dbConnection *gorm.DB, redisClient *redis.Client, msgChannel chan []byte) {
	for {
		rawData := <-msgChannel

		order := models.Order{}
		err := json.Unmarshal(rawData, &order)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable unmarshal data from Nats: %v\n", err)
			continue
		}

		jsonOrder, err := json.Marshal(order)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable marshal data from struct Order: %v\n", err)
			continue
		}

		redisClient.Set(order.OrderUID, jsonOrder, 0)

		order.Create(dbConnection)
	}
}

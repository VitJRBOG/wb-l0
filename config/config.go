package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	PostgresUser     string
	PostgresPassword string
	DBHost           string
	DBHostPort       string
	DBName           string
}

func GetDBConfig() *DBConfig {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetDBConfig godotenv load failed: %v\n", err.Error())
		os.Exit(1)
	}

	cfg := &DBConfig{
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_HOST_PORT"),
		os.Getenv("DB_NAME"),
	}

	return cfg
}

type NatsConfig struct {
	NatsClusterID  string
	NatsClientID   string
	NatsChannel    string
	NatsDurableID  string
	NatsQueueGroup string
}

func GetNatsConfig() *NatsConfig {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetNatsConfig godotenv load failed: %v\n", err.Error())
		os.Exit(1)
	}

	cfg := &NatsConfig{
		os.Getenv("NATS_CLUSTER_ID"),
		os.Getenv("NATS_CLIENT_ID"),
		os.Getenv("NATS_CHANNEL"),
		os.Getenv("NATS_DURABLE_ID"),
		os.Getenv("NATS_QUEUE_GROUP"),
	}

	return cfg
}

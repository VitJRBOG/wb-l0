package cmd

import (
	"fmt"

	"wb-l0/config"
	"wb-l0/internal/db"
	"wb-l0/internal/nats"
	"wb-l0/internal/redis"
	"wb-l0/internal/server"
)

func Execute() {
	dbCfg := config.GetDBConfig()
	databaseURL := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s",
		dbCfg.DBHost, dbCfg.PostgresUser, dbCfg.PostgresPassword, dbCfg.DBHostPort, dbCfg.DBName)
	dbConnection := db.NewConnection(databaseURL)

	redisCfg := config.GetRedisConfig()
	redisClient := redis.NewRedisClient(redisCfg)
	redis.RestoreFromDB(redisClient, dbConnection)

	natsCfg := config.GetNatsConfig()
	sc := nats.Connect(natsCfg)
	sub := nats.Subcribe(sc, dbConnection, redisClient)

	serverCfg := config.GetServerConfig()
	server.Run(serverCfg, redisClient)

	defer sub.Unsubscribe()
	defer sc.Close()
}

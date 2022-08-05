package cmd

import (
	"fmt"

	"wb-l0/config"
	"wb-l0/internal/db"
	"wb-l0/internal/nats"
	"wb-l0/internal/server"
)

func Execute() {
	dbCfg := config.GetDBConfig()
	databaseURL := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s",
		dbCfg.DBHost, dbCfg.PostgresUser, dbCfg.PostgresPassword, dbCfg.DBHostPort, dbCfg.DBName)
	dbConnection := db.NewConnection(databaseURL)

	natsCfg := config.GetNatsConfig()
	sc := nats.Connect(natsCfg)
	sub := nats.Subcribe(sc, dbConnection)

	serverCfg := config.GetServerConfig()
	server.Run(serverCfg, dbConnection)

	defer sub.Unsubscribe()
	defer sc.Close()
}

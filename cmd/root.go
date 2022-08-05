package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"wb-l0/config"
	"wb-l0/internal/db"
	"wb-l0/internal/models"
	"wb-l0/internal/nats"

	"gorm.io/gorm"
)

func Execute() {
	dbCfg := config.GetDBConfig()

	databaseURL := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s",
		dbCfg.DBHost, dbCfg.PostgresUser, dbCfg.PostgresPassword, dbCfg.DBHostPort, dbCfg.DBName)

	dbConnection := db.NewConnection(databaseURL)

	natsCfg := config.GetNatsConfig()

	sc := nats.Connect(natsCfg)

	msgChannel := make(chan []byte)
	sub := nats.Subcribe(sc, msgChannel)

	parseReceivedData(dbConnection, msgChannel)
	parseSelectedData(dbConnection)

	defer sub.Unsubscribe()
	defer sc.Close()
}

func parseReceivedData(dbConnection *gorm.DB, msgChannel chan []byte) {
	for {
		order := models.Order{}
		err := json.Unmarshal(<-msgChannel, &order)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable unmarshal data from Nats: %v\n", err)
		} else {
			order.Create(dbConnection)
		}
	}
}

func parseSelectedData(dbConnection *gorm.DB) {
	order := models.Order{}
	order.Select(dbConnection, "b563feb7b2b84b6test") // FIXME
	fmt.Println(order)
}

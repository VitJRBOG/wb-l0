package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(dsn string) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to Postgres: %v\n", err)
		os.Exit(1)
	}

	return conn
}

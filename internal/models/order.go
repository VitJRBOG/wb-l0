package models

import (
	"fmt"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Order struct {
	OrderUID          string   `gorm:"primaryKey" json:"order_uid"`
	TrackNumber       string   `json:"track_number"`
	Entry             string   `json:"entry"`
	Delivery          Delivery `gorm:"foreignKey:phone" json:"delivery"`
	Payment           Payment  `gorm:"foreignKey:transaction" json:"payment"`
	Items             []Item   `gorm:"foreignKey:rid" json:"items"`
	Locale            string   `json:"locale"`
	InternalSignature string   `json:"internal_signature"`
	CustomerID        string   `json:"customer_id"`
	DeliveryService   string   `json:"delivery_service"`
	Shardkey          string   `json:"shardkey"`
	SmID              int      `json:"sm_id"`
	DateCreated       string   `json:"date_created"`
	OofShard          string   `json:"oof_shard"`
}

func (order *Order) Create(dbConnection *gorm.DB) {
	err := dbConnection.Create(order).Error
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable push data to db: %v\n", err)
	}
}

func (order *Order) Select(dbConnection *gorm.DB, id string) {
	order.OrderUID = id
	err := dbConnection.Preload(clause.Associations).First(&order).Error
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable pull data from db: %v\n", err)
	}
}

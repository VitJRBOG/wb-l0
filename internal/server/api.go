package server

import (
	"encoding/json"
	"wb-l0/internal/models"

	"gorm.io/gorm"
)

func order(dbConnection *gorm.DB, orderUID string) ([]byte, error) {
	order := models.Order{}
	err := order.Select(dbConnection, orderUID)
	if err != nil {
		return nil, err
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	return orderJSON, nil
}

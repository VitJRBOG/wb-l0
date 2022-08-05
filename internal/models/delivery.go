package models

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `gorm:"primaryKey" json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

package models

import "time"

type (
	// Transaction
	Transaction struct {
		Id         int   `gorm:"primary_key" json:"id"`
		UserId     int    `json:"userId"`
		ProductsId  int    `json:"productsId"`
		Quantity   int       `json:"quantity"`
		TotalPrice string    `json:"totalPrice"`
		CreatedAt  time.Time `json:"createdAt"`
		Payment string `json:"payment"`
		Products []Products `gorm:"many2many:transaction_products;" json:"-"`
	}
)
package models

import "time"

type (
	// Products
	Products struct {
		Id          int    	`gorm:"primary_key" json:"id"`
		UserId      int    	`json:"userId"`
		Category    string    	`json:"category"`
		Gender      string    	`json:"gender"`
		ProductName string    	`json:"productName"`
		Price       int64     	`json:"price"`
		DescProduct string    	`json:"descProduct"`
		Stock       int       	`json:"stock"`
		Images      string  	`json:"images"`
		CreatedAt   time.Time 	`json:"createdAt"`
		UpdatedAt   time.Time 	`json:"updatedAt"`
		IsActive    int       	`json:"isActive"`
		User 		User 		`json:"-"`
		Comments 	[]Comments 	`json:"-"`
		Promotion 	[]Promotion `josn:"-"`
		Transaction []Transaction `json:"-" gorm:"many2many:products_transaction;"`
	}
)
package models

import "time"

type (
	// Comments
	Comments struct {
		Id            int    `gorm:"primary_key" json:"id"`
		ProductsId     int    `json:"productsId"`
		UserIdComment int    `json:"userIdComment"`
		Comment       string    `json:"comment"`
		Rating        int       `json:"rating"`
		CreatedAt     time.Time `json:"createdAt"`
		UpdatedAt     time.Time `json:"updatedAt"`
		Products 	Products 	`json:"-"`
	}
)
package models

import "time"

type (
	// Promotion
	Promotion struct {
		Id        int    `gorm:"primary_key" json:"id"`
		ProductsId int    `json:"productsId"`
		PromoName string    `json:"promoName"`
		PromoDesc string    `json:"promoDesc"`
		Promo     int       `json:"promo"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		IsActive  int       `json:"isActive"`
		Products Products `json:"-"`
	}
)
package models

import "time"

type (
	// User
	User struct {
		Id         int `gorm:"primary_key" json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		Gender       string `json:"gender"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		IsActive int `json:"isActive"`
		Products []Products `josn:"-"`
	}
)
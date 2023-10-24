package config

import (
	"FP-BDS-Sanbercode-Go-50-anggi/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB{
	username := "root"
	password := "admin"
	host := "tcp(127.0.0.1:3306)"
	database := "clone_tokopedia"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err.Error())
	}
	if err := db.AutoMigrate(&models.Products{}, &models.User{}, &models.Comments{}, &models.Transaction{}, &models.Promotion{}); err != nil{
		fmt.Println("Database Cannot Connect")
	}
	return db
}
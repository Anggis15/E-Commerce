package config

import (
	"FP-BDS-Sanbercode-Go-50-anggi/models"
	"FP-BDS-Sanbercode-Go-50-anggi/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB{
	username := utils.GetEnv("DATABASE_USERNAME", "root")
  	password := utils.GetEnv("DATABASE_PASSWORD", "password")
  	host := utils.GetEnv("DATABASE_HOST", "127.0.0.1")
  	port := utils.GetEnv("DATABASE_PORT", "3306")
  	database := utils.GetEnv("DATABASE_NAME", "clone_tokopedia")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err.Error())
	}
	if err := db.AutoMigrate(&models.Products{}, &models.User{}, &models.Comments{}, &models.Transaction{}, &models.Promotion{}); err != nil{
		fmt.Println("Database Cannot Connect")
	}
	return db
}
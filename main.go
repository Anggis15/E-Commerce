package main

import (
	"FP-BDS-Sanbercode-Go-50-anggi/config"
	"FP-BDS-Sanbercode-Go-50-anggi/docs"
	"FP-BDS-Sanbercode-Go-50-anggi/routes"
	"FP-BDS-Sanbercode-Go-50-anggi/utils"
	"log"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/


func main() {
	// for load godotenv
    // for env
    environment := utils.GetEnv("ENVIRONMENT", "development")

    if environment == "development" {
      err := godotenv.Load()
      if err != nil {
        log.Fatal("Error loading .env file")
      }
    }
	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Ecomerce."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "localhost:8080"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	
	db := config.ConnectToDatabase()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	r := routes.SetUpRouter(db)
 	r.Run()
}
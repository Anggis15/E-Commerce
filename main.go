package main

import (
	"FP-BDS-Sanbercode-Go-50-anggi/config"
	"FP-BDS-Sanbercode-Go-50-anggi/docs"
	"FP-BDS-Sanbercode-Go-50-anggi/routes"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/


func main() {
	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Movie."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	
	db := config.ConnectToDatabase()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	r := routes.SetUpRouter(db)
 	 r.Run()
}
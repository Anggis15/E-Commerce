package routes

import (
	"FP-BDS-Sanbercode-Go-50-anggi/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
)

func SetUpRouter(db *gorm.DB) *gin.Engine{
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	r.GET("/users", controller.GetAllUsers)
	r.POST("/users", controller.CreateUser)
	r.GET("/users/:id", controller.GetUserById)
	r.PATCH("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
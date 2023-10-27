package routes

import (
	"FP-BDS-Sanbercode-Go-50-anggi/controller"
	"FP-BDS-Sanbercode-Go-50-anggi/middlewares"

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

	// User
	r.GET("/users", controller.GetAllUsers)
	r.POST("/user/login", controller.Login)
	r.POST("/users/registrasi", controller.CreateUser)
	// r.GET("/users/:id", controller.GetUserById)
	// r.PATCH("/users/:id", controller.UpdateUser)
	// r.DELETE("/users/:id", controller.DeleteUser)

	userMiddleWare := r.Group("/user")
	userMiddleWare.Use(middlewares.JwtAuthMiddleware())
	userMiddleWare.GET("/users/:id", controller.GetUserById)
	userMiddleWare.PATCH("/users/:id", controller.UpdateUser)
	userMiddleWare.DELETE("/users/:id", controller.DeleteUser)

	// Products
	r.GET("/product", controller.GetAllProduct)
	// r.POST("/product", controller.CreateProduct)
	r.GET("/product/:id", controller.GetProductById)
	// r.PATCH("/product/:id", controller.UpdateProduct)
	// r.DELETE("/product/:id", controller.DeleteProduct)

	productMiddleWare := r.Group("/product")
	productMiddleWare.Use(middlewares.JwtAuthMiddleware())
	productMiddleWare.POST("/product", controller.CreateProduct)
	productMiddleWare.PATCH("/product/:id", controller.UpdateProduct)
	productMiddleWare.DELETE("/product/:id", controller.DeleteProduct)

	// Transaction

	// r.GET("/transaction/user/:id", controller.GetAllTransaction)
	// r.POST("/transaction", controller.CreateTransaction)
	// r.GET("/transaction/:id", controller.GetTransactionByUserId)
	// r.DELETE("/transaction/:id", controller.DeleteTransation)

	transactionMiddleWare := r.Group("/transaction")
	transactionMiddleWare.Use(middlewares.JwtAuthMiddleware())
	transactionMiddleWare.GET("/transaction/user/:id", controller.GetAllTransaction)
	transactionMiddleWare.POST("/transaction", controller.CreateTransaction)
	transactionMiddleWare.GET("/transaction/:id", controller.GetTransactionByUserId)
	transactionMiddleWare.DELETE("/transaction/:id", controller.DeleteTransation)

	// Comment
	r.GET("/comment/{productsId}", controller.GetAllComments)
	// r.POST("/comment", controller.CreateComment)
	// r.GET("/comment/user/:useridcomment", controller.GetCommentByUserId)
	// r.POST("/comment/:id", controller.UpdateComment)
	r.DELETE("/comment/:id", controller.DeleteCommet)

	commentMiddleWare := r.Group("/comment")
	commentMiddleWare.Use(middlewares.JwtAuthMiddleware())
	commentMiddleWare.POST("/comment", controller.CreateComment)
	commentMiddleWare.GET("/comment/user/:useridcomment", controller.GetCommentByUserId)
	commentMiddleWare.POST("/comment/:id", controller.UpdateComment)
	commentMiddleWare.DELETE("/comment/:id", controller.DeleteCommet)
	
	commentMiddleWare.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
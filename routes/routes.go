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
	r.POST("/users/login", controller.Login)
	r.POST("/users/registrasi", controller.CreateUser)
	// r.GET("/users/:id", controller.GetUserById)
	// r.PATCH("/users/:id", controller.UpdateUser)
	// r.DELETE("/users/:id", controller.DeleteUser)

	userMiddleWare := r.Group("/users")
	userMiddleWare.Use(middlewares.JwtAuthMiddleware())
	userMiddleWare.GET("/:id", controller.GetUserById)
	userMiddleWare.PATCH("/:id", controller.UpdateUser)
	userMiddleWare.DELETE("/:id", controller.DeleteUser)

	// Products
	r.GET("/product", controller.GetAllProduct)
	// r.POST("/product", controller.CreateProduct)
	r.GET("/product/:id", controller.GetProductById)
	// r.PATCH("/product/:id", controller.UpdateProduct)
	// r.DELETE("/product/:id", controller.DeleteProduct)

	productMiddleWare := r.Group("/product")
	productMiddleWare.Use(middlewares.JwtAuthMiddleware())
	productMiddleWare.POST("/", controller.CreateProduct)
	productMiddleWare.PATCH("/:id", controller.UpdateProduct)
	productMiddleWare.DELETE("/:id", controller.DeleteProduct)

	// Transaction

	// r.GET("/transaction/user/:id", controller.GetAllTransaction)
	// r.POST("/transaction", controller.CreateTransaction)
	// r.GET("/transaction/:id", controller.GetTransactionByUserId)
	// r.DELETE("/transaction/:id", controller.DeleteTransation)

	transactionMiddleWare := r.Group("/transaction")
	transactionMiddleWare.Use(middlewares.JwtAuthMiddleware())
	transactionMiddleWare.GET("/user/:id", controller.GetAllTransaction)
	transactionMiddleWare.POST("/", controller.CreateTransaction)
	transactionMiddleWare.GET("/:id", controller.GetTransactionByUserId)
	transactionMiddleWare.DELETE("/:id", controller.DeleteTransation)

	// Comment
	r.GET("/comment/{productsId}", controller.GetAllComments)
	// r.POST("/comment", controller.CreateComment)
	// r.GET("/comment/user/:useridcomment", controller.GetCommentByUserId)
	// r.POST("/comment/:id", controller.UpdateComment)
	// r.DELETE("/comment/:id", controller.DeleteCommet)

	commentMiddleWare := r.Group("/comment")
	commentMiddleWare.Use(middlewares.JwtAuthMiddleware())
	commentMiddleWare.POST("/", controller.CreateComment)
	commentMiddleWare.GET("/user/:useridcomment", controller.GetCommentByUserId)
	commentMiddleWare.POST("/:id", controller.UpdateComment)
	commentMiddleWare.DELETE("/:id", controller.DeleteCommet)

	r.GET("/promo", controller.GetAllPromotion)
	r.GET("/promo/:id", controller.GetPromoById)
	promoMiddleWare := r.Group("/promo")
	promoMiddleWare.Use(middlewares.JwtAuthMiddleware())
	promoMiddleWare.POST("/promo", controller.CreatePromotion)
	// promoMiddleWare.GET("/:id", controller.GetPromoById)
	promoMiddleWare.PATCH("/:id", controller.UpdatePromotion)
	promoMiddleWare.DELETE("/:id", controller.DeletePromotion)
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
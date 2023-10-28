package controller

import (
	"FP-BDS-Sanbercode-Go-50-anggi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productInput struct {
	UserId      int    `json:"userId"`
	Category    string    `json:"category"`
	Gender  string    `json:"gender"`
	ProductName    string    `json:"productName"`
	Price int64 `json:"price"`
	DescProduct string `json:"descProduct"`
	Stock int `json:"stock"`
	Images string `json:"images"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsActive  int       `json:"isActive"`
}

// GetAllProducts godoc
// @Summary Get all Products.
// @Description Get a list of Products.
// @Tags Products
// @Produce json
// @Success 200 {object} []models.Products
// @Router /product [get]
func GetAllProduct(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var products []models.Products
    db.Find(&products)

    c.JSON(http.StatusOK, gin.H{"data": products})
}

// CreateProducts godoc
// @Summary Create New Product.
// @Description Creating a new Products.
// @Tags Products
// @Param Body body productInput true "the body to create a new Products"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Products
// @Router /product [post]
func CreateProduct(c *gin.Context) {
    // Validate input
    var input productInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create product
    product := models.Products{UserId: input.UserId, Category: input.Category, Gender: input.Gender, 
							ProductName: input.ProductName,
							Price: input.Price,
							DescProduct: input.DescProduct,
							Stock: input.Stock,
							Images: input.Images,
							CreatedAt: time.Now(),
							UpdatedAt: time.Now(),
							IsActive: 1,}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&product)

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// GetProductById godoc
// @Summary Get product.
// @Description Get an Product by id.
// @Tags Products
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} models.Products
// @Router /product/{id} [get]
func GetProductById(c *gin.Context) { // Get model if exist
    var product models.Products

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateProduct godoc
// @Summary Update Product.
// @Description Update Product by id.
// @Tags Products
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Product id"
// @Param Body body productInput true "the body to update product"
// @Success 200 {object} models.Products
// @Router /product/{id} [patch]
func UpdateProduct(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var product models.Products
    if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input productInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Products
    updatedInput.Category = input.Category
	updatedInput.Gender = input.Gender
	updatedInput.ProductName = input.ProductName
	updatedInput.Price = input.Price
	updatedInput.DescProduct = input.DescProduct
	updatedInput.Stock = input.Stock
	updatedInput.Images = input.Images
	updatedInput.IsActive = input.IsActive
	updatedInput.UpdatedAt = time.Now()

    db.Model(&product).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct godoc
// @Summary Delete one Product.
// @Description Delete a Product by id.
// @Tags Products
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} map[string]boolean
// @Router /product/{id} [delete]
func DeleteProduct(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var product models.Products
    if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&product)

    c.JSON(http.StatusOK, gin.H{"data": true})
}
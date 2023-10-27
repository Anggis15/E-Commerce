package controller

import (
	"FP-BDS-Sanbercode-Go-50-anggi/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type transactionInput struct {
	Name      string    `json:"name"`
	UserId     int    `json:"userId"`
	ProductsId  int    `json:"productsId"`
	Quantity   int       `json:"quantity"`
	TotalPrice string    `json:"totalPrice"`
	CreatedAt  time.Time `json:"createdAt"`
	Payment string `json:"payment"`
}

// GetAllTransaction godoc
// @Summary Get all Transaction By UserId.
// @Description Get a list of transaction.
// @Tags Transaction
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} []models.Transaction
// @Router /transaction/{userId} [get]
func GetAllTransaction(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var transaction []models.Transaction
    if err := db.Where("user_id = ?", c.Param("userId")).Find(&transaction).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": transaction})
}

// CreateTransaction godoc
// @Summary Create New Transaction.
// @Description Creating a new Transaction.
// @Tags Transaction
// @Param Body body transactionInput true "the body to create a new comment"
// @Produce json
// @Success 200 {object} models.Transaction
// @Router /transaction [post]
func CreateTransaction(c *gin.Context) {
    // Validate input
    var input transactionInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create transaction
	db := c.MustGet("db").(*gorm.DB)
    user := models.Transaction{
		UserId: input.UserId,
		ProductsId: input.ProductsId,
		Quantity: input.Quantity,
		TotalPrice: strconv.Itoa(input.Quantity * 1000),
		CreatedAt: time.Now(),
		Payment: input.Payment,
	}
    
    db.Create(&user)

    c.JSON(http.StatusOK, gin.H{"data": user})
}

// GetTransactionByuserId godoc
// @Summary Get transaction.
// @Description Get an Transaction by userId.
// @Tags Transaction
// @Produce json
// @Param id path int true "User id"
// @Success 200 {object} models.Transaction
// @Router /transaction/{id} [get]
func GetTransactionByUserId(c *gin.Context) { // Get model if exist
    var transaction models.Transaction

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&transaction).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": transaction})
}


// DeleteTransaction godoc
// @Summary Delete one Trannsaction.
// @Description Delete a User by user id.
// @Tags Transaction
// @Produce json
// @Param id path int true "User id"
// @Success 200 {object} map[string]boolean
// @Router /transation/{id} [delete]
func DeleteTransation(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var transaction models.Transaction
    if err := db.Where("id = ?", c.Param("id")).First(&transaction).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&transaction)

    c.JSON(http.StatusOK, gin.H{"data": true})
}
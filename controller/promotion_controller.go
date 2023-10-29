package controller

import (
	"FP-BDS-Sanbercode-Go-50-anggi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type promotionInput struct {
	ProductsId int    `json:"productsId"`
	PromoName string    `json:"promoName"`
	PromoDesc string    `json:"promoDesc"`
	Promo     int       `json:"promo"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsActive  int       `json:"isActive"`
}

// GetAllPromotion godoc
// @Summary Get all Promotion.
// @Description Get a list of Promotion.
// @Tags Promotion
// @Produce json
// @Success 200 {object} []models.Promotion
// @Router /promotion [get]
func GetAllPromotion(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var promotion []models.Promotion
    db.Find(&promotion)

    c.JSON(http.StatusOK, gin.H{"data": promotion})
}

// CreatePromotion godoc
// @Summary Create New Promotion.
// @Description Creating a new Promotion.
// @Tags Promotion
// @Param Body body promotionInput true "the body to create a new promotion"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Promotion
// @Router /promotion [post]
func CreatePromotion(c *gin.Context) {
    // Validate input
    var input promotionInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create user
    promo := models.Promotion{
		ProductsId: input.ProductsId,
		PromoName: input.PromoName,
		PromoDesc: input.PromoDesc,
		Promo: input.Promo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive: 1,
	}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&promo)

    c.JSON(http.StatusOK, gin.H{"data": promo})
}

// GetpromoById godoc
// @Summary Get promo.
// @Description Get a promo by id.
// @Tags Promotion
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} models.Promotion
// @Router /promotion/{id} [get]
func GetPromoById(c *gin.Context) { // Get model if exist
    var promo models.Promotion

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("products_id = ?", c.Param("id")).First(&promo).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": promo})
}

// UpdateIsActive godoc
// @Summary Update Active Promo.
// @Description Update User by id.
// @Tags Promotion
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "User id"
// @Param Body body promotionInput true "the body to update promo"
// @Success 200 {object} models.Promotion
// @Router /promotion/{id} [patch]
func UpdatePromotion(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var promo models.Promotion
    if err := db.Where("id = ?", c.Param("id")).First(&promo).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input promotionInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Promotion
    updatedInput.IsActive = input.IsActive
	updatedInput.UpdatedAt = time.Now()

    db.Model(&promo).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": promo})
}

// DeletePromotion godoc
// @Summary Delete one Promotion.
// @Description Delete a Promotion by id.
// @Tags Promotion
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path int true "User id"
// @Success 200 {object} map[string]boolean
// @Router /promotion/{id} [delete]
func DeletePromotion(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var promo models.Promotion
    if err := db.Where("id = ?", c.Param("id")).First(&promo).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&promo)

    c.JSON(http.StatusOK, gin.H{"data": true})
}
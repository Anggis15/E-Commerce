package controller

import (
	"FP-BDS-Sanbercode-Go-50-anggi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentInput struct {
	ProductsId     int    `json:"productsId"`
	UserIdComment int    `json:"userIdComment"`
	Comment       string    `json:"comment"`
	Rating        int       `json:"rating"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// GetAllCommentinProduct godoc
// @Summary Get all Comments.
// @Description Get a list of Comments in Products.
// @Tags Comments
// @Produce json
// @Param id path int true "products id"
// @Success 200 {object} []models.Comments
// @Router /comment/{productsId} [get]
func GetAllComments(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var comments []models.Comments
    if err := db.Where("product_id = ?", c.Param("productsiId")).Find(&comments).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": comments})
}

// CreateComment godoc
// @Summary Create New Comment.
// @Description Creating a new Comment.
// @Tags Comments
// @Param Body body commentInput true "the body to create a new Comments"
// @Produce json
// @Success 200 {object} models.Comments
// @Router /comment [post]
func CreateComment(c *gin.Context) {
    // Validate input
    var input commentInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create comment
    comment := models.Comments{
		ProductsId: input.ProductsId,
		UserIdComment: input.UserIdComment,
		Comment: input.Comment,
		Rating: input.Rating,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&comment)

    c.JSON(http.StatusOK, gin.H{"data": comment})
}

// GetCommentByUserIdComment godoc
// @Summary Get comment.
// @Description Get an Comment by UserIdComment.
// @Tags Comments
// @Produce json
// @Param id path int true "User id comment"
// @Success 200 {object} models.Comments
// @Router /comment/user/{useridcomment} [get]
func GetCommentByUserId(c *gin.Context) { // Get model if exist
    var comments []models.Comments

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("user_id_comment = ?", c.Param("useridcomment")).Find(&comments).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": comments})
}

// UpdateComment godoc
// @Summary Update Comment.
// @Description Update Comment by id.
// @Tags Comments
// @Produce json
// @Param id path int true "id"
// @Param Body body commentInput true "the body to update comment"
// @Success 200 {object} models.Comments
// @Router /comment/{id} [patch]
func UpdateComment(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var comment models.Comments
    if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input commentInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Comments
    updatedInput.Comment = input.Comment
	updatedInput.UpdatedAt = time.Now()

    db.Model(&comment).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": comment})
}

// DeleteCommet godoc
// @Summary Delete one Comment.
// @Description Delete a Comment by id.
// @Tags Comment
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} map[string]boolean
// @Router /comment/{id} [delete]
func DeleteCommet(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var comment models.Comments
    if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&comment)

    c.JSON(http.StatusOK, gin.H{"data": true})
}
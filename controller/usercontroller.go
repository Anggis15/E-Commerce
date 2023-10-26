package controller

import (
	"FP-BDS-Sanbercode-Go-50-anggi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userInput struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsActive  int       `json:"isActive"`
}

// GetAllUsers godoc
// @Summary Get all Users.
// @Description Get a list of Users.
// @Tags Users
// @Produce json
// @Success 200 {object} []models.User
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var users []models.User
    db.Find(&users)

    c.JSON(http.StatusOK, gin.H{"data": users})
}

// CreateAgeRatingCategory godoc
// @Summary Create New User.
// @Description Creating a new User.
// @Tags User
// @Param Body body userInput true "the body to create a new users"
// @Produce json
// @Success 200 {object} models.User
// @Router /users [post]
func CreateUser(c *gin.Context) {
    // Validate input
    var input userInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create user
    user := models.User{Name: input.Name, Email: input.Email, Password: input.Password, Gender: input.Gender, CreatedAt: time.Now(), UpdatedAt: time.Now(), IsActive: 1}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&user)

    c.JSON(http.StatusOK, gin.H{"data": user})
}

// GetUserById godoc
// @Summary Get user.
// @Description Get an User by id.
// @Tags User
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func GetUserById(c *gin.Context) { // Get model if exist
    var user models.User

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUserPassword godoc
// @Summary Update User.
// @Description Update User by id.
// @Tags User
// @Produce json
// @Param id path string true "User id"
// @Param Body body userInput true "the body to update user"
// @Success 200 {object} models.User
// @Router /users/{id} [patch]
func UpdateUser(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var user models.User
    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input userInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.User
    updatedInput.Password = input.Password
	updatedInput.UpdatedAt = time.Now()

    db.Model(&user).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser godoc
// @Summary Delete one User.
// @Description Delete a User by id.
// @Tags User
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} map[string]boolean
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var user models.User
    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&user)

    c.JSON(http.StatusOK, gin.H{"data": true})
}
package controllers

import (
	"net/http"
	"root/config"
	"root/models"
	"root/validators"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)

	c.JSON(200, &users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	result := config.DB.Where("email = ?", user.Email).First(&user)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User with this email already exist"})
		return
	}

	bool, string := validators.ValidateEmail(user.Email)
	if user.Email == "" || !bool {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email", "error": string})
		return
	}
	validpw, string := validators.ValidatePassword(user.Password, user.Email)
	if !validpw {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid password", "error": string})
		return
	}
	validperm := models.IsValidPermessi(user.Permessi)
	if !validperm {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid permission"})
		return
	}
	hashedPassword, err := validators.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword
	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	return
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id=?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id=?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}

package handler

import (
	"net/http"

	"go-gin-crud/config"
	"go-gin-crud/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GET /users
func GetUsers(c *gin.Context) {
	var users []model.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GET /users/:id
func GetUserByID(c *gin.Context) {
	var user model.User

	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// PUT /users/:id
func UpdateUser(c *gin.Context) {
	var user model.User

	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ValidationError(err))
		return
	}

	config.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	if err := config.DB.Delete(&model.User{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, ValidationError(err))
		return
	}

	c.Status(http.StatusNoContent)
}

func ValidationError(err error) gin.H {
	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = e.Tag()
	}

	return gin.H{"errors": errors}
}

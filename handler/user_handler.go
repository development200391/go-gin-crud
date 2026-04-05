package handler

import (
	"errors"
	"net/http"
	"time"

	"go-gin-crud/config"
	"go-gin-crud/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type updateUserRequest struct {
	Name     string `json:"name" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"omitempty,min=6"`
	Role     string `json:"role"`
}

// GET /users
func GetUsers(c *gin.Context) {
	var users []model.User
	config.DB.Where("is_active = ?", true).Find(&users)
	c.JSON(http.StatusOK, users)
}

// GET /users/:id
func GetUserByID(c *gin.Context) {
	var user model.User

	if err := config.DB.Where("id = ? AND is_active = ?", c.Param("id"), true).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// PUT /users/:id
func UpdateUser(c *gin.Context) {
	var user model.User
	var req updateUserRequest

	if err := config.DB.Where("id = ? AND is_active = ?", c.Param("id"), true).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ValidationError(err))
		return
	}

	updates := map[string]interface{}{
		"name":  req.Name,
		"email": req.Email,
		"role":  req.Role,
	}

	if req.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		updates["password"] = string(hashed)
	}

	if err := config.DB.Model(&user).Updates(updates).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.First(&user, user.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	deletedAt := time.Now()

	result := config.DB.Model(&model.User{}).
		Where("id = ? AND is_active = ?", c.Param("id"), true).
		Updates(map[string]interface{}{
			"is_active":  false,
			"deleted_at": &deletedAt,
		})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

func ValidationError(err error) gin.H {
	var validationErrors validator.ValidationErrors
	errors := make(map[string]string)

	if !AsValidationErrors(err, &validationErrors) {
		return gin.H{"error": err.Error()}
	}

	for _, e := range validationErrors {
		errors[e.Field()] = e.Tag()
	}

	return gin.H{"errors": errors}
}

func AsValidationErrors(err error, target *validator.ValidationErrors) bool {
	return errors.As(err, target)
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"go-gin-crud/config"
	"go-gin-crud/dto"
	"go-gin-crud/model"
	"go-gin-crud/utils"
)

type registerRequest struct {
	Name     string `json:"name" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	RoleID   *uint  `json:"role_id"`
}

// POST /register
func Register(c *gin.Context) {
	var req registerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ValidationError(err))
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashed),
		IsActive: true,
		RoleID:   req.RoleID,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"is_active":  user.IsActive,
		"role_id":    user.RoleID,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
}

// POST /login
func Login(c *gin.Context) {
	var req dto.LoginRequest
	var user model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Where("email = ? AND is_active = ?", req.Email, true).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

package handler

import (
	"net/http"
	"time"

	"go-gin-crud/config"
	"go-gin-crud/model"

	"github.com/gin-gonic/gin"
)

type roleRequest struct {
	RoleName    string `json:"role_name" binding:"required"`
	Description string `json:"description"`
}

func CreateRole(c *gin.Context) {
	var req roleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ValidationError(err))
		return
	}

	role := model.Role{
		RoleName:    req.RoleName,
		Description: req.Description,
		IsActive:    true,
	}

	if err := config.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

func GetRoles(c *gin.Context) {
	var roles []model.Role
	config.DB.Where("is_active = ?", true).Find(&roles)
	c.JSON(http.StatusOK, roles)
}

func GetRoleByID(c *gin.Context) {
	var role model.Role

	if err := config.DB.Where("id = ? AND is_active = ?", c.Param("id"), true).First(&role).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

func UpdateRole(c *gin.Context) {
	var role model.Role
	var req roleRequest

	if err := config.DB.Where("id = ? AND is_active = ?", c.Param("id"), true).First(&role).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ValidationError(err))
		return
	}

	if err := config.DB.Model(&role).Updates(map[string]interface{}{
		"role_name":   req.RoleName,
		"description": req.Description,
	}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.First(&role, role.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

func DeleteRole(c *gin.Context) {
	deletedAt := time.Now()

	result := config.DB.Model(&model.Role{}).
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

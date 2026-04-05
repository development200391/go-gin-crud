package handler

import (
	"net/http"
	"time"

	"go-gin-crud/config"
	"go-gin-crud/model"

	"github.com/gin-gonic/gin"
)

type permissionRequest struct {
	PermName    string `json:"perm_name" binding:"required"`
	Description string `json:"description"`
}

func CreatePermission(c *gin.Context) {
	var req permissionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ValidationError(err))
		return
	}

	permission := model.Permission{
		PermName:    req.PermName,
		Description: req.Description,
		IsActive:    true,
	}

	if err := config.DB.Create(&permission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, permission)
}

func GetPermissions(c *gin.Context) {
	var permissions []model.Permission
	config.DB.Where("is_active = ?", true).Find(&permissions)
	c.JSON(http.StatusOK, permissions)
}

func GetPermissionByID(c *gin.Context) {
	var permission model.Permission

	if err := config.DB.Where("id = ? AND is_active = ?", c.Param("id"), true).First(&permission).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		return
	}

	c.JSON(http.StatusOK, permission)
}

func UpdatePermission(c *gin.Context) {
	var permission model.Permission
	var req permissionRequest

	if err := config.DB.Where("id = ? AND is_active = ?", c.Param("id"), true).First(&permission).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ValidationError(err))
		return
	}

	if err := config.DB.Model(&permission).Updates(map[string]interface{}{
		"perm_name":   req.PermName,
		"description": req.Description,
	}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.First(&permission, permission.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, permission)
}

func DeletePermission(c *gin.Context) {
	deletedAt := time.Now()

	result := config.DB.Model(&model.Permission{}).
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

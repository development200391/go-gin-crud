package handler

import (
	"net/http"

	"go-gin-crud/config"
	"go-gin-crud/model"

	"github.com/gin-gonic/gin"
)

type rolePermissionRequest struct {
	RoleID uint `json:"role_id" binding:"required"`
	PermID uint `json:"perm_id" binding:"required"`
}

func CreateRolePermission(c *gin.Context) {
	var req rolePermissionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ValidationError(err))
		return
	}

	rolePermission := model.RolePermission{
		RoleID: req.RoleID,
		PermID: req.PermID,
	}

	if err := config.DB.Create(&rolePermission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, rolePermission)
}

func GetRolePermissions(c *gin.Context) {
	var rolePermissions []model.RolePermission
	config.DB.Find(&rolePermissions)
	c.JSON(http.StatusOK, rolePermissions)
}

func GetRolePermissionByID(c *gin.Context) {
	var rolePermission model.RolePermission

	if err := config.DB.Where("role_id = ? AND perm_id = ?", c.Param("role_id"), c.Param("perm_id")).First(&rolePermission).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role permission not found"})
		return
	}

	c.JSON(http.StatusOK, rolePermission)
}

func UpdateRolePermission(c *gin.Context) {
	var existing model.RolePermission
	var req rolePermissionRequest

	if err := config.DB.Where("role_id = ? AND perm_id = ?", c.Param("role_id"), c.Param("perm_id")).First(&existing).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role permission not found"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ValidationError(err))
		return
	}

	if err := config.DB.Model(&existing).
		Where("role_id = ? AND perm_id = ?", existing.RoleID, existing.PermID).
		Updates(map[string]interface{}{
			"role_id": req.RoleID,
			"perm_id": req.PermID,
		}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Where("role_id = ? AND perm_id = ?", req.RoleID, req.PermID).First(&existing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existing)
}

func DeleteRolePermission(c *gin.Context) {
	result := config.DB.Where("role_id = ? AND perm_id = ?", c.Param("role_id"), c.Param("perm_id")).Delete(&model.RolePermission{})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role permission not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

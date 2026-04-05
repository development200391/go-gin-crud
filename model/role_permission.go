package model

import "time"

type RolePermission struct {
	RoleID    uint      `json:"role_id" binding:"required" gorm:"primaryKey;column:role_id"`
	PermID    uint      `json:"perm_id" binding:"required" gorm:"primaryKey;column:perm_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

func (RolePermission) TableName() string {
	return "auth.role_permissions"
}

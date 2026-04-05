package model

import "time"

type Role struct {
	ID          uint       `json:"id" gorm:"primaryKey;column:id"`
	RoleName    string     `json:"role_name" binding:"required" gorm:"column:role_name;type:varchar(50);not null;unique"`
	Description string     `json:"description" gorm:"column:description;type:text"`
	IsActive    bool       `json:"is_active" gorm:"column:is_active;default:true"`
	CreatedAt   time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}

func (Role) TableName() string {
	return "auth.roles"
}

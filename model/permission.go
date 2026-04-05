package model

import "time"

type Permission struct {
	ID          uint       `json:"id" gorm:"primaryKey;column:id"`
	PermName    string     `json:"perm_name" binding:"required" gorm:"column:perm_name;type:varchar(100);not null;unique"`
	Description string     `json:"description" gorm:"column:description;type:text"`
	IsActive    bool       `json:"is_active" gorm:"column:is_active;default:true"`
	CreatedAt   time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}

func (Permission) TableName() string {
	return "permissions"
}

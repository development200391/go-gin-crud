package model

import "time"

type User struct {
	ID        uint       `json:"id" gorm:"primaryKey;column:id"`
	Name      string     `json:"name" binding:"required,min=3" gorm:"column:name;type:varchar(100)"`
	Email     string     `json:"email" binding:"required,email" gorm:"column:email;type:varchar(150)"`
	Password  string     `json:"password,omitempty" binding:"required,min=6" gorm:"column:password;type:varchar(255)"`
	IsActive  bool       `json:"is_active" gorm:"column:is_active;default:true"`
	RoleID    *uint      `json:"role_id" gorm:"column:role_id"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "auth.users"
}

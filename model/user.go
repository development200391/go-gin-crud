package model

import "time"

type User struct {
	ID        uint       `json:"id" gorm:"primaryKey;column:id"`
	Name      string     `json:"name" binding:"required,min=3" gorm:"column:name;type:varchar(100);not null"`
	Email     string     `json:"email" binding:"required,email" gorm:"column:email;type:varchar(150);uniqueIndex:uni_users_email;not null"`
	Password  string     `json:"password,omitempty" binding:"required,min=6" gorm:"column:password;type:varchar(255);not null"`
	IsActive  bool       `json:"is_active" gorm:"column:is_active;default:true"`
	Role      string     `json:"role" gorm:"column:role;type:varchar(20);default:user"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}

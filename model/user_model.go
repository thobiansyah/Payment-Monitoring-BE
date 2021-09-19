package model

import "gorm.io/gorm"

//entity
type User struct {
	gorm.Model
	ID       uint   `gorm:"not null"`
	Name     string `gorm:"not null"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	RoleID   uint   `gorm:"not null"`
	Role     Role
}

//request

//response
type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     Role   `json:"role"`
}

package model

import "gorm.io/gorm"

//entity
type Role struct {
	gorm.Model
	ID   uint   `gorm:"not null"`
	Name string `gorm:"not null"`
}

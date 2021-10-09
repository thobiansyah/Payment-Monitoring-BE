package model

import (
	"time"

	"gorm.io/gorm"
)

//entity
type Payment struct {
	gorm.Model
	ID                   uint      `gorm:"not null"`
	UserID               uint      `gorm:"not null"`
	RequestBy            string    `gorm:"size:255;not null"`
	Necessity            string    `gorm:"size:255;not null"`
	PaymentDate          time.Time `gorm:"not null"`
	PaymentAmount        string    `gorm:"not null"`
	PaymentCalculate     string    `gorm:"not null"`
	PaymentAccountName   string    `gorm:"size:255;not null"`
	PaymentAccountNumber string    `gorm:"size:255;not null"`
	StatusID             uint      `gorm:"not null"`
	Reason               *string   `gorm:"default:null"`
	User                 User
	Status               Status
}

package main

import (
	"backend-c-payment-monitoring/config"
	"backend-c-payment-monitoring/model"
	"log"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestMigration(t *testing.T) {
	//setup configuration
	configuration := config.New()
	database := config.NewMysqlDatabase(configuration)

	database.AutoMigrate(&model.Role{}, &model.User{})
}

func TestRoleSeeder(t *testing.T) {
	//setup configuration
	configuration := config.New()
	database := config.NewMysqlDatabase(configuration)

	var roles = []model.Role{
		{Name: "Admin"},
		{Name: "Unit Kerja (Customer)"},
		{Name: "General Support"},
		{Name: "Accounting"},
	}

	err := database.Create(&roles).Error

	if err != nil {
		log.Println("Role Seed Failed")
	}

	log.Println("Role Seed Success")
}

func TestUserSeeder(t *testing.T) {
	//setup configuration
	configuration := config.New()
	database := config.NewMysqlDatabase(configuration)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.MinCost)

	var user = model.User{
		Name:     "Admin",
		Username: "admin",
		Password: string(passwordHash),
		RoleID:   1,
	}

	err = database.Create(&user).Error

	if err != nil {
		log.Println("User Seed Failed")
	}

	log.Println("User Seed Success")

}

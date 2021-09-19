package repository

import (
	"backend-c-payment-monitoring/config"
	"backend-c-payment-monitoring/model"
)

func FindUserByUsername(username string) (model.User, error) {

	configuration := config.New()
	db := config.NewMysqlDatabase(configuration)

	var user model.User
	err := db.Preload("Role").Where("username = ?", username).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

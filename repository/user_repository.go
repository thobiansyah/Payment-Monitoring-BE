package repository

import (
	"backend-c-payment-monitoring/config"
	"backend-c-payment-monitoring/model"

	"golang.org/x/crypto/bcrypt"
)

func FindAllUser(pagination model.Pagination) (model.Pagination, error) {

	configuration := config.New()
	db := config.NewMysqlDatabase(configuration)

	var users []model.User
	keyword := "%" + pagination.Keyword + "%"
	err := db.
		Preload("Role").
		Where("name LIKE ?", keyword).
		Scopes(model.UserPaginate(keyword, users, &pagination, db)).
		Find(&users).Error

	if err != nil {
		return pagination, err
	}

	pagination.Rows = model.FormatGetAllUserResponse(users)

	return pagination, nil
}

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

func DeleteUser(id int) (model.User, error) {

	configuration := config.New()
	db := config.NewMysqlDatabase(configuration)

	var user model.User

	err := db.Where("id = ?", id).Delete(&user).Error
  
  if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(payload model.User) (model.User, error) {
	configuration := config.New()
	db := config.NewMysqlDatabase(configuration)

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	var user = model.User{
		Name:     payload.Name,
		Username: payload.Username,
		Password: string(passwordHash),
		RoleID:   payload.RoleID,
	}

	err := db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

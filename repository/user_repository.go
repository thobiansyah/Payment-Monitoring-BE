package repository

import (
	"backend-c-payment-monitoring/config"
	"backend-c-payment-monitoring/model"
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

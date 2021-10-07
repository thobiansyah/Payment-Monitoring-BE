package service

import (
	"backend-c-payment-monitoring/model"
	"backend-c-payment-monitoring/repository"
)

func GetAllUser(pagination model.Pagination) (model.Pagination, error) {

	users, err := repository.FindAllUser(pagination)

	if err != nil {
		return users, err
	}

	return users, err
}

func CreateUser(user model.User) (model.User, error) {
	inserted, err := repository.CreateUser(user)

	if err != nil {
		return model.User{}, err
	}

	return inserted, nil
}

func DeleteUser(id int) bool {
	//check availability
	_, err := repository.FindUserById(id)

	if err != nil {
		return false
	}

	_, err = repository.DeleteUser(id)

	if err != nil {
		return false
	}

	return true
}

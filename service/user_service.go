package service

import (
	"backend-c-payment-monitoring/model"
	"backend-c-payment-monitoring/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GetAllUser(pagination model.Pagination) (model.Pagination, error) {

	users, err := repository.FindAllUser(pagination)

	if err != nil {
		return users, err
	}

	return users, err
}

func UpdateUser(id int, payload model.User) (int, error) {

	user, err := repository.FindUserById(id)

	if err != nil {
		return id, errors.New("Id Not Found")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)

	user.Name = payload.Name
	user.Username = payload.Username
	user.Password = string(password)
	user.RoleID = payload.RoleID

	user, errUpdate := repository.SaveUser(user)
	if errUpdate != nil {
		return id, errUpdate
	}

	return id, nil
}

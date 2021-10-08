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

func CreateUser(payload model.User) (model.User, error) {
	//check username
	username, err := repository.FindUserByUsername(payload.Username)

	if err == nil {
		return username, errors.New("Username already exist")
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	var user = model.User{
		Name:     payload.Name,
		Username: payload.Username,
		Password: string(passwordHash),
		RoleID:   payload.RoleID,
	}

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

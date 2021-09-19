package service

import (
	"backend-c-payment-monitoring/model"
	"backend-c-payment-monitoring/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Login(input model.LoginUserRequest) (model.User, error) {

	username := input.Username
	password := input.Password

	user, err := repository.FindUserByUsername(username)

	if err != nil {
		return user, errors.New("Username Not Found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("Username/Password was wrong")
	}

	return user, nil
}

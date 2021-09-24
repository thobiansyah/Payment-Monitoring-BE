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

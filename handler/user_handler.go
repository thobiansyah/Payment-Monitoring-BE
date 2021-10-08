package handler

import (
	"backend-c-payment-monitoring/exception"
	"backend-c-payment-monitoring/model"
	"backend-c-payment-monitoring/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.FormValue("limit"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString("limit required."),
			Data:    nil,
		})
	}

	page, err := strconv.Atoi(c.FormValue("page"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString("page required."),
			Data:    nil,
		})
	}

	keyword := c.FormValue("keyword")

	set_paginate := model.Pagination{}
	set_paginate.Limit = limit
	set_paginate.Page = page
	set_paginate.Keyword = keyword
	set_paginate.Sort = "Id asc"

	responses, err := service.GetAllUser(set_paginate)

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(model.ApiResponse{
		Code:    http.StatusOK,
		Message: "Get Data Success",
		Error:   nil,
		Data:    model.Pagination(responses),
	})
}

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := service.GetUserById(id)

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(model.ApiResponse{
		Code:    http.StatusOK,
		Message: "Get Data Success",
		Error:   nil,
		Data:    responses,
	})

}

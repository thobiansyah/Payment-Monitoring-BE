package handler

import (
	"backend-c-payment-monitoring/exception"
	"backend-c-payment-monitoring/exception/validation"
	"backend-c-payment-monitoring/model"
	"backend-c-payment-monitoring/service"
	"backend-c-payment-monitoring/util"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(ctx *fiber.Ctx) error {
	var input model.LoginUserRequest
	err := ctx.BodyParser(&input)
	if err != nil {
		return err
	}

	//validation
	validation.LoginValidate(input)

	responses, err := service.Login(input)

	if err != nil {
		//error
		return ctx.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Login Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	token, err := util.GenerateNewAccessToken(responses)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(model.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Login Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(model.ApiResponse{
		Code:    http.StatusOK,
		Message: "Login Success",
		Error:   nil,
		Data:    model.FormatLoginUserResponse(responses, token),
	})
}

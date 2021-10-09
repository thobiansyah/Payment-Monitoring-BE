package main

import (
	"backend-c-payment-monitoring/config"
	"backend-c-payment-monitoring/handler"
	"backend-c-payment-monitoring/middleware"
	"backend-c-payment-monitoring/model"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configuration := config.New()
	database := config.NewMysqlDatabase(configuration)

	//migration
	database.AutoMigrate(&model.Role{}, &model.User{}, &model.Status{}, &model.Payment{})

	app := fiber.New(config.NewFiberConfig())
	setupRoutes(app)

	port := configuration.Get("APP_PORT")
	app.Listen(fmt.Sprintf(":%v", port))
}

func setupRoutes(app *fiber.App) {
	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api := app.Group("/api/v1")

	api.Post("/login", handler.LoginHandler)

	api.Get("/users", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.GetAllUser)
	api.Get("/users/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.GetUserById)
	api.Post("/users", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.CreateUser)
	api.Put("/users/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.UpdateUser)
	api.Delete("/users/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.DeleteUser)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(model.ApiResponse{
			Code:  http.StatusNotFound,
			Error: &fiber.ErrNotFound.Message,
			Data:  nil,
		})
	})
}

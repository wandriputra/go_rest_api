package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wandriputra/go_rest_api/src/controllers/rest"
	"github.com/wandriputra/go_rest_api/src/pkg/database"
	"github.com/wandriputra/go_rest_api/src/services"
)

func main() {
	// The main function is the entry point of the application
	database.Connect()

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:   "${time} ${method} ${path} ${status} ${latency}\n",
		TimeZone: "Local",
	}))
	app.Use(cors.New())

	userService := services.NewUserService()

	rest.InitRestUser(app, userService)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found",
		})
	})

	app.Listen(":3000")

}

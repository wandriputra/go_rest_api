package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wandriputra/go_rest_api/src/domains"
	"github.com/wandriputra/go_rest_api/src/pkg/utils"
)

type App struct {
	Service domains.IUserService
}

func InitRestUser(app *fiber.App, service domains.IUserService) {
	rest := App{Service: service}

	api := app.Group("/api")
	user := api.Group("/user")

	user.Use(func(c *fiber.Ctx) error {
		fmt.Println("Middleware")
		return c.Next()
	})

	user.Get("/", rest.CreateUser)
	user.Get("/:id", rest.CreateUser)
	user.Post("/", rest.CreateUser)
}

func (handler *App) CreateUser(c *fiber.Ctx) error {
	response, err := handler.Service.GetAllUsers(c.UserContext())
	utils.PanicIfNeeded(err)

	return c.JSON(utils.ResponseData{
		Status:  200,
		Code:    "SUCCESS",
		Message: "Login success",
		Results: response,
	})
}

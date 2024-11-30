package handler

import "github.com/gofiber/fiber/v2"

func GetAllUsers(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "All users",
	})
}

func GetSingleUser(c *fiber.Ctx) error {
	return c.SendString("Single user")
}

func CreateUser(c *fiber.Ctx) error {
	return c.SendString("Create user")
}

func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Update user")
}

func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Delete user")
}

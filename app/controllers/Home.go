package controllers

import (
	"github.com/brahimrizqHireme/go-fiber/app/utils"
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Welcome to Golang, Fiber, SQLite, and GORM",
	})
}

func Home(c *fiber.Ctx) error {
	userID := utils.GenerateUUID().String()

	return c.JSON(fiber.Map{
		"message": "User was created! " + userID,
	})
}

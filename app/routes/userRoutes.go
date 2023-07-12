package routes

import (
	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(app fiber.Router) {
	userGroup := app.Group("/user")
	userGroup.Get("/", getUserListHandler)
	// userGroup.Get("/:id", getUserHandler)
	// userGroup.Post("/", createUserHandler)
	// userGroup.Put("/:id", updateUserHandler)
	// userGroup.Delete("/:id", deleteUserHandler)
}

func getUserListHandler(c *fiber.Ctx) error {
	user := c.Locals("user")
	return c.JSON(user)

}

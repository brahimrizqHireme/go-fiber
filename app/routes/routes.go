package routes

import (
	"github.com/brahimrizqHireme/go-fiber/app/controllers"
	"github.com/brahimrizqHireme/go-fiber/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	publicRoutes(app)
	privateRoutes(app)
}

func publicRoutes(app *fiber.App) {
	app.Get("/", controllers.Home)
	app.Get("/api/healthcheck", controllers.HealthCheck)

	// userService := container.Resolve("UserRepository").(*repositories.UserRepository)

	app.Post("/auth", controllers.NewAuthHandler().Login)
	app.Post("/api/register", controllers.NewAuthHandler().Register)
}

func privateRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1", middlewares.NewJwtMiddleware().JWTProtected())
	setupUserRoutes(apiV1)
}

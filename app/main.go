package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	"github.com/brahimrizqHireme/go-fiber/app/configs"
	"github.com/brahimrizqHireme/go-fiber/app/middlewares"
	"github.com/brahimrizqHireme/go-fiber/app/routes"

	"github.com/brahimrizqHireme/go-fiber/app/database"
	"github.com/brahimrizqHireme/go-fiber/app/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(FiberConfig())
	initProject(app)
	utils.StartServer(app)
}

func initProject(app *fiber.App) {
	err := godotenv.Load(".env.app")
	if err != nil {
		log.Fatal("Error loading .env.app file")
	}
	utc := time.FixedZone("UTC", 0)
	time.Local = utc
	database.Connect()
	configs.LoadConfig()
	// container := container.NewAppContainer()
	// container.RegisterServices()
	// app.Use(func(c *fiber.Ctx) error {
	// 	c.Locals("container", container)
	// 	c.Locals("config", config.LoadConfig())
	// 	return c.Next()
	// })

	// client, err := database.MongoDbClient()
	// if err != nil {
	// 	log.Fatal("Error connecting to MongoDB:", err)
	// }
	middlewares.SetupMiddlewares(app)
	routes.SetupRoutes(app)

	// defer client.Close()
}

func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout:  time.Second * time.Duration(readTimeoutSecondsCount),
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: utils.HandleError,
	}
}

package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/brahimrizqHireme/go-fiber/app/configs"
	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := a.Listen(configs.AppConfig.ServerUrl); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

func StartServer(a *fiber.App) {
	if err := a.Listen(configs.AppConfig.ServerUrl); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}

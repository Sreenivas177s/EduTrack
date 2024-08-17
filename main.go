package main

import (
	ApiHandler "chat-server/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {

	log.Info("Starting Server ....")
	app := initAppInstance()
	app.Group("/api", ApiHandler.HandleApiCall)

	// ws := app.Group("/ws/v1")

	log.Fatal(app.Listen(":3000"))
}

func initAppInstance() *fiber.App {
	app := fiber.New()
	app.Static("/ui", "./static")
	return app
}

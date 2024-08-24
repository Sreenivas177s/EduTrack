package main

import (
	ApiHandler "chat-server/handlers"
	"chat-server/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := initAppInstance()

	ApiHandler.HandleApiCall(app.Group(`/api/:version<regex(v\d{1,2})>`))
	// ws := app.Group("/ws/v1")

	// return not found
	app.All("/*", ApiHandler.UrlNotFound)

	log.Fatal(app.Listen(":3000"))
}

func initAppInstance() *fiber.App {
	app := fiber.New()
	app.Static(`/ui`, `./static`)
	app.Use(utils.GetLogger())
	log.Info("Server Initialized ", time.Now(), " ....")

	return app
}

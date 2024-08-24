package main

import (
	ApiHandler "chat-server/handlers"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

func main() {
	// load env variables
	viper.SetConfigFile("./.env")

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
	// app.Use(utils.GetLogger())
	log.Info("Server Initialized ", time.Now(), " ....")

	return app
}

package main

import (
	apiframework "chat-server/api-framework"
	"chat-server/auth"
	"chat-server/database"
	"chat-server/utils"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/joho/godotenv"
)

func main() {
	// load env variables
	if err := godotenv.Load(); err != nil {
		panic("Error while loading env file")
	}
	// initializing data base
	database.InitDataBase()

	app := initAppInstance()

	// init logger
	accessLogger := utils.RegisterAccessLogger(app)
	// authorizing apis
	authRouter := app.Group(`/auth`)
	auth.HandleAuth(authRouter)
	authenticatedRouter := auth.InitAuthMiddleWare(app.Group("/*"))
	// execute apis
	apiframework.HandleApiCall(authenticatedRouter.Group(`/api/:version<regex(v\d{1,2})>`))
	// ws := app.Group("/ws/v1")

	// return not found
	app.All("/*", utils.ServeNotFoundHTML)
	//print all registered routes in configured-routes.json
	utils.GenerateConfiguredRoutesJSON(app)

	defer accessLogger.Close()
	url := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	log.Fatal(app.Listen(url))
}

func initAppInstance() *fiber.App {
	app := fiber.New()
	//helmet middleware to enforce security rules
	app.Use(helmet.New())
	// initialize readiness and liveliness endpoints for healthcheck
	initHealthCheckApis(app)
	// initialize static files
	app.Static(`/ui`, `./static`)
	log.Info("Server Initialized ", time.Now().Format(time.DateTime), " ....")
	return app
}

func initHealthCheckApis(app *fiber.App) {
	config := healthcheck.Config{
		LivenessEndpoint:  "is_live",
		ReadinessEndpoint: "is_ready",
	}
	app.Use(healthcheck.New(config))
}

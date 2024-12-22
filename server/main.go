package main

import (
	api "chat-server/api"
	"chat-server/auth"
	"chat-server/database"
	"chat-server/utils"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/joho/godotenv"
)

func main() {
	// load env variables
	if err := godotenv.Load(filepath.Join("..", ".env")); err != nil {
		panic("Error while loading env file")
	}
	// initializing data base
	database.InitDataBase()
	database.InitializeRedis()

	app := initAppInstance()

	// init logger
	accessLogger := utils.RegisterAccessLogger(app)

	authMiddleware := auth.GetAuthMiddleWare()
	app.Use(authMiddleware)
	// authorizing apis
	auth.HandleAuth(app)
	api.HandleApiCall(app)
	// ws := app.Group("/ws/v1")

	// return not found
	app.All("/*", utils.ServeNotFoundHTML).Name("Not found - Generic")
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

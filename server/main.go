package main

import (
	api "chat-server/api"
	"chat-server/auth"
	"chat-server/database"
	"chat-server/utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func main() {
	// load env variables
	utils.LoadEnv()
	// initializing data base
	database.InitDataBase()
	database.InitializeRedis()

	app := initAppInstance()

	// init logger
	accessLogger := utils.RegisterAccessLogger(app)

	app.Use(auth.GetAuthMiddleWare())
	// authorizing apis
	auth.HandleAuth(app)
	api.HandleApiCall(app)
	// ws := app.Group("/ws/v1")

	// return not found
	app.All("/*", utils.ServeNotFoundHTML).Name("Not found - Generic")
	//print all registered routes in configured-routes.json
	utils.GenerateConfiguredRoutesJSON(app)

	url := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	go func() {
		if err := app.Listen(url); err != nil {
			log.Panic(err)
		}
	}()

	serverChannel := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(serverChannel, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-serverChannel // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	defer accessLogger.Close()
	database.CloseRedis()
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
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

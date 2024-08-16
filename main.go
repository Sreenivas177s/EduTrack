package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()

	app.Static("/", "./static")

	api := app.Group("/api/v1")

	ws := app.Group("/ws/v1")
}

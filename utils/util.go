package utils

import (
	"chat-server/api-framework/entity"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RegisterAccessLogger(app *fiber.App) *os.File {
	// register accesslog
	path := fmt.Sprintf("access-log-%s.txt", time.Now().Format(time.DateOnly))
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		log.Fatalf("error while creating logs directory")
	}
	path = filepath.Join("logs", path)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	config := logger.Config{
		Format: "${time} - ${method} | ${path} | ${status} | ${latency} | ${ip} | ${error}\n",
		Output: file,
	}
	app.Use(logger.New(config))
	return file
}
func ServeNotFoundHTML(ctx *fiber.Ctx) error {
	path := filepath.Join(".", "static", "not-found.html")
	ctx.Status(fiber.StatusNotFound)
	return ctx.SendFile(path)
}
func ConstructResponse(status int, message string, responseData entity.Entity) fiber.Map {
	response := fiber.Map{
		"status": status,
	}
	if message != "" {
		response["message"] = message
	}
	if responseData != nil {
		response["data"] = responseData
	}
	return response
}

package utils

import (
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

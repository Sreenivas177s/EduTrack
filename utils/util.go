package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func GetLogger() fiber.Handler {
	config := logger.Config{
		Format: "${time} - ${method} | ${path} | ${status} | ${latency} | ${ip} | ${error}\n",
	}
	return logger.New(config)
}

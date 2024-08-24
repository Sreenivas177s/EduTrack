package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func GetLogger() fiber.Handler {
	path := fmt.Sprintf("logs/access-log-%s.txt", time.Now().Format(time.DateOnly))
	fmt.Println(path)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	config := logger.Config{
		Format: "${time} - ${method} | ${path} | ${status} | ${latency} | ${ip} | ${error}\n",
		Output: file,
	}
	return logger.New(config)
}

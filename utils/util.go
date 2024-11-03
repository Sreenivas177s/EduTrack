package utils

import (
	"chat-server/api/entity"
	"encoding/json"
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
func ConstructResponse(status int, message string, entity string, responseData entity.ApiEntity) fiber.Map {
	response := fiber.Map{
		"status": status,
	}
	if message != "" {
		response["message"] = message
	}
	if responseData != nil {
		response["data"] = responseData
	}
	if entity != "" {
		response["entity"] = entity
	}
	return response
}
func WriteFileAtomic(fileName string, data []byte) error {
	tmpName := fmt.Sprintf("%s.tmp", fileName)
	fp, err := os.CreateTemp(".", tmpName)
	// fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0664)
	if err != nil {
		return err
	}
	tmpName = fp.Name()
	_, err = fp.Write(data)
	if err != nil {
		// os.Remove(tmpName)
		return err
	}
	err = fp.Sync()
	if err != nil {
		return err
	}
	fp.Close()
	defer func() {
		if err != nil {
			fp.Close()
			os.Remove(tmpName)
		}
	}()
	return os.Rename(tmpName, fileName)
}

func GenerateConfiguredRoutesJSON(app *fiber.App) {
	data, _ := json.MarshalIndent(app.GetRoutes(true), "", "  ")
	err := WriteFileAtomic("generated-routes.json", data)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info("Routes file generated ", time.Now().Format(time.DateTime), "...")
}

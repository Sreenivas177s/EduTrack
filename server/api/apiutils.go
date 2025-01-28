package api

import (
	"chat-server/api/entity"

	"github.com/gofiber/fiber/v2"
)

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

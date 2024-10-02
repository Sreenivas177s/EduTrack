package utils

import "github.com/gofiber/fiber/v2"

const (
	EntityEventData string = "entity_event_data"
	EntityResponse  string = "entity_response"
)

var NOT_FOUND_JSON = fiber.Map{
	"status":  fiber.StatusNotFound,
	"message": "Provided URL not found",
}

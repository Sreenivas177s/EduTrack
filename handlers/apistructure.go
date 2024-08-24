package ApiHandler

import "github.com/gofiber/fiber/v2"

var NOT_FOUND_JSON = fiber.Map{
	"status":  fiber.StatusNotFound,
	"message": "Provided URL not found",
}

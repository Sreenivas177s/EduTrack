package chats

import (
	"github.com/gofiber/fiber/v2"
)

func GetMiddlewares(httpMethod string) ([]fiber.Handler, error) {
	chain := make([]fiber.Handler, 0, 10)
	chain[0] = func(c *fiber.Ctx) error { return c.SendString("reached middlewares") }
	return chain, nil
}

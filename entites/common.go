package entites

import "github.com/gofiber/fiber/v2"

type Common struct {
	name string
}

func (common *Common) getMiddlewares(httpMethod string) []fiber.Handler {
	chain := make([]fiber.Handler, 0, 10)

	return chain
}

package entityhandlers

import "github.com/gofiber/fiber/v2"

type MiddlewareChainFactory interface {
	GetMiddlewares(httpMethod string) ([]fiber.Handler, error)
}

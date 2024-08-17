package entites

import "github.com/gofiber/fiber/v2"

type Entity interface {
	Name() string
	IsApiEnabled() bool
}

type MiddlewareChainFactory interface {
	getMiddlewares(httpMethod string) []fiber.Handler
}

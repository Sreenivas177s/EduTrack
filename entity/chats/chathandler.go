package chats

import (
	"github.com/gofiber/fiber/v2"
)

func Accumulator(ctx *fiber.Ctx) error {
	ctx.SendString("reached")
	return ctx.Next()
}

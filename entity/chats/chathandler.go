package chats

import (
	"github.com/gofiber/fiber/v2"
)

func Validator(ctx *fiber.Ctx) error {
	inputData := new(Chat)
	if err := ctx.BodyParser(inputData); err != nil {
		return err
	}

	return ctx.Next()
}

func Accumulator(ctx *fiber.Ctx) error {

	return ctx.Next()
}

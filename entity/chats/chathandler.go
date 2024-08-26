package chats

import (
	"chat-server/utils"

	"github.com/gofiber/fiber/v2"
)

func Validator(ctx *fiber.Ctx) error {
	inputData := new(Chat)
	if err := ctx.BodyParser(inputData); err != nil {
		return err
	}
	ctx.Locals(utils.EntityResponse, inputData)
	return ctx.Next()
}

func Preprocessor(ctx *fiber.Ctx) error {

	return ctx.Next()
}

func Accumulator(ctx *fiber.Ctx) error {

	return ctx.Next()
}

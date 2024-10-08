package apiframework

import (
	"chat-server/utils"

	"github.com/gofiber/fiber/v2"
)

func EnforceHeaders(ctx *fiber.Ctx) error {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	return ctx.Next()
}

func UrlNotFound(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusNotFound)
	return ctx.JSON(utils.NOT_FOUND_JSON)
}

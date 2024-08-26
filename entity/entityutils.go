package entity

import (
	"chat-server/utils"

	"github.com/gofiber/fiber/v2"
)

// POST methods implementation -----------------------------------------------------------------------
func Validator(ctx *fiber.Ctx) error {
	event := getEntityEvent(ctx)
	entityAcc := GetEntityHandler(event.entityName, utils.VALIDATOR)
	return entityAcc(ctx)
}
func Accumulator(ctx *fiber.Ctx) error {
	event := getEntityEvent(ctx)
	entityAcc := GetEntityHandler(event.entityName, utils.ACCUMULATOR)
	return entityAcc(ctx)
}
func Postprocessor(ctx *fiber.Ctx) error {
	response := ctx.Locals(utils.EntityResponse)
	if response != nil {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		ctx.Status(fiber.StatusOK)
		return ctx.JSON(response)
	}
	return ctx.Next()
}

// ------------------------------------------------------------------------------------------------------
func EnforceHeaders(ctx *fiber.Ctx) error {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	return ctx.Next()
}

var EmptyHandler = func(ctx *fiber.Ctx) error { return ctx.Next() }

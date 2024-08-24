package entity

import (
	"chat-server/utils"

	"github.com/gofiber/fiber/v2"
)

// POST methods implementation

func Accumulator(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	event := getEntityEvent(ctx)
	entityAcc := GetEntityHandler(event.entityName, utils.ACCUMULATOR)
	return entityAcc(ctx)
}

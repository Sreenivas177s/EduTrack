package apiframework

import (
	"chat-server/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EntityEvent struct {
	method     string
	entityid   int64
	entityName string
	apiversion string
}

func ParseEntityEvent(ctx *fiber.Ctx) error {
	var (
		entityid   int64
		err        error
		tempId     = ctx.Params("entityid")
		entityName = ctx.Params("entity")
		event      *EntityEvent
	)

	if tempId != "" {
		entityid, err = strconv.ParseInt(tempId, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}
	}

	if entityName == "" {
		return fiber.NewError(fiber.StatusNotFound)
	}
	event = &EntityEvent{
		method:     ctx.Method(),
		entityName: entityName,
		entityid:   entityid,
		apiversion: ctx.Params("version"),
	}
	ctx.Locals(utils.EntityEventData, event)
	return ctx.Next()
}

func getEntityEvent(ctx *fiber.Ctx) *EntityEvent {
	data := ctx.Locals(utils.EntityEventData)
	return data.(*EntityEvent)
}

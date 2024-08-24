package entity

import (
	"chat-server/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type EntityEvent struct {
	method     string
	entityid   int64
	entityName string
	apiversion string
}

func ParseEntityData(ctx *fiber.Ctx) error {
	var (
		entityid int64
		err      error
	)
	tempId := ctx.Params("entityid")
	if tempId != "" {
		entityid, err = strconv.ParseInt(tempId, 10, 64)
		if err != nil {
			return err
		}
	}
	entityName := ctx.Params("entity")
	event := &EntityEvent{
		method:     ctx.Method(),
		entityName: entityName,
		entityid:   entityid,
		apiversion: ctx.Params("version"),
	}
	ctx.Locals(utils.EntityEventData, event)
	log.Debug(event)
	return nil
}

func getEntityEvent(ctx *fiber.Ctx) EntityEvent {
	return ctx.Locals(utils.EntityEventData).(EntityEvent)
}

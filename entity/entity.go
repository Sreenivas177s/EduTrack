package entity

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

	event := &EntityEvent{
		method:     ctx.Method(),
		entityName: ctx.Params("entity"),
		entityid:   entityid,
		apiversion: ctx.Params("version"),
	}
	ctx.Locals(utils.EntityEventData, event)
	return nil
}

package entity

import (
	"chat-server/entity/chats"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var handlerMapping map[string]fiber.Handler = map[string]fiber.Handler{
	"chats-accumulator": chats.Accumulator,
}

func GetEntityHandler(entityName string, handlerType string) fiber.Handler {
	key := fmt.Sprintf("%s-%s", entityName, handlerType)
	mcf := handlerMapping[key]
	log.Debug(mcf)
	if mcf == nil {
		return EmptyHandler
	}
	return mcf
}
func EmptyHandler(ctx *fiber.Ctx) error {
	return ctx.Next()
}

package entity

import (
	"chat-server/entity/chats"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var EmptyHandler = func(ctx *fiber.Ctx) error { return ctx.Next() }

var handlerMapping map[string]fiber.Handler = map[string]fiber.Handler{
	// "chats-validator": chats.Validator,
	"chats-accumulator": chats.Accumulator,
	// "chats-postprocessor": chats.Postprocessor,
}

func GetEntityHandler(entityName string, handlerType string) fiber.Handler {
	key := fmt.Sprintf("%s-%s", entityName, handlerType)
	mcf := handlerMapping[key]
	if mcf == nil {
		return EmptyHandler
	}
	return mcf
}

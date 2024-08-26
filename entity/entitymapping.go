package entity

import (
	"chat-server/entity/chats"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var handlerMapping map[string]fiber.Handler = map[string]fiber.Handler{
	// methods related to chat api
	"chats-validator":    chats.Validator,
	"chats-accumulator":  chats.Accumulator,
	"chats-preprocessor": chats.Preprocessor,
}

// can allow 1 custom middle ware
func GetEntityHandler(entityName string, handlerType string) fiber.Handler {
	key := fmt.Sprintf("%s-%s", entityName, handlerType)
	mcf := handlerMapping[key]
	if mcf == nil {
		return EmptyHandler
	}
	return mcf
}

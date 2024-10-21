package apiframework

import (
	"chat-server/utils"
	"reflect"

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

func ExecuteEntityMethod(value reflect.Value, methodName string, params []reflect.Value) []reflect.Value {
	preProcessor := value.MethodByName(methodName)
	if preProcessor.IsValid() {
		return preProcessor.Call(params)
	}
	return nil
}

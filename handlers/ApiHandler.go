package ApiHandler

import (
	"chat-server/entites"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HandleApiCall(ctx *fiber.Ctx) error {

	httpMethod := ctx.Method()

	apiApp := ctx.App()
	apiApp.Group("/:entity")
	entity := ctx.Params("entity")
	mwcbuilder := entites.GetEntityMapping(entity)

	switch httpMethod {
	case fiber.MethodGet:
		mwc, _ := mwcbuilder.getEntityMiddlewares(httpMethod)
		apiApp.Get("/:entityid?", mwc)
	default:
		panic("savu da")
	}

	return ctx.SendString(fmt.Sprintf("e = %s ei = %s", ctx.Params("entity"), ctx.Params("entityid")))
}
func getEntityHandler(entity string) string {
	return "asdasd"
}

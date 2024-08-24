package ApiHandler

import (
	"chat-server/entity"
	"chat-server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func HandleApiCall(app fiber.Router) error {
	// route will start as '/api'

	// parse entity data
	app.Use(entity.ParseEntityData)

	//handle necessary middle wares
	// POST HANDLER
	app.Post(`/:entity`, entity.Accumulator)

	// GET HANDLER
	app.Get(`/:entity/:entityid?`, func(ctx *fiber.Ctx) error {
		log.Debug(ctx.Locals(utils.EntityEventData))
		return ctx.Next()
	})

	// // PUT HANDLER
	// app.Put(`/:entity/:entityid`)

	// // DELETE HANDLER
	// app.Delete(`/:entity/:entityid`)

	// httpMethod := ctx.Method()

	// entityName := ctx.Params("entity")
	// // version := ctx.Params("version")
	// mwcbuilder := entity.GetEntityMapping(entityName)
	// log.Debug(ctx.AllParams())

	// switch httpMethod {
	// case fiber.MethodGet:
	// 	mwc, _ := mwcbuilder.GetMiddlewares(httpMethod)
	// default:
	// 	panic("savu da")
	// }

	return nil
}

func UrlNotFound(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusNotFound)
	return ctx.JSON(NOT_FOUND_JSON)
}

func handlePOST() []fiber.Handler {
	return []fiber.Handler{
		// validator(),
		// preprocessor(),
		entity.Accumulator,
		// postprocessor(),
	}
}

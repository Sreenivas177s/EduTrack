package Handler

import (
	"chat-server/entity"
	"chat-server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func HandleApiCall(app fiber.Router) {
	// route will start as '/api'

	//handle entity independant middle wares here

	// POST HANDLER
	app.Post(`/:entity`, handlePOST()...)

	// GET HANDLER
	app.Get(`/:entity/:entityid?`, func(ctx *fiber.Ctx) error {
		log.Debug(ctx.Locals(utils.EntityEventData))
		return ctx.Next()
	})

	// // PUT HANDLER
	// app.Put(`/:entity/:entityid`)

	// // DELETE HANDLER
	// app.Delete(`/:entity/:entityid`)

	app.All("/*", UrlNotFound)
}

func UrlNotFound(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusNotFound)
	return ctx.JSON(NOT_FOUND_JSON)
}

func handlePOST() []fiber.Handler {
	return []fiber.Handler{
		entity.ParseEntityData,
		entity.Validator,
		// entity.preprocessor(),
		// entity.Accumulator,
		entity.Postprocessor,
	}
}

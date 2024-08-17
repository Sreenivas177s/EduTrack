package ApiHandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HandleApiCall(ctx *fiber.Ctx) error {
	return ctx.SendString(fmt.Sprintf("e = %s ei = %s", ctx.Params("entity"), ctx.Params("entityid")))
}

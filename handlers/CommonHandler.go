package Handler

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func ServeNotFoundHTML(ctx *fiber.Ctx) error {
	path := filepath.Join(".", "static", "not-found.html")
	ctx.Status(fiber.StatusNotFound)
	return ctx.SendFile(path)
}

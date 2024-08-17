package entites

import "github.com/gofiber/fiber/v2"

type Chat struct {
	id        uint32
	name      string
	owner     string
	createdAt string //unix-epoch
}

func (chat *Chat) Name() string {
	return chat.name
}
func (chat *Chat) IsApiEnabled() bool {
	return true
}

type ChatMWCFactory struct{}

func (*ChatMWCFactory) getMiddlewares(httpMethod string) []fiber.Handler {
	chain := make([]fiber.Handler, 0, 10)

	return chain
}

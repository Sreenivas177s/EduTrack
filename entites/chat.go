package entites

// import "github.com/gofiber/fiber/v2"

type Chat struct {
	id        uint32
	name      string
	owner     string
	createdAt string //unix-epoch
}

// func (chat *Chat) HandleApi(ctx *fiber.Ctx) string {
// }

package auth

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func HandleAuth(app fiber.Router) {

}

func initJWTMiddleware(app *fiber.App) {
	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
		},
	}
}

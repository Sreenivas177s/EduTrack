package auth

import (
	"chat-server/api/entity"
	"chat-server/database"
	"fmt"
	"os"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

const TOKEN_USER = "token_user"
const LOGGEDIN_USER = "loggedin_user"

func HandleAuth(app fiber.Router) {
	// auth related apis

	app.Post("/login", authorizeLogin)

	// app.Post("/signup", signUpUser)

	// app.Post("/logout")
}

func signUpUser(ctx *fiber.Ctx) error {

	return ctx.Next()
}

func authorizeLogin(ctx *fiber.Ctx) error {
	userIdentifier := ctx.FormValue("userEmail")
	userPass := ctx.FormValue("userPassword")

	user, err := AuthorizeUser(userIdentifier, userPass)
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	// if authorized generate JWT creds and send as asresponse
	jwtClaims := jwt.MapClaims{
		"user_id": user.ID(),
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	//sign token and send with cookie
	signedToken, err := token.SignedString(GetJWTSigningKey())
	if err != nil || signedToken == "" {
		ctx.SendStatus(fiber.StatusInternalServerError)
	}
	ctx.Status(fiber.StatusAccepted)
	if os.Getenv("USE_COOKIE_AUTH") == "true" {
		cookieData := &fiber.Cookie{
			Name:  fiber.HeaderAuthorization,
			Value: signedToken,
		}
		ctx.Cookie(cookieData)
		return ctx.SendStatus(fiber.StatusAccepted)
	}
	return ctx.JSON(fiber.Map{"token": signedToken})
}

func InitAuthMiddleWare(app fiber.Router) fiber.Router {
	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    GetJWTSigningKey(),
		},
		SuccessHandler: authSuccessHandler,
		ContextKey:     TOKEN_USER,
	}

	if os.Getenv("USE_COOKIE_AUTH") == "true" {
		config.TokenLookup = fmt.Sprintf("%s:%s", "cookie", fiber.HeaderAuthorization)
	}
	app.Use(jwtware.New(config))

	return app
}

func authSuccessHandler(ctx *fiber.Ctx) error {
	jwtUser := ctx.Locals(TOKEN_USER).(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	// parse user details and fetch user
	userID := uint(claims["user_id"].(float64))
	log.Debugf("user id = %s", userID)
	// fetch user and set context
	user := new(entity.User)
	result := database.GetDBRef().First(&user, userID)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusUnauthorized)
	}
	ctx.Locals(LOGGEDIN_USER, user)
	return ctx.Next()
}

package auth

import (
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
const typeAuthorization = "Authorization"

func HandleAuth(app fiber.Router) {
	// auth related apis

	app.Post("/login", authorizeLogin)

	app.Post("/signup", signUpUser)

	// app.Post("/logout")
}

func signUpUser(ctx *fiber.Ctx) error {

	return ctx.Next()
}

func authorizeLogin(ctx *fiber.Ctx) error {
	userIdentifier := ctx.FormValue("userIdentifier")
	userPass := ctx.FormValue("userPassword")

	user, err := AuthorizeUser(userIdentifier, userPass)
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	// if authorized generate JWT creds and send as asresponse
	jwtClaims := jwt.MapClaims{
		"name":  user.FirstName,
		"email": user.EmailId,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	//sign token and send with cookie
	signedToken, err := token.SignedString(GetJWTSigningKey())
	if err != nil || signedToken == "" {
		ctx.SendStatus(fiber.StatusInternalServerError)
	}
	ctx.Status(fiber.StatusAccepted)
	ctx.Redirect("//index.html")
	if os.Getenv("USE_COOKIE_AUTH") == "true" {
		cookieData := &fiber.Cookie{
			Name:  typeAuthorization,
			Value: signedToken,
		}
		ctx.Cookie(cookieData)
		return ctx.SendStatus(fiber.StatusAccepted)
	}
	return ctx.JSON(fiber.Map{"token": signedToken})
}

func InitAuthMiddleWare(app fiber.Router) fiber.Router {
	authType := "header"
	if os.Getenv("USE_COOKIE_AUTH") == "true" {
		authType = "cookie"
	}
	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    GetJWTSigningKey(),
		},
		SuccessHandler: authSuccessHandler,
		ContextKey:     TOKEN_USER,
		TokenLookup:    fmt.Sprintf("%s:%s", authType, typeAuthorization),
	}
	app.Use(jwtware.New(config))

	return app
}

func authSuccessHandler(ctx *fiber.Ctx) error {
	user := ctx.Locals(TOKEN_USER).(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// parse user details and fetch user
	emailid := claims["email"].(string)
	name := claims["name"].(string)
	log.Debugf("user : %s email : %s", name, emailid)
	//fetch user and set context
	loggedInUser := database.GetUser(emailid)
	ctx.Locals(LOGGEDIN_USER, loggedInUser)
	return ctx.Next()
}

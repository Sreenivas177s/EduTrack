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

func HandleAuth(app *fiber.App) {
	authRouter := app.Group("/auth")
	// auth related apis
	authRouter.Post("/login", authorizeLogin)

	authRouter.Post("/signup", signUpUser)

	authRouter.Put("/logout", handleLogout)
}

func signUpUser(ctx *fiber.Ctx) error {

	return ctx.Next()
}
func handleLogout(ctx *fiber.Ctx) error {
	jwtUser := ctx.Locals(TOKEN_USER).(*jwt.Token)
	database.SetBlacklistToken(jwtUser.Raw, fmt.Sprintf("%v", jwtUser.Claims.(jwt.MapClaims)["sub"]))
	ctx.ClearCookie(fiber.HeaderAuthorization)
	return ctx.SendStatus(fiber.StatusOK)
}
func authorizeLogin(ctx *fiber.Ctx) error {

	form := new(LoginForm)
	if err := ctx.BodyParser(form); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userIdentifier := form.Email_id
	userPass := form.Password

	user, err := AuthorizeUser(userIdentifier, userPass)
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	// if authorized generate JWT creds and send as asresponse
	currentTime := time.Now().Local()
	jwtClaims := jwt.MapClaims{
		"exp": currentTime.Add(time.Hour * 1).Unix(),
		"iat": currentTime.Unix(),
		"sub": user.ID(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	//sign token and send with cookie
	signedToken, err := token.SignedString(GetJWTSigningKey())
	if err != nil {
		log.Errorf("Error signing token: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	if os.Getenv("USE_COOKIE_AUTH") == "true" {
		cookieData := &fiber.Cookie{
			Name:     fiber.HeaderAuthorization,
			Value:    signedToken,
			Secure:   os.Getenv("PROD_ENV") == "true",
			HTTPOnly: true,
			SameSite: fiber.CookieSameSiteStrictMode,
			Expires:  currentTime.Add(time.Hour),
		}
		ctx.Cookie(cookieData)
		return ctx.SendStatus(fiber.StatusOK)
	}
	ctx.Status(fiber.StatusOK)
	return ctx.JSON(fiber.Map{"token": signedToken})
}

func GetAuthMiddleWare() fiber.Handler {
	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    GetJWTSigningKey(),
		},
		SuccessHandler: authSuccessHandler,
		ErrorHandler:   authErrorHandler,
		ContextKey:     TOKEN_USER,
		Filter:         authFilter,
	}

	if os.Getenv("USE_COOKIE_AUTH") == "true" {
		config.TokenLookup = fmt.Sprintf("%s:%s", "cookie", fiber.HeaderAuthorization)
	}
	return jwtware.New(config)
}

func authSuccessHandler(ctx *fiber.Ctx) error {
	jwtUser := ctx.Locals(TOKEN_USER).(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	// check if token is blacklisted
	isBlacklisted, err := database.IsBlacklistToken(jwtUser.Raw)
	if err != nil || isBlacklisted {
		ctx.ClearCookie(fiber.HeaderAuthorization)
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	// parse user details and fetch user
	userID := uint(claims["sub"].(float64))
	// fetch user and set context
	user, err := database.GetUserByID(userID)
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	ctx.Locals(LOGGEDIN_USER, user)
	return ctx.Next()
}

func authErrorHandler(ctx *fiber.Ctx, err error) error {
	log.Errorf("Error in auth middleware: %v", err)
	return ctx.SendStatus(fiber.StatusUnauthorized)
}

func authFilter(ctx *fiber.Ctx) bool {
	if ctx.Path() == "/auth/login" || ctx.Path() == "/auth/signup" || ctx.Path() == "/auth/logout" || (os.Getenv("DEBUG_MODE") == "true" && ctx.Path() == "/api/v1/users" && ctx.Method() == fiber.MethodPost) {
		return true
	}
	return false
}

package apiframework

import (
	"chat-server/api-framework/entity"
	"chat-server/utils"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func HandleApiCall(app fiber.Router) {
	// route will start as '/api/version'

	//handler needs to know all params before parsing them so calling parse event method for every configured url
	app.Use(EnforceHeaders)
	entityApi := app.Group("/:entity")
	// POST HANDLER
	entityApi.Post(``, apiMiddlewareChain(fiber.MethodPost)...)

	// GET HANDLER
	entityApi.Get(`/:entityid?`, func(ctx *fiber.Ctx) error {
		log.Debug(ctx.Locals(utils.EntityEventData))
		return ctx.Next()
	})

	// // PUT HANDLER
	// app.Put(`/:entity/:entityid`)

	// // DELETE HANDLER
	// app.Delete(`/:entity/:entityid`)

	app.All("/*", UrlNotFound)
}

func apiMiddlewareChain(httpMethod string) []fiber.Handler {
	return []fiber.Handler{
		ParseEntityEvent,
		handlePOST,
	}
}

func handlePOST(ctx *fiber.Ctx) error {
	event := getEntityEvent(ctx)
	if event == nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	// allocate and populate data
	inputType := event.structType
	inputValueAllocatedPointer := reflect.New(inputType)
	inputData := inputValueAllocatedPointer.Interface()
	if err := ctx.BodyParser(inputData); err != nil {
		return err
	}
	methodParams := []reflect.Value{reflect.ValueOf(ctx.Method())}
	// validate input data
	executeMethod(inputValueAllocatedPointer, entity.METHOD_VALIDATOR, methodParams)

	// pre persistence handling
	executeMethod(inputValueAllocatedPointer, entity.METHOD_PRE_PROCESSOR, methodParams)

	// add the provided data into persistence layer

	//post persistence handling
	executeMethod(inputValueAllocatedPointer, entity.METHOD_POST_PROCESSOR, methodParams)

	isSuccess := true
	if isSuccess {
		response := utils.ConstructResponse(fiber.StatusCreated, "", inputData.(entity.Entity))
		ctx.Status(fiber.StatusCreated)
		return ctx.JSON(response)
	}
	return ctx.Next()
}

func executeMethod(value reflect.Value, methodName string, params []reflect.Value) {
	preProcessor := value.MethodByName(methodName)
	if preProcessor.IsValid() {
		preProcessor.Call(params)
	}
}

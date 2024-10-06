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

	//parse entity event
	app = app.Use("/*", EnforceHeaders, ParseEntityEvent)

	// POST HANDLER
	app.Post(`/:entity`, handlePOST)

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

func handlePOST(ctx *fiber.Ctx) error {
	event := getEntityEvent(ctx)
	if event == nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	// allocate and populate data
	ipType := GetDefinedType(event.entityName)
	ipValue := reflect.New(ipType)
	inputData := ipValue.Interface()
	if err := ctx.BodyParser(inputData); err != nil {
		return err
	}
	methodParams := make([]reflect.Value, 1)
	methodParams[0] = reflect.ValueOf(ctx.Method())
	// validate input data
	executeMethod(ipValue, entity.METHOD_VALIDATOR, methodParams)

	// pre persistence handling
	executeMethod(ipValue, entity.METHOD_PRE_PROCESSOR, methodParams)

	// add the provided data into persistence layer

	//post persistence handling
	executeMethod(ipValue, entity.METHOD_POST_PROCESSOR, methodParams)

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

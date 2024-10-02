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
	app = app.Use("api", EnforceHeaders, ParseEntityEvent)

	// POST HANDLER
	app.Post(`/:entity`, ParseEntityEvent, handlePOST)

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
	// allocate and populate data
	ipType := GetDefinedType(event.entityName)
	ipValue := reflect.New(ipType)
	inputData := ipValue.Interface()
	if err := ctx.BodyParser(inputData); err != nil {
		return err
	}
	// validate input data
	executeMethod(ipValue, entity.METHOD_VALIDATOR, nil)

	// pre persistence handling
	executeMethod(ipValue, entity.METHOD_PRE_PROCESSOR, nil)

	// add the provided data into persistence layer

	//post persistence handling
	executeMethod(ipValue, entity.METHOD_POST_PROCESSOR, nil)

	ctx.Locals(utils.EntityResponse, inputData)
	return ctx.Next()
}

func executeMethod(value reflect.Value, methodName string, params []reflect.Value) {
	preProcessor := value.MethodByName(methodName)
	if preProcessor.IsValid() {
		preProcessor.Call(params)
	}
}

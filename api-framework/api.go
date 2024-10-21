package apiframework

import (
	"chat-server/api-framework/entity"
	"chat-server/database"
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
	entityApi.Post(``, ParseEntityEvent, handlePOST).Name("Add Entity")

	// GET HANDLER
	entityApi.Get(`/:entityid<regex(\d{1,19})>?`, func(ctx *fiber.Ctx) error {
		log.Debug(ctx.Locals(utils.EntityEventData))
		return ctx.Next()
	}).Name("Entity Get")

	// // PUT HANDLER
	// app.Put(`/:entity/:entityid`)

	// // DELETE HANDLER
	// app.Delete(`/:entity/:entityid`)

	entityApi.All("/*", UrlNotFound).Name("Not Found - API")
}

func handlePOST(ctx *fiber.Ctx) error {
	event := getEntityEvent(ctx)
	if event == nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	// allocate and populate data
	inputType := event.structType
	inputValueAllocatedPointer := reflect.New(inputType)
	var inputData entity.ApiEntity = inputValueAllocatedPointer.Interface().(entity.ApiEntity)
	if err := ctx.BodyParser(inputData); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	methodParams := []reflect.Value{reflect.ValueOf(ctx.Method())}
	//check for user authorization
	authorizationResult := ExecuteEntityMethod(inputValueAllocatedPointer, entity.METHOD_AUTHORIZER, methodParams)
	if !authorizationResult[0].Bool() || authorizationResult[1].Interface() != nil {
		return fiber.NewError(fiber.StatusUnauthorized)
	}
	// validate input data
	ExecuteEntityMethod(inputValueAllocatedPointer, entity.METHOD_VALIDATOR, methodParams)

	// pre persistence handling
	ExecuteEntityMethod(inputValueAllocatedPointer, entity.METHOD_PRE_PROCESSOR, methodParams)

	// add the provided data into persistence layer
	dbReference := database.GetDBRef()
	txn := dbReference.Begin()
	if err := txn.Create(inputData).Error; err != nil {
		txn.Rollback()
		return fiber.NewError(fiber.StatusBadRequest)
	}
	if err := txn.Commit().Error; err != nil {
		txn.Rollback()
		return fiber.NewError(fiber.StatusBadRequest)
	}
	//refetch data for post processing
	insertedDataPointer := reflect.New(inputType)
	refetchedData := insertedDataPointer.Interface().(entity.ApiEntity)
	insertedID := inputData.ID()
	dbReference.Take(&refetchedData, insertedID)

	//post persistence handling
	ExecuteEntityMethod(inputValueAllocatedPointer, entity.METHOD_POST_PROCESSOR, methodParams)

	response := utils.ConstructResponse(fiber.StatusCreated, "Created successfully", event.entityName, refetchedData)
	ctx.Status(fiber.StatusCreated)
	return ctx.JSON(response)
}

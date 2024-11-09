package api

import (
	"chat-server/api/entity"
	"chat-server/database"
	"chat-server/utils"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func HandleApiCall(app *fiber.App) {
	apiRouter := app.Group("/api")
	// route will start as '/api/version'

	//handler needs to know all params before parsing them so calling parse event method for every configured url
	apiRouter.Use(EnforceHeaders)

	entityApi := apiRouter.Group("/:entity")
	// POST HANDLER
	entityApi.Post(``, ParseEntityEvent, handlePOST).Name("Entity Add")

	// GET HANDLER
	entityApi.Get(`/:entityid<regex(\d{1,19})>?`, ParseEntityEvent, handleGET).Name("Entity Get")

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
	if err := Authorize(inputValueAllocatedPointer, methodParams); err != nil {
		return err
	}
	if err := Validate(inputValueAllocatedPointer, inputType, methodParams); err != nil {
		return err
	}

	// add the provided data into persistence layer
	dbReference := database.GetDBRef()
	txn := dbReference.Begin()
	// pre persistence handling
	ExecuteEntityMethod(inputValueAllocatedPointer, utils.METHOD_PRE_PROCESSOR, methodParams)
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
	ExecuteEntityMethod(inputValueAllocatedPointer, utils.METHOD_POST_PROCESSOR, methodParams)

	response := utils.ConstructResponse(fiber.StatusCreated, "Created successfully", event.entityName, refetchedData)
	ctx.Status(fiber.StatusCreated)
	return ctx.JSON(response)
}

func handleGET(ctx *fiber.Ctx) error {
	event := getEntityEvent(ctx)
	if event == nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	entityType := event.structType
	fetchedData := reflect.New(entityType)
	castedValue := fetchedData.Interface()
	entityID := event.entityid

	db := database.GetDBRef()
	result := db.First(&castedValue, entityID)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	methodParams := []reflect.Value{reflect.ValueOf(ctx.Method())}
	if err := Authorize(fetchedData, methodParams); err != nil {
		return err
	}
	return ctx.JSON(castedValue)
}

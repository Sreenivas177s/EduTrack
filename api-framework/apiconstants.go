package apiframework

import (
	"chat-server/api-framework/validator"
	"chat-server/utils"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func EnforceHeaders(ctx *fiber.Ctx) error {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	return ctx.Next()
}

func UrlNotFound(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusNotFound)
	return ctx.JSON(utils.NOT_FOUND_JSON)
}

func ExecuteEntityMethod(value reflect.Value, methodName string, params []reflect.Value) []reflect.Value {
	preProcessor := value.MethodByName(methodName)
	if preProcessor.IsValid() {
		return preProcessor.Call(params)
	}
	return nil
}

func Authorize(inputData reflect.Value, methodParams []reflect.Value) *fiber.Error {
	authorizationResult := ExecuteEntityMethod(inputData, utils.METHOD_AUTHORIZER, methodParams)
	if !authorizationResult[0].Bool() || authorizationResult[1].Interface() != nil {
		if authorizationResult[1].Interface() != nil {
			return fiber.NewError(fiber.StatusUnauthorized, authorizationResult[1].Interface().(error).Error())
		}
		return fiber.NewError(fiber.StatusUnauthorized)
	}
	return nil
}

func Validate(inputData reflect.Value, inputType reflect.Type, methodParams []reflect.Value) *fiber.Error {
	//standard field validation and preProcess
	if validationErr := validator.BasicFieldValidation(inputData, inputType); validationErr != nil {
		return fiber.NewError(fiber.StatusBadRequest, validationErr.Error())
	}
	// validate input data
	if customValidErr := ExecuteEntityMethod(inputData, utils.METHOD_VALIDATOR, methodParams)[0].Interface(); customValidErr != nil {
		return fiber.NewError(fiber.StatusBadRequest, customValidErr.(error).Error())
	}
	return nil
}

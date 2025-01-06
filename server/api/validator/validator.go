package validator

import (
	"chat-server/api/entity"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func BasicFieldValidation(data reflect.Value, dataType reflect.Type) error {
	resolvedData := data
	if resolvedData.Kind() == reflect.Ptr {
		resolvedData = resolvedData.Elem()
	}
	for fieldIdx := 0; fieldIdx < dataType.NumField(); fieldIdx++ {
		typeField := dataType.Field(fieldIdx)
		// skip non exported fields
		if !typeField.IsExported() {
			continue
		}
		valField := resolvedData.Field(fieldIdx)
		if valField.Kind() == reflect.Struct && (isParsableType(typeField.Type)) {
			err := BasicFieldValidation(valField, typeField.Type)
			if err != nil {
				return err
			}
		}
		if validateKey, keyPresent := typeField.Tag.Lookup("validate"); keyPresent {
			for _, key := range strings.Split(validateKey, ",") {
				isValid := validate(key, valField.Interface())
				if !isValid {
					message := fmt.Sprintf("%s validation failed for - %s", strings.ToTitle(validateKey), typeField.Tag.Get("json"))
					return errors.New(message)
				}
			}
		}
	}
	return nil
}

var structsToParse []reflect.Type = []reflect.Type{reflect.TypeOf((*entity.ApiEntity)(nil)).Elem()}

func isParsableType(dtype reflect.Type) bool {
	if dtype.Name() == "ApiBase" {
		return true
	}
	for _, typeVal := range structsToParse {
		if dtype.Implements(typeVal) {
			return true
		}
	}
	return false
}

func validate(valFn string, data any) bool {
	switch valFn {
	case "email":
		return EvaluateRegex(TYPE_EMAIL, data.(string))
	}
	// no validations are present so returning true
	return true
}

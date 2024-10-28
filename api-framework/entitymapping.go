package apiframework

import (
	"chat-server/api-framework/entity"
	"reflect"
)

// api url to entity struct mapping
var entityMapping map[string]reflect.Type = map[string]reflect.Type{
	"users":        reflect.TypeOf(entity.User{}),
	"institutions": reflect.TypeOf(entity.Institution{}),
	"campuses":     reflect.TypeOf(entity.Campus{}),
}

func GetDefinedType(dtype string) reflect.Type {
	if dtype == "" {
		return nil
	}
	return entityMapping[dtype]
}

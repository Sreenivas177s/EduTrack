package generators

import (
	"chat-server/api/entity"
	"chat-server/utils"
	"fmt"
	"reflect"
	"strings"
)

const FILENAME = "./entity-types.ts"

var entitiesToGenerate = []entity.ApiEntity{
	&entity.User{},
}

// generator method for generating typescript types
func main() {
	fileData := generateTsTypes()
	utils.WriteFileAtomic(FILENAME, []byte(fileData))
}

func generateTsTypes() string {
	tsFileContent := "// This file is auto-generated. Do not edit this file manually.\n// Author : sreenivas \n\n"
	for _, entity := range entitiesToGenerate {
		rsType := reflect.TypeOf(entity)
		typeContent := ""
		for idx := 0; idx < rsType.NumField(); idx++ {
			typeField := rsType.Field(idx)
			// skip non exported fields
			if !typeField.IsExported() {
				continue
			}
			if apikey, keyPresent := typeField.Tag.Lookup("json"); keyPresent && apikey != "-" && !strings.Contains(apikey, "omitempty") {
				typeContent += getMemberType(apikey, typeField.Type.Name(), false)
			}
		}
		if typeContent != "" {
			tsFileContent += fmt.Sprintf("%s {\n%s}\n\n", getTypeHeader(rsType.Name()), typeContent)
		}
	}
	return tsFileContent
}

const (
	CHAR_OPEN_BRACE          = "{"
	CHAR_CLOSE_BRACE         = "}"
	CHAR_COLON               = ":"
	CHAR_QUESTION_MARK       = "?"
	CHAR_SEMI_COLON          = ";"
	CHAR_OPEN_ANGLE_BRACKET  = "<"
	CHAR_CLOSE_ANGLE_BRACKET = ">"
)

func getTypeHeader(typeName string) string {
	return "export type " + typeName + " = "
}
func getMemberType(memberName string, memberType string, optional bool) string {
	optionalStr := ""
	if optional {
		optionalStr = CHAR_QUESTION_MARK
	}
	return fmt.Sprintf("\t%s%s%s %s\n", memberName, optionalStr, CHAR_COLON, memberType)
}

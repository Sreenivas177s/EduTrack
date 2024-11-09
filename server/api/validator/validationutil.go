package validator

import (
	"regexp"

	"github.com/gofiber/fiber/v2/log"
)

const (
	TYPE_EMAIL    string = "email"
	TYPE_PASSWORD string = "password"
)

var regexMap map[string]string = map[string]string{
	TYPE_EMAIL:    `\b[\w\.-]+@[\w\.-]+\.\w{2,4}\b`,
	TYPE_PASSWORD: `([a-z0-9_\-]{1,5}:\/\/)?(([a-z0-9_\-]{1,}):([a-z0-9_\-]{1,})\@)?((www\.)|([a-z0-9_\-]{1,}\.)+)?([a-z0-9_\-]{3,})(\.[a-z]{2,4})(\/([a-z0-9_\-]{1,}\/)+)?([a-z0-9_\-]{1,})?(\.[a-z]{2,})?(\?)?(((\&)?[a-z0-9_\-]{1,}(\=[a-z0-9_\-]{1,})?)+)?`,
}

func EvaluateRegex(vtype string, value string) bool {
	isValid, err := regexp.MatchString(regexMap[vtype], value)
	if err != nil {
		log.Warn(err.Error())
	}
	return isValid
}

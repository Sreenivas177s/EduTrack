package utils

import "regexp"

const emailRegex = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`

type EmailID string

func (email EmailID) Validate() bool {
	result, err := regexp.MatchString(emailRegex, string(email))
	if err != nil || !result {
		return false
	}
	return true
}

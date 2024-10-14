package entity

import (
	"github.com/gofiber/fiber/v2/log"
)

type User struct {
	ApiBase
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	EmailId   string        `json:"email_id"`
	Status    CurrentStatus `json:"user_status"`
}

// api handler methods
func (user *User) Validate(httpMethod string) {
	// httpMethod := params[0]

}

// function option methods
func (user *User) FillDefaults() error {
	log.Debug("reached method")
	// user.LastUpdatedTime = time.Now()
	return nil
}

func (user *User) RemoveInternalFields() error {
	log.Debug("reached method")
	return nil
}

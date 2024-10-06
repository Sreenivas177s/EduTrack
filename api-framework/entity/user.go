package entity

import "reflect"

type UserStatus int

const (
	ACTIVE UserStatus = iota
	INACTIVE
	RESIGNED
)

type User struct {
	ApiBase
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	EmailId      string     `json:"email_id"`
	Organization any        `json:"organization"`
	Status       UserStatus `json:"user_status"`
}

func (user *User) Validate(params []reflect.Value) {
	// httpMethod := params[0]

}

package entity

import (
	"gorm.io/gorm"
)

type UserStatus int

const (
	ACTIVE UserStatus = iota
	INACTIVE
	RESIGNED
)

type User struct {
	gorm.Model
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	EmailId      string     `json:"email_id"`
	Organization any        `json:"organization"`
	Status       UserStatus `json:"user_status"`
}

func (user *User) Validate(httpMethod string) {
	// httpMethod := params[0]

}
func (c *User) New() Entity {
	return Chat{}
}

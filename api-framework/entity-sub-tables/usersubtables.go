package entitysubtables

import (
	"gorm.io/gorm"
)

// user auth storage table
type UserAuthDetails struct {
	gorm.Model
	Password string `json:"password"`
}

// user institution mapping

//user campus mapping

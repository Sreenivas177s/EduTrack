package entity

import (
	"gorm.io/gorm"
)

type ApiBase struct {
	gorm.Model
	ImageURL string `json:"image_url"`
}

type Entity interface {
	FillDefaults() error
	RemoveInternalFields() error
}

type ApiEntity interface {
	Authorize(httpMethod string) (bool, error)
	Validate(httpMethod string) error
	ID() uint
	Preprocessor(httpMethod string) error
	HandleOperation(operation string) error
}

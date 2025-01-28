package entity

import (
	"gorm.io/gorm"
)

type ApiBase struct {
	gorm.Model
	ImageURL string `json:"image_url"`
}

type Entity interface {
	ID() uint
	FillDefaults() error
	RemoveInternalFields() error
}

type ApiEntity interface {
	Authorize(httpMethod string) (bool, error)
	Validate(httpMethod string) error
	Preprocessor(httpMethod string) error
	HandleOperation(operation string) error
}

package entity

import (
	"gorm.io/gorm"
)

type ApiBase struct {
	// Id              int       `json:"id"`
	// CreatedTime     time.Time `json:"created_time"`
	// LastUpdatedTime time.Time `json:"last_updated_time"`
	gorm.Model
	ImageURL string `json:"image_url"`
}

type Entity interface {
	FillDefaults() error
	RemoveInternalFields() error
}

type ApiEntity interface {
	Authorize(httpMethod string) (bool, error)
	Validate(httpNethod string) error
	ID() uint
}

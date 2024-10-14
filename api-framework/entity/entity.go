package entity

import "time"

type ApiBase struct {
	Id              int       `json:"id"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
	ImageURL        string    `json:"image_url"`
}

type Entity interface {
	FillDefaults() error
	RemoveInternalFields() error
}

package entity

const (
	METHOD_VALIDATOR      = "Validate"
	METHOD_PRE_PROCESSOR  = "Preprocessor"
	METHOD_POST_PROCESSOR = "Postprocessor"
)

type ApiBase struct {
	Id              int   `json:"id"`
	CreatedTime     int64 `json:"created_time"`
	LastUpdatedTime int64 `json:"last_updated_time"`
}

type Entity interface {
	// basic builder
	New() Entity
}

package entity

type Institution struct {
	ApiBase
	Name   string
	Email  string // custom type
	Owner  User
	Status CurrentStatus
}

// need custom operations to transfer ownership

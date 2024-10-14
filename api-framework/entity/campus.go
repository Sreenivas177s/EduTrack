package entity

type Campus struct {
	ApiBase
	Name        string
	Location    string // custom struct
	Type        CampusType
	HomePageURL string // url
}

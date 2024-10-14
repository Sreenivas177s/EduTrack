package entity

// ---------- running status of a type
type CurrentStatus int

const (
	INACTIVE CurrentStatus = iota
	ACTIVE
	RESIGNED
)

// ----------------- academia type
type CampusType int

const (
	PRIMARY_SCHOOL CampusType = iota
	SECONDARY_SCHOOL
	HIGHER_SECONDARY_SCHOOL
	UNIVERSITY
)

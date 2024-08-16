package entites

type Entity struct {
	name        string
	endpoint    string
	httpHandler EntityMethods
}

// http methods to be implemented by
type EntityMethods interface {
	post()
	put()
	delete()
	get()
	getall()
}

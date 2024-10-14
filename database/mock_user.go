package database

import (
	"chat-server/api-framework/entity"
)

var userMap map[string]entity.User = map[string]entity.User{
	"test@email.com": entity.User{EmailId: "test@email.com", FirstName: "testfname", LastName: "testlname"},
	"user@email.com": entity.User{EmailId: "user@email.com", FirstName: "user1"},
}

func GetUser(email string) entity.User {
	return userMap[email]
}

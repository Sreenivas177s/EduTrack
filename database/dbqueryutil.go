package database

import (
	"chat-server/api/entity"
	"errors"
)

func GetUserByEmail(email string) (*entity.User, error) {
	dbref := GetDBRef()
	var user entity.User
	result := dbref.Where("email_id = ?", email).First(&user)
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func GetUserByID(id uint) (*entity.User, error) {
	dbref := GetDBRef()
	var user entity.User
	result := dbref.First(&user, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

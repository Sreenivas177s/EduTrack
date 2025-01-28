package auth

import (
	"bytes"
	"chat-server/api/entity"
	"chat-server/database"
	"chat-server/utils"
	"crypto"
	"errors"
	"os"
)

type LoginForm struct {
	Email_id    string `json:"email_id"`
	Password    string `json:"password"`
	CallbackUrl string `json:"callback_url"`
	Source      string `json:"source"`
}

func AuthorizeUser(identifier, password string) (*entity.User, error) {
	if identifier != "" && password != "" {
		user, err := database.GetUserByEmail(identifier)
		if err != nil {
			return nil, errors.New("invalid credentials")
		}
		hashedPass, _ := utils.GetHashedPassword(password, user.Salt)
		result := bytes.Equal(hashedPass, user.HashedPassword)
		if result {
			return user, nil
		}
	}
	return nil, errors.New("unable to Find user")
}

func GetJWTSigningKey() []byte {
	plainSecret := os.Getenv("AUTH_SECRET")
	if plainSecret == "" {
		panic("AUTH_SECRET environment variable is not set")
	}
	return crypto.SHA256.New().Sum([]byte(plainSecret))
}

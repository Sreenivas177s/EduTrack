package auth

import (
	"chat-server/api/entity"
	"chat-server/database"
	"crypto"
	"errors"
	"os"
)

type LoginForm struct {
	Email_id string `json:"email_id"`
	Password string `json:"password"`
}

func AuthorizeUser(identifier, password string) (*entity.User, error) {
	if identifier != "" && password != "" {
		user, err := database.GetUserByEmail(identifier)
		if err != nil {
			return nil, errors.New("invalid credentials")
		}
		return user, nil
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

package auth

import (
	"chat-server/api-framework/entity"
	"crypto"
	"errors"
	"os"
)

func AuthorizeUser(identifier, password string) (*entity.User, error) {
	if identifier != "" && password != "" {
		return &entity.User{
			EmailId:   identifier,
			FirstName: identifier,
		}, nil
	}
	return nil, errors.New("unable to Find user")
}

func GetJWTSigningKey() []byte {
	plainSecret := os.Getenv("AUTH_SECRET")
	return crypto.SHA256.New().Sum([]byte(plainSecret))
}

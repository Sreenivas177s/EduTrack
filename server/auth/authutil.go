package auth

import (
	"chat-server/api/entity"
	"chat-server/database"
	"crypto"
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
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
		err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			return user, nil
		}
		return nil, err
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

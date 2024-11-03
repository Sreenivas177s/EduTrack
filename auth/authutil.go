package auth

import (
	"chat-server/api/entity"
	"chat-server/database"
	"crypto"
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func AuthorizeUser(identifier, password string) (*entity.User, error) {
	if identifier != "" && password != "" {
		user := entity.User{
			EmailId: identifier,
		}
		dbref := database.GetDBRef()
		result := dbref.Where(&user).First(&user)
		if result.RowsAffected == 1 {
			salted := append([]byte(password), user.Salt...)
			if err := bcrypt.CompareHashAndPassword(user.HashedPassword, salted); err != nil {
				return nil, errors.New("invalid username/password")
			}
			return &user, nil
		}
	}
	return nil, errors.New("unable to Find user")
}

func GetJWTSigningKey() []byte {
	plainSecret := os.Getenv("AUTH_SECRET")
	return crypto.SHA256.New().Sum([]byte(plainSecret))
}

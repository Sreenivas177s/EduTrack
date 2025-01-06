package entity

import (
	"crypto/rand"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ApiBase
	FirstName   string        `json:"first_name" gorm:"unique;index"`
	LastName    string        `json:"last_name"`
	DateOfBirth time.Time     `json:"date_of_birth"`
	EmailId     string        `json:"email_id" gorm:"unique;uniqueIndex" validate:"email"`
	Status      CurrentStatus `json:"user_status" gorm:"default:1"`

	// write only data
	Password string `json:"password,omitempty" gorm:"-:all"`
	// hidden fields
	Salt           []byte `json:"-"`
	HashedPassword []byte `json:"-"`
}

// api entity methods
func (user *User) Authorize(httpMethod string) (bool, error) {
	switch httpMethod {
	case fiber.MethodPost:
		return true, nil
	case fiber.MethodGet:
		return true, nil
	}

	return false, fiber.NewError(fiber.StatusUnauthorized)
}

func (user *User) Preprocessor(httpMethod string) error {
	if fiber.MethodPost != httpMethod && fiber.MethodPut != httpMethod {
		return errors.New("invalid Params")
	}
	// process data
	if user.Password != "" {
		salt, err := generateSalt()
		if err != nil {
			return err
		}
		hashBytes := append([]byte(user.Password), salt...)
		hashedPassword, err := bcrypt.GenerateFromPassword(hashBytes, bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user.Salt = salt
		user.HashedPassword = hashedPassword
	}
	return nil
}
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}
func (user *User) SubTables() {

}

func (user *User) ID() uint {
	return user.ApiBase.ID
}

func (user *User) Validate(httpMethod string) error {
	return nil
}
func (user *User) HandleOperation(operation string) error {
	return nil
}

// function option methods
func (user *User) FillDefaults() error {
	log.Debug("reached method")
	// user.LastUpdatedTime = time.Now()
	return nil
}

func (user *User) RemoveInternalFields() error {
	log.Debug("reached method")
	return nil
}

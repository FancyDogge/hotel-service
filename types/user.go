package types

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 7
)

// for a request scope, ?data-oriented approach?
// валидация данных при json запросе к серверу
type CreateUserParams struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() []string {
	errors := []string{}
	if len(params.FirstName) < minFirstNameLen {
		err := fmt.Sprintf("firstName length should be at least %d characters", minFirstNameLen)
		errors = append(errors, err)
	}
	if len(params.LastName) < minLastNameLen {
		err := fmt.Sprintf("lastName length should be at least %d characters", minLastNameLen)
		errors = append(errors, err)
	}
	if len(params.Password) < minPasswordLen {
		err := fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
		errors = append(errors, err)
	}
	if !isEmailValid(params.Email) {
		err := fmt.Sprintf("email is invalid")
		errors = append(errors, err)
	}
	// if errors != nil {
	// 	return errors
	// }
	return errors
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username          string             `bson:"username,omitempty" json:"username,omitempty"`
	FirstName         string             `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName          string             `bson:"lastName,omitempty" json:"lastName,omitempty"`
	Email             string             `bson:"email,omitempty" json:"email,omitempty"`
	EncryptedPassword string             `bson:"EncryptedPassword,omitempty" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		Username:          params.Username,
		EncryptedPassword: string(encpw), //we neet to convert to string потому что encpw сейчас []byte
	}, nil
}

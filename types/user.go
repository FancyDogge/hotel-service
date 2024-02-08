package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

// for a request scope, ?data-oriented approach?
// валидация данных при json запросе к серверу
type CreateUserParams struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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

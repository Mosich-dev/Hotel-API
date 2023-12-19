package types

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	bcryptCost   = 12
	MinFirstName = 2
	MinLastName  = 2
	minPassword  = 7
)

type baseUserParams struct {
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}

type CreateUserParams struct {
	baseUserParams
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserParams struct {
	baseUserParams
}

func (p UpdateUserParams) ToBsonM() bson.M {

	d := bson.M{}
	if len(p.FirstName) > 0 {
		d["firstName"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		d["lastName"] = p.LastName
	}
	return d
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"EncryptedPassword" json:"-"`
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func (params CreateUserParams) ValidateAll() map[string]string {
	errors := map[string]string{}

	if len(params.FirstName) < MinFirstName {
		errors["firstName"] = fmt.Sprintf("first name length most be atleast %d characters", MinFirstName)
	}
	if len(params.LastName) < MinLastName {
		errors["lastName"] = fmt.Sprintf("last name length most be atleast %d characters", MinLastName)
	}
	if len(params.Password) < minPassword {
		errors["password"] = fmt.Sprintf("password length most be atleast %d characters", minPassword)
	}
	if !isEmailValid(params.Email) {
		errors["email"] = fmt.Sprintf("email is unvalid")
	}

	return errors
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
		EncryptedPassword: string(encpw),
	}, nil
}

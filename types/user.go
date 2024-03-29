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
	MinPassword  = 7
)

type InsertUserParams struct {
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password"`
}

type UpdateUserParams struct {
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}

type UpdateUserParamsBsonM struct {
	bson.M
}

func (p UpdateUserParams) ToBsonM() UpdateUserParamsBsonM {

	d := bson.M{}

	if len(p.FirstName) > 0 {
		d["firstName"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		d["lastName"] = p.LastName
	}
	return UpdateUserParamsBsonM{d}
}

func (m UpdateUserParamsBsonM) Len(key string) int {
	return len(fmt.Sprintf("%v", m.M[key]))
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

func (params InsertUserParams) ValidateAll() map[string]string {
	errors := map[string]string{}

	if len(params.FirstName) < MinFirstName {
		errors["firstName"] = fmt.Sprintf("first name length most be atleast %d characters", MinFirstName)
	}
	if len(params.LastName) < MinLastName {
		errors["lastName"] = fmt.Sprintf("last name length most be atleast %d characters", MinLastName)
	}
	if len(params.Password) < MinPassword {
		errors["password"] = fmt.Sprintf("password length most be atleast %d characters", MinPassword)
	}
	if !isEmailValid(params.Email) {
		errors["email"] = fmt.Sprintf("email is invalid")
	}

	return errors
}

func NewUserFromParams(params InsertUserParams) (*User, error) {
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

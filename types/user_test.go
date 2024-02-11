package types

import (
	"strings"
	"testing"
)

func TestInsertUserValidator(t *testing.T) {
	scenarios := map[string]InsertUserParams{
		"validParams": {
			FirstName: "reza",
			LastName:  "chegeni",
			Email:     "rrr.reza18@gmail.com",
			Password:  "1234567",
		},
		"invalidParams": {
			FirstName: "r",
			LastName:  "c",
			Email:     "wrong_email12345",
			Password:  "123456",
		},
		"invalidEmailParams": {
			FirstName: "reza",
			LastName:  "chegeni",
			Email:     "wrong_email12345",
			Password:  "1234567",
		},
		"invalidFirstNameParams": {
			FirstName: "r",
			LastName:  "chegeni",
			Email:     "rrr.reza18@gmail.com",
			Password:  "1234567",
		},
		"invalidLastNameParams": {
			FirstName: "reza",
			LastName:  "c",
			Email:     "rrr.reza18@gmail.com",
			Password:  "1234567",
		},
		"invalidPasswordParams": {
			FirstName: "reza",
			LastName:  "chegeni",
			Email:     "rrr.reza18@gmail.com",
			Password:  "123456",
		},
	}

	for scenario, params := range scenarios {
		valErrs := params.ValidateAll()
		switch scenario {
		case "validParams":
			if len(valErrs) != 0 {
				t.Errorf("all params are valid. NONE of params most fail but %d did:\n", len(valErrs))
				for param, err := range valErrs {
					t.Errorf("%s validator failed. error: %s\n", strings.ToUpper(param), err)
				}
			}

		case "invalidParams":
			if len(valErrs) != 4 {
				t.Errorf("none of the params are valid. ALL 4 params most fail but only %d did:\n", len(valErrs))
			}

		case "invalidEmailParams":
			if len(valErrs["firstName"]) != 0 || len(valErrs["lastName"]) != 0 || len(valErrs["password"]) != 0 || len(valErrs["email"]) == 0 {
				t.Errorf("email validator failed.")
			}

		case "invalidFirstNameParams":
			if len(valErrs["firstName"]) == 0 || len(valErrs["lastName"]) != 0 || len(valErrs["password"]) != 0 || len(valErrs["email"]) != 0 {
				t.Error("firstName validator failed.")
			}

		case "invalidLastNameParams":
			if len(valErrs["firstName"]) != 0 || len(valErrs["lastName"]) == 0 || len(valErrs["password"]) != 0 || len(valErrs["email"]) != 0 {
				t.Error("firstName validator failed.")
			}

		case "invalidPasswordParams":
			if len(valErrs["firstName"]) != 0 || len(valErrs["lastName"]) != 0 || len(valErrs["password"]) == 0 || len(valErrs["email"]) != 0 {
				t.Error("password validator failed.")
			}
		}
	}
}

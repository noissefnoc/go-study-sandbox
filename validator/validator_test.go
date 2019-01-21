package main

import (
	"gopkg.in/go-playground/validator.v9"
	"regexp"
	"testing"
)

// validation target
type CreateUserPayload struct {
	Username string `validate:"username"`
}

// custom validate function
func usernameValidation(fl validator.FieldLevel) bool {
	tf, err := regexp.Match(`^[a-z][a-z0-9-]{1,32}$`, []byte(fl.Field().String()))

	if tf && err == nil {
		return true
	}

	return false;
}

func TestUsernameCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", usernameValidation)

	// valid username
	err := validate.Struct(CreateUserPayload{Username: "hoge"})
	if err != nil {
		t.Fatalf("failed test: %v", err)
	}

	// invalid username length under 1
	err = validate.Struct(CreateUserPayload{Username: ""})
	if err == nil {
		t.Fatalf("failed test: %v", err)
	}
}

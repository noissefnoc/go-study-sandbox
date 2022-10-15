package validator

import (
    "github.com/go-playground/validator/v10"
	"regexp"
	"testing"
)

// validation target
type CreateUserPayload struct {
	Username            string `validate:"required,username"`
	Token               string `validate:"required,token"`
	AgreeTermsOfService string `validate:"required,oneof=yes no"`
	NotMinor            string `validate:"required,oneof=yes no"`
}

// custom validate function
func usernameValidation(fl validator.FieldLevel) bool {
	tf, err := regexp.Match(`^[a-z][a-z0-9-]{1,32}$`, []byte(fl.Field().String()))

	if tf && err == nil {
		return true
	}

	return false;
}

func tokenValidation(fl validator.FieldLevel) bool {
	tf, err := regexp.Match(`^[ -~]{8,128}$`, []byte(fl.Field().String()))

	if tf && err == nil {
		return true
	}

	return false;
}

func TestCreateUserPayloadValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", usernameValidation)
	validate.RegisterValidation("token", tokenValidation)

	// valid payload
	err := validate.Struct(
		CreateUserPayload{
			Username:            "hoge-2",
			Token:               "hogehogehoge",
			AgreeTermsOfService: "yes",
			NotMinor:            "yes"})

	if err != nil {
		t.Fatalf("failed test: %v", err)
	}

	err = validate.Struct(
		CreateUserPayload{
			Username:            "hoge-2",
			Token:               "hogehogehoge",
			AgreeTermsOfService: "no",
			NotMinor:            "no"})

	if err != nil {
		t.Fatalf("failed test: %v", err)
	}

	// invalid payload
	// case length under 1
	err = validate.Struct(CreateUserPayload{Username: ""})
	if err == nil {
		t.Fatalf("failed test: %v", err)
	}

	// case length over 32
	err = validate.Struct(CreateUserPayload{Username: "hogehogehogehogehogehogehogehogehogehoge"})
	if err == nil {
		t.Fatalf("failed test: %v", err)
	}
}

package validator

import (
	"gopkg.in/go-playground/validator.v9"
	"testing"
)

type User struct {
	Name string `validate:"gte=2"`
}

func TestValidate(t *testing.T) {
	validate := validator.New()

	err := validate.Struct(User {Name: "a"})

	if err == nil {
		t.Errorf("wrong")
	}
}

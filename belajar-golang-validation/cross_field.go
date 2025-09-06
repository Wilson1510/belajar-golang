package main

import (
	"fmt"
	"strings"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Username string `validate:"required,must_equal_ignore_case=Password|must_equal_ignore_case=Email"`
	Password string `validate:"required"`
	Email string `validate:"required,email"`
}

func MustEqualIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("field not found")
	}
	firstValue := strings.ToLower(field.Field().String())
	secondValue := strings.ToLower(value.String())
	return firstValue == secondValue
}

func main() {
	validate := validator.New()
	validate.RegisterValidation("must_equal_ignore_case", MustEqualIgnoreCase)

	user := User{
		Username: "wiLSon@gmail.com",
		Password: "wilson1",
		Email: "wilson@gmail.com",
	}
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
	}
}
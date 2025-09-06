package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type RegisterUser struct {
	Username string `validate:"required,min=3"`
	Email string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

func MustMatch(field validator.StructLevel) {
	registerUser := field.Current().Interface().(RegisterUser)
	if registerUser.Password == registerUser.Email || registerUser.Password == registerUser.Username {
	} else{
		field.ReportError(registerUser.Password, "Password", "Password", "must_match", "")
	}
}

func main() {
	validate := validator.New()
	validate.RegisterStructValidation(MustMatch, RegisterUser{})

	registerUser := RegisterUser{
		Username: "wilson5",
		Email: "wilson@gmail.com",
		Password: "wilson",
	}
	err := validate.Struct(registerUser)
	if err != nil {
		fmt.Println(err)
	}
}
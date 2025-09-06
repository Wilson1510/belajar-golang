package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	ID int `validate:"required,min=1"`
	Name string `validate:"required,min=3"`
	Email string `validate:"required,email"`
}

func main() {
	user := User{
		ID: 0,
		Name: "",
		Email: "john.doeexample.com",
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			fmt.Println(validationError.Field(), validationError.Tag(), validationError.Error())
		}
	}
}
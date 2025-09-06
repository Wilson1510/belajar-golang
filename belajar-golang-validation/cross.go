package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserInput struct {
	Name string `validate:"required,min=3"`
	Password string `validate:"required,min=6,eqfield=ConfirmPassword"`
	ConfirmPassword string `validate:"required"`
}

func main() {
	userInput := UserInput{
		Name: "John Doe",
		Password: "123456",
		ConfirmPassword: "123456",
	}

	validate := validator.New()
	err := validate.Struct(userInput)
	if err != nil {
		fmt.Println(err)
	}
}
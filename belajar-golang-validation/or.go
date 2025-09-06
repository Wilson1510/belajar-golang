package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)


func main() {
	validate := validator.New()

	err := validate.Var("wilson", "required,email|numeric")
	if err != nil {
		fmt.Println(err)
	}

	err = validate.Var("123456", "required,email|numeric")
	if err != nil {
		fmt.Println(err)
	}

	err = validate.Var("wilson@gmail.com", "required,email|numeric")
	if err != nil {
		fmt.Println(err)
	}
}
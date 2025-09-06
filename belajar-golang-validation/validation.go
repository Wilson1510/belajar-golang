package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func main() {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		panic("validate is nil")
	}

	user := "john"
	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err)
	}

	password := "123456"
	confirmPassword := "123456"
	err = validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err)
	}

	phone := "1234567890"
	err = validate.Var(phone, "required,numeric,min=10,max=15")
	if err != nil {
		fmt.Println(err)
	}

}
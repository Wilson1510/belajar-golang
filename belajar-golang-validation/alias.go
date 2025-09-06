package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	String1 string `validate:"varchar"`
	String2 string `validate:"varchar"`
	String3 string `validate:"varchar"`
	String4 string `validate:"varchar"`
}

func main() {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,min=3")
	user := User{
		String1: "1",
		String2: "2",
		String3: "3",
		String4: "4",
	}
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
	}
}
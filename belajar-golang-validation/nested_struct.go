package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Address struct {
	Street string `validate:"required"`
	City string `validate:"required"`
	Province string `validate:"required"`
}

type UserInfo struct {
	Name string `validate:"required"`
	Address Address `validate:"required"`
}

func main() {
	userInfo := UserInfo{
		Name: "John Doe",
		Address: Address{
			Street: "Jl. Raya",
			City: "Jakarta",
			Province: "DKI Jakarta",
		},
	}

	validate := validator.New()
	err := validate.Struct(userInfo)
	if err != nil {
		fmt.Println(err)
	}
}
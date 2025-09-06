package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Education struct {
	School string `validate:"required"`
}

type Address struct {
	Street string `validate:"required"`
	City string `validate:"required"`
	Province string `validate:"required"`
}

type UserInfo struct {
	Name string `validate:"required"`
	Addresses []Address `validate:"required,dive"`
	Education map[string]Education `validate:"dive,keys,required,min=2,endkeys,required"`
}

func main() {
	userInfo := UserInfo{
		Name: "John Doe",
		Addresses: []Address{
			{
				Street: "",
				City: "Bandung",
				Province: "",
			},
			{
				Street: "Jl. Raya 2",
				City: "",
				Province: "",
			},
		},
		Education: map[string]Education{
			"": {
				School: "SDN 1",
			},
			"secondary": {
				School: "",
			},
		},
	}

	validate := validator.New()
	err := validate.Struct(userInfo)
	if err != nil {
		fmt.Println(err)
	}
}
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"github.com/go-playground/validator/v10"
)

func Pin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}
	regexNumber := regexp.MustCompile(`^[0-9]+$`)
	return regexNumber.MatchString(field.Field().String()) && len(field.Field().String()) == length
}

func ValidateCustom(field validator.FieldLevel) bool {
	return field.Field().String() == "wilson"
}

func main() {
	validate := validator.New()
	validate.RegisterValidation("custom", ValidateCustom)

	username := "wilson"
	err := validate.Var(username, "required,custom")
	if err != nil {
		fmt.Println(err)
	}

	validate.RegisterValidation("pin", Pin)

	pin := "123456"
	err = validate.Var(pin, "required,pin=6")
	if err != nil {
		fmt.Println(err)
	}
}


package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func main() {
	var validate *validator.Validate = validator.New()

	user := []string{"john", "jane", "jim"}
	err := validate.Var(user, "dive,required,min=5")
	if err != nil {
		fmt.Println(err)
	}

	balance := map[string]int{
		"john": 100,
		"j": 200,
		"jim": 300,
	}
	err = validate.Var(balance, "dive,keys,required,min=2,endkeys,required,gt=100")
	if err != nil {
		fmt.Println(err)
	}
}
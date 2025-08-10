package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func saveData(id int) error {
	if id == 0 {
		return &ValidationError{Message: "Validation Error"}
	}
	if id == 1 {
		return &NotFoundError{Message: "Not Found"}
	}
	return nil
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	err = saveData(2)
	if err != nil {
		if vali, ok := err.(*ValidationError); ok {
			fmt.Println("Validation Error:", vali.Message)
		} else if notFou, ok := err.(*NotFoundError); ok {
			fmt.Println("Not Found Error:", notFou.Message)
		} else {
			fmt.Println("Unknown Error:", err)
		}
	}
}

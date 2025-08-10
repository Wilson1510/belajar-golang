package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{
		"name": "John",
		"age": "20",
	}
	fmt.Println(person)

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}

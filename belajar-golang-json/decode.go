package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street string
	City string
	Country string
}

type Person struct {
	Name string
	Age int
	Single bool
	Hobbies []string
	Address []Address
}

func main() {
	json_object := `{"name": "John", "age": 20, "single": true, "hobbies": ["Coding", "Reading", "Traveling"], "address": [{"street": "Jl. Raya", "city": "Jakarta", "country": "Indonesia"}, {"street": "Jl. Raya 2", "city": "Bandung", "country": "Indonesia"}]}`

	var person Person
	err := json.Unmarshal([]byte(json_object), &person)
	if err != nil {
		panic(err)
	}
	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Println(person.Single)
	fmt.Println(person.Hobbies)
	fmt.Println(person.Address)
}
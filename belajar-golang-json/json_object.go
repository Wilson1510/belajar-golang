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
	person := Person{
		Name: "John",
		Age: 20,
		Single: true,
		Hobbies: []string{"Coding", "Reading", "Traveling"},
		Address: []Address{
			{Street: "Jl. Raya", City: "Jakarta", Country: "Indonesia"},
			{Street: "Jl. Raya 2", City: "Bandung", Country: "Indonesia"},
		},
	}

	json, _ := json.Marshal(person)
	fmt.Println(string(json))
	
}
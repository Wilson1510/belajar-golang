package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Single bool `json:"single"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	reader, _ := os.Open("sample.json")
	decoder := json.NewDecoder(reader)

	var person Person
	err := decoder.Decode(&person)
	if err != nil {
		panic(err)
	}
	fmt.Println(person)
}
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"KjH"`
	Married bool `json:"is_mArried"`
}

func main() {
	person := Person{
		Name: "John", 
		Age: 20,
		Married: false,
	}

	json_object, _ := json.Marshal(person)
	fmt.Println(string(json_object))

	json_string := `{"name": "John", "kjh": 20, "is_marrIed": false}`
	var person2 Person
	err := json.Unmarshal([]byte(json_string), &person2)
	if err != nil {
		panic(err)
	}
	fmt.Println(person2)
}

package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name string
	Age int `required:"true" min:"18"`
}

func readField(data any) {
	var valueType reflect.Type = reflect.TypeOf(data)
	fmt.Println(valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		fmt.Println(valueType.Field(i).Name, valueType.Field(i).Type)
		fmt.Println(valueType.Field(i).Tag.Get("required"))
		fmt.Println(valueType.Field(i).Tag.Get("max"))
		fmt.Println(valueType.Field(i).Tag.Get("min"))
	}
}

func isValid(data any) bool {
	result := true
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			result = value != ""
			if result == false {
				return result
			}	
		}
	}
	return result
}

func main() {
	person := Person{
		Name: "",
		Age: 20,
	}

	sample := Sample{
		Name: "",
	}

	readField(person)
	readField(sample)

	fmt.Println(isValid(person))
	fmt.Println(isValid(sample))
}
package main

import (
	"encoding/json"
	"os"
)

type Category struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func main() {
	writer, _ := os.Create("sample.json")
	encoder := json.NewEncoder(writer)

	category := Category{
		Id: "1",
		Name: "Category 1",
	}

	encoder.Encode(category)
}
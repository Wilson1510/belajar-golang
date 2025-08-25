package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	json_object := `{"name": "John", "age": 20, "single": true, "hobbies": ["Coding", "Reading", "Traveling"]}`

	var person map[string]interface{}
	err := json.Unmarshal([]byte(json_object), &person)
	if err != nil {
		panic(err)
	}
	fmt.Println(person)

	map_json := map[string]interface{}{
		"name": "John",
		"age": 20,
		"single": true,
		"hobbies": []string{"Coding", "Reading", "Traveling"},
	}

	json_map, _ := json.Marshal(map_json)
	fmt.Println(string(json_map))
}	
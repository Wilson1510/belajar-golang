package main

import (
	"encoding/json"
	"fmt"
)

func EncodeJson(data interface{}) string {
	byte, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	return string(byte)
}

func main() {
	json := EncodeJson("Eko")
	fmt.Println(json)

	json = EncodeJson(1)
	fmt.Println(json)

	json = EncodeJson(true)
	fmt.Println(json)

	json = EncodeJson([]string{"Eko", "Kurniawan", "Khannedy"})
	fmt.Println(json)

	var b byte = '9' // b adalah byte dengan nilai karakter 'A'
	fmt.Println(b)
	s := string(b)   // s adalah string yang isinya "A"
	fmt.Println(s)   // Output: A
}
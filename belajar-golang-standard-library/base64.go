package main

import (
	"fmt"
	"encoding/base64"
)

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte("Eko Kurniawan Khannedy"))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}
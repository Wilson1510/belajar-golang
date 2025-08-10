package main

import (
	"fmt"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Hostname:", hostname)
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println("Username:", username)
	fmt.Println("Password:", password)

	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

}